package base

import (
	"fmt"
	"main/config"
	"main/dtos"

	"gorm.io/gorm"
)

type Repository[Model any] struct {
	model Model 
}

func (repository Repository[Model]) CreateOne(record *Model) *gorm.DB {
	result := config.Db.Create(&record)
	return result
}

func (repository Repository[Model]) FindOne(filter *Model) (Model, error) {
	var record Model
	result := config.Db.Model(&filter).Where(&filter).First(&record)
	return record, result.Error
}

func (respoitory Repository[Model]) FindAll(filter *Model, query dtos.Query) ([]Model, int64) {
	var records []Model

	result := config.Db.Model(&filter).Where(&filter).Offset(query.Offset).Limit(query.Limit).Find(&records)

	return records, result.RowsAffected
}

func (repository Repository[Model]) GetAll(filter *Model) []Model {

	var records []Model

	config.Db.Model(&filter).Where(&filter).Find(&records)

	return records
}

func (repository Repository[Model]) Updates(filter interface{}, value interface{}) {

	result := config.Db.Model(&repository.model).Where(&filter).Updates(value)
	fmt.Printf("result: %v\n", result)
}
