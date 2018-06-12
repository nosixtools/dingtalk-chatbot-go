package dingtalkclient

import (
	"dingtalk-chatbot-go/result"
	"io/ioutil"
	"net/http"
	"strings"
)

func Send(webHook string, message string) (*result.SendResult, error) {

	client := http.Client{}
	request, err := http.NewRequest("POST", webHook, strings.NewReader(message))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sr, err := result.GetNewSendResult(body)
	if err != nil {
		return nil, err
	}
	return sr, nil
}
