package pointusecase_dto

type Uri struct {
  ID int64 `uri:"id" binding:"is_number"`
}

type Data struct {
	Description string `json:"description"`
	Value       int    `json:"value"`
}
