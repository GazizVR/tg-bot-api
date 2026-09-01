package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	sendMessageMethod   = "sendMessage"
	deleteMessageMethod = "deleteMessage"
	sendDocumentMethod  = "sendDocument"
)

func (c *Client) SendMessage(
	chatId int64,
	text string,
	markup *InlineMarkup,
	msgIdToReply *int64,
) (*MessageResponse, error) {
	params := map[string]string{
		"chat_id": fmt.Sprintf("%d", chatId),
		"text":    text,
	}
	if markup != nil {
		jsonMarkup, err := json.Marshal(markup)
		if err != nil {
			log.Println("Ошибка преобразования markup json: ", err)
			return nil, err
		}
		params["reply_markup"] = string(jsonMarkup)
	}
	if msgIdToReply != nil {
		params["reply_parameters"] = fmt.Sprintf(
			`{"message_id": %d}`,
			*msgIdToReply,
		)
	}

	response, err := c.messageRequest(
		params,
		sendMessageMethod,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteMessage(
	chatId int64,
	messageId int64,
) (*CommonResponse, error) {
	params := map[string]string{
		"chat_id":    fmt.Sprintf("%d", chatId),
		"message_id": fmt.Sprintf("%d", messageId),
	}
	response, body, err := c.commonRequest(params, deleteMessageMethod)
	if err != nil {
		return nil, err
	}

	resp := CommonResponse{Ok: response.Ok}
	if err := CheckError(
		resp,
		body,
		getUpdatesMethod,
	); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SendDocument(
	chatId int64,
	mediaType string,
	media os.File,
	msgIdToReply *int64,
) (*MessageResponse, error) {
	params := map[string]string{
		"chat_id": fmt.Sprintf("%d", chatId),
		"document": fmt.Sprintf(`{
			"type": "%s",
			"media": "attach://%s"
		}`, mediaType, mediaType),
	}
	if msgIdToReply != nil {
		params["reply_parameters"] = fmt.Sprintf(
			`{"message_id": %d}`,
			*msgIdToReply,
		)
	}

	response, err := c.mediaRequest(
		params,
		sendDocumentMethod,
		map[string]os.File{mediaType: media},
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
