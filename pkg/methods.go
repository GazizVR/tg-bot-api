package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

func (c *Client) getFile(
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

func (c *Client) DownloadFile(
	id string,
	dirPath string,
) (*string, error) {
	fileResp, err := c.getFile(id)
	if err != nil {
		return nil, err
	}
	filePath := fileResp.Result.Path
	if len(strings.TrimSpace(filePath)) < 1 {
		return nil, ErrFileNotFound
	}
	if string(filePath[0]) != "/" {
		fileName := fmt.Sprint(
			fileResp.Result.UniqueId,
			filePath[strings.LastIndex(filePath, "."):],
		)
		newPath := filepath.Join(dirPath, fileName)
		urlPath := fmt.Sprint("file", "/", c.urlPath(filePath))
		resp, err := http.Get(urlPath)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, ErrFileNotFound
		}
		newFile, err := os.Create(newPath)
		if err != nil {
			return nil, err
		}
		defer newFile.Close()
		if _, err := io.Copy(newFile, resp.Body); err != nil {
			return nil, err
		}
		return &newPath, nil
	}
	return &filePath, nil
}
