package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rubichandrap/hello-go/database"
	"github.com/rubichandrap/hello-go/models"
	"github.com/rubichandrap/hello-go/router"
	"github.com/rubichandrap/hello-go/utils"
)

func init() {
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		log.Fatal("Please provide the DB_DRIVER value on the .env file")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("Please provide the DB_USER value on the .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("Please provide the DB_HOST value on the .env file")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("Please provide the DB_NAME value on the .env file")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("Please provide the DB_PORT value on the .env file")
	}
	if _, err := strconv.ParseUint(dbPort, 10, 32); err != nil {
		log.Fatal("Please provide a valid DB_PORT value on the .env file")
	}
}

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatal(fmt.Errorf("error when connecting to database: %v", err))
	}

	db.AutoMigrate(&models.User{}, &models.ConfigFoundation{})

	fmt.Println("Database migrated")

	app := fiber.New()

	router.Setup(app)

	utils.StartServerWithGracefulShutdown(app)
}
