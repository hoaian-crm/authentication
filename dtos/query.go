package dtos

type Query struct {
	Limit  int `json:"limit" form:"limit" default:"10"`
	Offset int `json:"offset" form:"offset" default:"0"`
}

func (query *Query) SetDefaults() {
	if query.Limit == 0 {
		query.Limit = 10
	}
}
