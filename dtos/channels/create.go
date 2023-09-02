package channels_dto

import "main/models"

type Create struct {
	ChannelID int                `json:"channelId"`
	Name      string             `json:"name"`
	Type      models.ChannelType `json:"channelType"`
}
