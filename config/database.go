package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDataBase() {
	dsn := EnvirontmentVariables.GormDSN

	localDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		println("Error when try to connecting with database: ", err)
		return
	}
	println("Database connected !")
	Db = localDB
}
