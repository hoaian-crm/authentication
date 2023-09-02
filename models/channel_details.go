package models

type ChannelDetail struct {
	BaseModel
	UserID            int
	ChannelID         int
	LastMessageSeenID *int            `json:"-"`
	LastMessageSeen   MessageAbstract `gorm:"foreginKey:LastMessageSeenID;references:ID" json:"lastMessageSeen"`
}
