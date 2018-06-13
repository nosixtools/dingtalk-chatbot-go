package dingtalkchatbotgo

import (
	"testing"
	"container/list"
	"github.com/nosixtools/dingtalk-chatbot-go/message"
	"time"
)

var WEBHOOK = "https://oapi.dingtalk.com/robot/send?access_token=c8e46870d0cd473e778763e4613b81644aa1b32f6ac59ae0ab398d404c600552"

func TestSend(t *testing.T) {
	messageText := "hello world"
	messageMobiles := list.New()
	messageMobiles.PushBack("13261656404")

	sc := GetNewSendClient(time.Second, time.Second)
	//发送数据
	sr, err := sc.Send(WEBHOOK, message.GetNewTextMessage(messageText, messageMobiles, false).ToJsonString())
	if err != nil {
		t.Error("send error", err)
	}

	if sr.ErrorCode != 0 {
		t.Error("send dingtalk message failed", sr)
	}
	t.Log(sr)
}
