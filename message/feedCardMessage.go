package message

import (
	"container/list"
	"encoding/json"
	"fmt"
)

type links struct {
	Links []Link `json:"links"`
}

type feedCardMessage struct {
	MsgType  string    `json:"msgtype"`
	FeedCard *links `json:"feedCard"`
}

func GetNewFeedCardMessage(feeds *list.List) *feedCardMessage {
	links := &links{}
	for p := feeds.Front(); p != nil; p = p.Next() {
		links.Links = append(links.Links, p.Value.(Link))
	}
	fmt.Println(links)
	fm := &feedCardMessage{MsgType: MESSAGE_TYPE_FEEDCART, FeedCard: links}
	return fm
}

func (fm *feedCardMessage) ToJsonString() string {
	content, err := json.Marshal(fm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
}
