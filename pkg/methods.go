package pkg

import (
	"fmt"

	"github.com/gazizvr/tg-bot-api/internal"
)

const (
	getUpdatesMethod      = "getUpdates"
	sendMessageMethod     = "sendMessage"
	editMediaMethod       = "editMessageMedia"
	editReplyMarkupMethod = "editMessageReplyMarkup"
	editTextMethod        = "editMessageText"
	answerQueryMethod     = "answerCallbackQuery"
	getFileMethod         = "getFile"
)

func (c *Client) getUpdates(
	offset int64,
	limit uint8,
	timeout uint8,
	allowedUpdates []string,
) (*UpdatesResponse, error) {
	var response UpdatesResponse

	params := map[string]string{
		"offset":          fmt.Sprintf("%d", offset),
		"limit":           fmt.Sprintf("%d", limit),
		"timeout":         fmt.Sprintf("%d", timeout),
		"allowed_updates": fmt.Sprintf("%v", allowedUpdates),
	}

	body, err := internal.GetRequest(
		c.BaseURL,
		c.urlPath(getUpdatesMethod),
		params,
		&response,
	)
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
	return &response, nil
}

func (c *Client) AnswerCallbackQuery(
	queryId string,
) (*CommonResponse, error) {
	params := map[string]string{
		"callback_query_id": queryId,
	}

	response, _, err := c.commonRequest(
		params,
		answerQueryMethod,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetFile(
	id string,
) (*FileResponse, error) {
	params := map[string]string{
		"file_id": id,
	}
	var response FileResponse
	body, err := internal.GetRequest(
		c.BaseURL,
		c.urlPath(getFileMethod),
		params,
		response,
	)
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
	return &response, nil
}
