package result

import (
	"encoding/json"
	"errors"
)

type SendResult struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}

func GetNewSendResult(body []byte) (*SendResult, error) {
	if nil == body {
		return nil, errors.New("empty body")
	}
	sr := &SendResult{}
	err := json.Unmarshal(body, sr)
	if err != nil {
		return nil, err
	}
	return sr, nil
}
