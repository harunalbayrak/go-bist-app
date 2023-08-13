package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harunalbayrak/go-finance/app/db"
)

func GetAllStocks(c *fiber.Ctx) error {
	// Create database connection.
	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	// Return status 500 and database connection error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	database, _ := db.CreateDB()

	// Get all books.
	stocks, err := db.GetAllStocks(database)
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   " were not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(stocks),
		"stocks": stocks,
	})
}
