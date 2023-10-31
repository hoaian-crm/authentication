package dtos

type Query struct {
	Limit     int    `json:"limit" form:"limit"`
	Offset    int    `json:"offset" form:"offset"`
	Order     string `json:"order" form:"order"`
	Direction string `json:"direction" form:"direction"`
}

func (query *Query) SetDefaults() {
	if query.Limit == 0 {
		query.Limit = 10
	}
}
