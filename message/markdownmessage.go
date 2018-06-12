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

func GetNewMarkdownMessge(title, content string, mobiles *list.List, isAtAll bool) *markdownMessage {

	markdown := &markdown{Title: title, Text: content}
	at := &at{IsAtAll: isAtAll}
	for p := mobiles.Front(); p != nil; p = p.Next() {
		at.AtMobiles = append(at.AtMobiles, p.Value.(string))
	}
	return &markdownMessage{MsgType: MESSAGE_TYPE_MARKDOWN, Markdown: markdown, At: at}
}

func (mm *markdownMessage) ToJsonString() string {
	content, err := json.Marshal(mm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
	return ""
}
