package models

type Permission struct {
	BaseModel
	Name        string `json:"name" binding:"is_not_empty"`
	Description string `json:"description" binding:"is_not_empty"`
	Policy      string `json:"policy" binding:"is_not_empty"`
}
