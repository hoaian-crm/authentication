package dtos

type Query struct {
	Limit        int      `json:"limit" form:"limit"`
	Offset       int      `json:"offset" form:"offset"`
	SearchFields []string `json:"search_fields" form:"search_fields"`
	Keywords     []string `json:"keywords" form:"keywords"`
}
