package pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gazizvr/tg-bot-api/internal"
)

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

func (c *Client) downloadFile(
	uniqueId, path, dirPath string,
) (*string, error) {
	fileName := fmt.Sprint(uniqueId, filepath.Ext(path))
	filePath := filepath.Join(dirPath, fileName)

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = fmt.Sprint("file", "/", c.urlPath(path))
	urlStr := u.String()

	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, ErrFileNotFound
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	if _, err := io.Copy(newFile, resp.Body); err != nil {
		return nil, err
	}
	return &filePath, nil
}

func (c *Client) DownloadFile(
	id string,
	dirPath string,
) (*string, error) {
	fileResp, err := c.getFile(id)
	if err != nil {
		return nil, err
	}
	filePath := *fileResp.Result.Path
	if len(strings.TrimSpace(filePath)) < 1 {
		return nil, ErrFileNotFound
	}
	if string(filePath[0]) != "/" {
		newPath, err := c.downloadFile(
			fileResp.Result.UniqueId,
			filePath,
			dirPath,
		)
		if err != nil {
			return nil, err
		}
		return newPath, nil
	}
	return &filePath, nil
}
