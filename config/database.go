package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func ConnectDataBase() {
	dsn := EnvirontmentVariables.GormDSN

	nameConfig := schema.NamingStrategy{}

	localDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: nameConfig,
	})

	if err != nil {
		println("Error when try to connecting with database: ", err)
		return
	}
	println("Database connected !")
	Db = localDB
}
