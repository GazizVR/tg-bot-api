package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type CommonResponse struct {
	Ok bool `json:"ok"`
}

type UpdatesResponse struct {
	CommonResponse
	Result []Update `json:"result"`
}

type MessageResponse struct {
	CommonResponse
	Result Message `json:"result"`
}

type FileResponse struct {
	CommonResponse
	Result File `json:"result"`
}

type ErrorResponse struct {
	CommonResponse
	Code        uint16 `json:"error_code"`
	Description string `json:"description"`
}

func CheckError(
	resp CommonResponse,
	body []byte,
	method string,
) error {
	if !resp.Ok {
		var errRespp ErrorResponse
		if err := json.Unmarshal(body, &errRespp); err != nil {
			log.Println("Ошибка преобразования raw bytes to json: ", err)
			return err
		}
		errStr := fmt.Sprintf(
			"Error %s %d: %s\n",
			method,
			errRespp.Code,
			errRespp.Description,
		)
		err := errors.New(errStr)
		log.Print(err.Error())
		return err
	}
	return nil
}
