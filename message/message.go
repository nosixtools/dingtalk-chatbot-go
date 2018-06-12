package message

type message interface {
	ToJsonString() string
}

type at struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
