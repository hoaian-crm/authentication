package models

type PointDetail struct {
	UserId    string `json:"user_id"`
	Amount    int    `json:"amount"`
	UseCaseId int    `json:"use_case_id"`
}
