package models

type MessageType string

const (
	TEXT_MESSAGE MessageType = "TEXT_MESSAGE"
)

type MessageAbstract struct {
	BaseModel
	UserID    int         `json:"userId"`
	ChannelID int         `json:"channelId"`
	Type      MessageType `json:"type"`
}

type TextMessage struct {
	MessageAbstract
	Content string `json:"content"`
}

func (MessageAbstract) TableName() string {
	return "messages"
}

func (TextMessage) TableName() string {
	return "messages"
}
