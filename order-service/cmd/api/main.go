package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"order-service/data"
	"order-service/data/seeds"
	"os"
	"time"
)

type Config struct {
	db *gorm.DB
}

const port = ":80"

func main() {
	db := openDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error while connecting to database %s", err)
		}

		sqlDB.Close()
	}()

	//call data model
	orderRepo := data.NewOrderConfig(db)
	_ = orderRepo.AutoMigrate()

	err := seeds.Run(db, seeds.All())
	if err != nil {
		log.Fatalf("Error while running seeds %s", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*", // "Content-Type,Authorization,Origin,X-Requested-With,Accept",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//inject routes
	routes(app, db)

	listen := app.Listen(port)
	if listen != nil {
		panic(listen)
	}
}

func openDB() *gorm.DB {

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
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
