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

// get instance of TextMessage
func GetNewTextMessage(content string, mobiles *list.List, isAtAll bool) *textMessage {
	if content == "" {
		return nil
	}

	text := &text{Content: content}
	at := &at{IsAtAll: isAtAll}
	if nil != mobiles && mobiles.Len() > 0 {
		for p := mobiles.Front(); p != nil; p = p.Next() {
			at.AtMobiles = append(at.AtMobiles, p.Value.(string))
		}
	}

	tm := textMessage{MsgType: MESSAGE_TYPE_TEXT, Text: *text, At: at}
	return &tm
}

// ToJsonString
func (tm *textMessage) ToJsonString() string {
	content, err := json.Marshal(tm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
}
