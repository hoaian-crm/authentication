package config

import (
	"strings"
	"unicode"

	"gorm.io/gorm/schema"
)

type NameConvention struct {
	*schema.NamingStrategy
}

// type Namer interface {
// 	TableName(table string) string
// 	SchemaName(table string) string
// 	ColumnName(table, column string) string
// 	JoinTableName(joinTable string) string
// 	RelationshipFKName(Relationship) string
// 	CheckerName(table, column string) string
// 	IndexName(table, column string) string
// }

func (n NameConvention) TableName(table string) string {
	return strings.ToLower(table)
}

func (n NameConvention) ColumnName(table, column string) string {
	if len(column) == 0 {
		return column
	}

	r := []rune(column)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
