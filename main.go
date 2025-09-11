package main

import (
	"fmt"
	"log"
	"os"

	"github.com/damnight/react-go-homepage/api"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Starting React Go Server!")
	if os.Getenv("ENV") != "production" {
		// Load the .env file if not in production
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	api.InitDB()
	app := fiber.New()

	app.Get("/weatherReports", api.GetWeatherReports)
	app.Get("/cities", api.GetCities)
	// app.Post("/api/weatherReports", createWeatherReport)
	// app.Patch("/api/weatherReports/:id", updateWeatherReport)
	// app.Delete("/api/weatherReports/:id", deleteWeatherReport)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
