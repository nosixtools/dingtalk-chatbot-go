package dingtalkclient

import (
	"dingtalk-chatbot-go/result"
	"io/ioutil"
	"net/http"
	"strings"
	"net"
	"time"
	"context"
)

type sendConfig struct {
	ConnectionTimeout time.Duration
	RequestTimeout time.Duration
}

type sendClient struct {
	SendConfig *sendConfig
	client *http.Client
}

// get instance of SendClient
func GetNewSendClient(connectionTimeout, requestTimeout time.Duration) *sendClient  {
	if connectionTimeout <= 0 {
		connectionTimeout = time.Second * 3
	}
	if requestTimeout <= 0 {
		requestTimeout = time.Second * 3
	}

	sc := &sendConfig{ConnectionTimeout:connectionTimeout, RequestTimeout:requestTimeout}
	sendClient := &sendClient{SendConfig:sc}
	return sendClient
}

// send message
func (sc *sendClient)Send(webHook string, message string) (*result.SendResult, error) {
	if sc.client == nil {
		client := &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					deadline := time.Now().Add(sc.SendConfig.ConnectionTimeout)
					c, err := net.DialTimeout(network, addr, time.Second*20)
					if err != nil {
						return nil, err
					}
					c.SetDeadline(deadline)
					return c, nil
				},
			},
			Timeout:time.Duration(sc.SendConfig.RequestTimeout),
		}
		sc.client = client
	}

	request, err := http.NewRequest("POST", webHook, strings.NewReader(message))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := sc.client.Do(request)
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
