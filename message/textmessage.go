package message

import (
	"container/list"
	"encoding/json"
	"fmt"
)

type textMessage struct {
	MsgType string `json:"msgtype"`
	Text    text   `json:"text"`
	At      *at    `json:"at"`
}

type text struct {
	Content string `json:"content"`
}

func GetNewTextMessage(content string, mobiles *list.List, isAtAll bool) *textMessage {
	at := &at{IsAtAll: isAtAll}
	for p := mobiles.Front(); p != nil; p = p.Next() {
		at.AtMobiles = append(at.AtMobiles, p.Value.(string))
	}
	text := &text{Content: content}
	tm := textMessage{MsgType: MESSAGE_TYPE_TEXT, Text: *text, At: at}
	return &tm
}

func (tm *textMessage) ToJsonString() string {
	content, err := json.Marshal(tm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
}
