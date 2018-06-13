package message

import (
	"encoding/json"
	"fmt"
)

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picURL"`
	MessageURL string `json:"messageURL"`
}

type linkMessage struct {
	MsgType string `json:"msgtype"`
	Link    *Link  `json:"link"`
}

// get instance of LinkMessage
func GetNewLinkMessage(title,content,picUrl,messageUrl string) (*linkMessage) {
	if content == "" || messageUrl == "" {
		return nil
	}
	link := &Link{Text:content, Title:title, PicURL:picUrl, MessageURL:messageUrl}
	lk := &linkMessage{MsgType:MESSAGE_TYPE_LINK, Link:link}
	return lk
}

// ToJsonString
func (lm *linkMessage) ToJsonString() string  {
	content, err := json.Marshal(lm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
}
