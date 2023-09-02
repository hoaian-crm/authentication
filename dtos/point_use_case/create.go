package pointusecase_dto

type Create struct {
	Value       int64  `json:"value" binding:"is_number"`
	Description string `json:"description" binding:"is_not_empty"`
}
