package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetWeatherReports(c *fiber.Ctx) error {
	var reports []WeatherReportHourly

	reports, err := GetAll(c, "weather_reports")
	if err != nil {
		return err
	}

	return c.JSON(reports)
}
