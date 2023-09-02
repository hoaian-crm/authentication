package pointuse_dto

type Create struct {
	UserID    int64
	UseCaseId int64 `json:"use_case_id"`
}
