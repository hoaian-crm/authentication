package dtos

type Query struct {
  Limit int `json:"limit" form:"limit"`
  Offset int `json:"offset" form:"offset" default=10`
}
