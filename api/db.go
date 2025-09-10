package api

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type City struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Elevation float64 `db:"elevation"`
}

type WeatherReportHourly struct {
	ID                       int     `db:"id"`
	Time                     string  `db:"time"`
	CityID                   int     `db:"city"`
	ForecastDays             int     `db:"forecast_days"`
	Temperature              float64 `db:"temperature"`
	PercipitationProbability int     `db:"precipitation_probability"`
	Precipitation            float64 `db:"precipitation"`
	CloudCover               int     `db:"cloud_cover"`
	WindDirection            int     `db:"wind_direction"`
	UVIndex                  float64 `db:"uv_index"`
	SurfacePressure          float64 `db:"surface_pressure"`
}

type Table interface {
	stringifyRowValues() string
	stringifyColumns() string
	getTablename() string
}

func (c *City) stringifyRowValues() string {
	return fmt.Sprintf("%v, %s, %f, %f, %f", c.ID, c.Name, c.Latitude, c.Longitude, c.Elevation)
}

func (c *City) stringifyColumns() string {
	return "id, name, latitude, longitude, elevation"
}

func (c *City) getTablename() string {
	return "cities"
}

func (w *WeatherReportHourly) stringifyRowValues() string {
	return fmt.Sprintf("%v, %s, %v, %v, %f, %v, %f, %v, %v, %f, %f", w.ID, w.Time, w.CityID, w.ForecastDays, w.Temperature, w.PercipitationProbability, w.Precipitation, w.CloudCover, w.WindDirection, w.UVIndex, w.SurfacePressure)
}

func (w *WeatherReportHourly) stringifyColumns() string {
	return "id, time, city, forecast_days, temperature, precipitation_probability, cloud_cover, wind_direction, uv_index, surface_pressure"
}

func (w *WeatherReportHourly) getTablename() string {
	return "weather_reports"
}

// InitDB initializes the SQLite database and creates the todos table if it doesn't exist
func InitDB() {
	var err error
	var dbUri = os.Getenv("SQLITE_URI")
	db, err = sqlx.Open("sqlite3", dbUri) // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}
	sql_create_cities := `CREATE TABLE IF NOT EXISTS cities (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	latitude REAL,
	longitude REAL,
	elevation REAL
	);`

	// SQL statement to create the todos table if it doesn't exist
	sql_create_weatherReports := `
	CREATE TABLE IF NOT EXISTS weather_reports (
	id INTEGER NOT NULL PRIMARY KEY,
	time TEXT,
	city INTEGER NOT NULL,
	forecast_days INTEGER,
	temperature REAL,
	precipitation_probability INTEGER,
	precipitation REAL,
	cloud_cover INTEGER,
	wind_direction INTEGER,
	uv_index REAL,
	surface_pressure REAL,
	FOREIGN KEY(city) REFERENCES city(id)
	);`

	_, err = db.Exec(sql_create_weatherReports)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sql_create_weatherReports) // Log an error if table creation fails
	}
	_, err = db.Exec(sql_create_cities)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sql_create_weatherReports) // Log an error if table creation fails
	}

	if len(os.Args) > 2 && os.Args[2] == "sample" {
		if err := AddSampleData(); err != nil {
			log.Fatalf("Error adding sample data: %v", err)
		}
	}
}

func AddSampleData() error {

	city := &City{ID: 1, Name: "Oberteuringen", Latitude: 47.7241, Longitude: 9.4698, Elevation: 450.0}
	report1 := &WeatherReportHourly{ID: 1, Time: "2025-09-10 06:00", CityID: city.ID, ForecastDays: 0, Temperature: 14.3, PercipitationProbability: 23, Precipitation: 0.0, CloudCover: 89, WindDirection: 135, UVIndex: 2.1, SurfacePressure: 950.3}

	_, err := InsertRow(city)
	if err != nil {
		return err
	}
	_, err = InsertRow(report1)
	if err != nil {
		return err
	}

	return nil
}

func NewTableEntry(tablename string) {
	switch tablename {
	case "cities":
		return &City{}
	case "weather_reports":
		return &WeatherReportHourly{}
	default:
		return fmt.Errorf("Can´t create table entry. Table does not exist.")
	}

}

// CRUD
func GetAll(tablename string) ([]Table, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablename)
	rows, err := db.Query(query)
	if err != nil {
		return []Table{}, err
	}
	defer rows.Close()

	data := []Table{}

	for rows.Next() {
		item := NewTableEntry(tablename)
		err = rows.StructScan(&item)
		if err != nil {
			return []Table{}, err
		}
		data = append(data, item)
	}

	return data, nil

}

func InsertRow(data Table) (int64, error) {

	tablename := data.getTablename()
	columnstring := data.stringifyColumns()
	datastring := data.stringifyRowValues()

	stmt := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tablename, columnstring, datastring)

	result, err := db.Exec(stmt)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()

}
