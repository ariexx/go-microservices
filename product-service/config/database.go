package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"product-service/pkg/model"
	"time"
)

func InitDatabase() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
	}

	//call auto migrate
	err = db.AutoMigrate(
		&model.Product{},
		&model.ProductDetail{},
	)

	if err != nil {
		log.Fatalf("Error while migrating database %s", err)
	}

	//config database
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
