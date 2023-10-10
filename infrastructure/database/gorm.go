package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fhrlzmn/hacktiv8-go/assignment-2/config"
	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	Connect()
	Migration()
}

func Connect() {
	dbConfig := config.GetConfig().Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func Migration() {
	db.AutoMigrate(entity.Order{}, entity.Item{})
}

func GetInstance() *gorm.DB {
	if db == nil {
		log.Panic("Database instance is not initialized")
	}

	return db
}
