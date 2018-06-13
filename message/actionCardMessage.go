package message

import (
	"encoding/json"
	"fmt"
)

type actionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
}

type actionCardMessage struct {
	MsgType    string      `json:"msgtype"`
	ActionCard *actionCard `json:"actionCard"`
}

// get instance of ActionCardMessage
func GetNewActionCardMessage(title, content, hideAvatar, btnOrientation, singleTitle, singleUrl string) *actionCardMessage {
	if title == "" || content == "" {
		return nil
	}
	ac := &actionCard{Title: title, Text: content, HideAvatar: hideAvatar, BtnOrientation: btnOrientation, SingleTitle: singleTitle, SingleURL: singleUrl}
	return &actionCardMessage{MsgType: MESSAGE_TYPE_ACTIONCART, ActionCard: ac}
}

// ToJsonString
func (am *actionCardMessage) ToJsonString() string {
	content, err := json.Marshal(am)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(content)
}
