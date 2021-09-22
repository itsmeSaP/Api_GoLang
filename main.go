package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/itsmeSaP/starting-go/api/endpoints"
	"github.com/itsmeSaP/starting-go/pkg/database"
	"github.com/itsmeSaP/starting-go/pkg/models"
)

func main() {
	//Connecting to DATABASE
	db, err := database.ConnectDB()
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Connected to DB")

	//RUNNING AUTOMIGRATIONS
	err = db.AutoMigrate(&models.Posts{})
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Migrations Done")

	//creating the fiber app
	app := fiber.New()
	//Middleware Initialization
	//WHAT IS THIS???
	app.Use(cors.New())
	app.Use(logger.New())

	//Mounting Routes
	endpoints.MountRoutes(app)

	//RUNNING THE SERVER on port 3000
	log.Fatal(app.Listen(":3000"))
}
