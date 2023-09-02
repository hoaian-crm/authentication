package messages_dto

type Create struct {
	Content   string `json:"content" binding:"is_not_empty"`
	ChannelID int    `json:"channelId,string"`
}
