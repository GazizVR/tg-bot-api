package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	editMediaMethod       = "editMessageMedia"
	editReplyMarkupMethod = "editMessageReplyMarkup"
	editTextMethod        = "editMessageText"
)

func (c *Client) EditMessageMedia(
	chatId int64,
	messageId int64,
	mediaType string,
	media os.File,
	markup *InlineMarkup,
) (*MessageResponse, error) {
	params := map[string]string{
		"chat_id":    fmt.Sprintf("%d", chatId),
		"message_id": fmt.Sprintf("%d", messageId),
		"media": fmt.Sprintf(`{
			"type": "%s",
			"media": "attach://%s"
		}`, mediaType, mediaType),
	}
	if markup != nil {
		jsonMarkup, err := json.Marshal(markup)
		if err != nil {
			log.Println("Ошибка преобразования markup json: ", err)
			return nil, err
		}
		params["reply_markup"] = string(jsonMarkup)
	}

	response, err := c.mediaRequest(
		params,
		editMediaMethod,
		map[string]os.File{mediaType: media},
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) EditMessageReplyMarkup(
	chatId int64,
	messageId int64,
	markup InlineMarkup,
) (*MessageResponse, error) {
	jsonMarkup, err := json.Marshal(markup)
	if err != nil {
		log.Println("Ошибка преобразования markup json: ", err)
		return nil, err
	}
	params := map[string]string{
		"chat_id":      fmt.Sprintf("%d", chatId),
		"message_id":   fmt.Sprintf("%d", messageId),
		"reply_markup": string(jsonMarkup),
	}

	response, err := c.messageRequest(
		params,
		editReplyMarkupMethod,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) EditMessageText(
	chatId int64,
	messageId int64,
	text string,
	markup *InlineMarkup,
) (*MessageResponse, error) {
	params := map[string]string{
		"message_id": fmt.Sprintf("%d", messageId),
		"chat_id":    fmt.Sprintf("%d", chatId),
		"text":       text,
	}
	if markup != nil {
		jsonMarkup, err := json.Marshal(markup)
		if err != nil {
			log.Println("Ошибка преобразования markup json: ", err)
			return nil, err
		}
		params["reply_markup"] = string(jsonMarkup)
	}

	response, err := c.messageRequest(
		params,
		editTextMethod,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
