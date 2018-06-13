package message

import (
	"container/list"
	"encoding/json"
	"fmt"
)

type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type markdownMessage struct {
	MsgType  string    `json:"msgtype"`
	Markdown *markdown `json:"markdown"`
	At       *at       `json:"at"`
}

// get instance of MarkdownMessage
func GetNewMarkdownMessage(title, content string, mobiles *list.List, isAtAll bool) *markdownMessage {
	if content == "" || title == "" {
		return nil
	}
	markdown := &markdown{Title: title, Text: content}

	at := &at{IsAtAll: isAtAll}
	if nil != mobiles && mobiles.Len() > 0 {
		for p := mobiles.Front(); p != nil; p = p.Next() {
			at.AtMobiles = append(at.AtMobiles, p.Value.(string))
		}
	}
	return &markdownMessage{MsgType: MESSAGE_TYPE_MARKDOWN, Markdown: markdown, At: at}
}

// ToJsonString
func (mm *markdownMessage) ToJsonString() string {
	content, err := json.Marshal(mm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
	return ""
}
