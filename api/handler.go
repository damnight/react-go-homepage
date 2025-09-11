package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetWeatherReports(c *fiber.Ctx) error {
	wrh := WeatherReportHourly{}
	reports, err := GetAll[*WeatherReportHourly](wrh.getTablename(), wrh.newEntry)
	if err != nil {
		return err
	}

	return c.JSON(reports)
}

func GetCities(c *fiber.Ctx) error {
	city := City{}
	cities, err := GetAll[*City](city.getTablename(), city.newEntry)
	if err != nil {
		return err
	}

	return c.JSON(cities)
}
