package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

type City struct {
	ID        int
	Latitude  float64
	Longitude float64
	Elevation float64
	Name      string
}

type WeatherReportHourly struct {
	ID           int
	City         City
	ForecastDays int

	Temperature              float64
	PercipitationProbability int
	Precipitaion             float64
	CloudCover               int
	WindDirection            int
	UVIndex                  float64
	SurfacePressure          float64
}

func main() {

	fmt.Println("Starting React Go Server!")
	if os.Getenv("ENV") != "production" {
		// Load the .env file if not in production
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	initDB()
	app := fiber.New()

	// app.Get("/api/weatherReports", getWeatherReport)
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

// initDB initializes the SQLite database and creates the todos table if it doesn't exist
func initDB() {
	var err error
	var dbUri = os.Getenv("SQLITE_URI")
	DB, err = sql.Open("sqlite3", dbUri) // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}
	sql_create_cities := `CREATE TABLE IF NOT EXISTS city (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	latitude REAL,
	longitude REAL,
	elevation REAL
	);`

	// SQL statement to create the todos table if it doesn't exist
	sql_create_weatherReports := `
	CREATE TABLE IF NOT EXISTS weatherReports (
	id INTEGER NOT NULL PRIMARY KEY,
	city INTEGER NOT NULL,
	forecast_days INTEGER,
	temperature REAL,
	precipitation_probability INTEGER,
	precipitation REAL,
	cloud_cover INTEGER,
	wind_direction INTEGER,
	uv_index REAL,
	air_pressure REAL,
	FOREIGN KEY(city) REFERENCES city(id)
	);`

	_, err = DB.Exec(sql_create_weatherReports)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sql_create_weatherReports) // Log an error if table creation fails
	}
	_, err = DB.Exec(sql_create_cities)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sql_create_weatherReports) // Log an error if table creation fails
	}
}
