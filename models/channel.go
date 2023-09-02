package models

type ChannelType string

const (
	DIRECT_CHAT ChannelType = "DIRECT_CHAT"
	BOT_CHAT    ChannelType = "BOT_CHAT"
)

type Channel struct {
	BaseModel
	Name          string
	LastMessageID *int
	LastMessage   MessageAbstract `gorm:"foreignKey:LastMessageID;references:ID"`
	Type          ChannelType
}
