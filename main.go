package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/harunalbayrak/go-finance/app/db"
	"github.com/harunalbayrak/go-finance/app/models"
	"github.com/harunalbayrak/go-finance/pkg/configs"
	"github.com/harunalbayrak/go-finance/pkg/routes"
	"github.com/harunalbayrak/go-finance/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("env loading error", err)
	}

	// spy, _ := quote.NewQuoteFromYahoo("ekiz.is", "2023-01-01", "2023-08-18", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

	// 1. Configure the example database connection.
	database := db.CreateDB()

	// AutoMigrate for player table
	database.AutoMigrate(&models.Stock{})

	stocks, _ := utils.FindAllStocks()

	db.CreateStocks(database, stocks)

	// deneme1()

	getdb, _ := db.GetAllStocks(database)

	for _, stock := range getdb {
		yahooChart, _ := stock.GetYahooChart("1d", "5d")
		fmt.Println("Stock:", yahooChart.Chart.Result[0].Meta.Symbol, "=", yahooChart.Chart.Result[0].Meta.RegularMarketPrice)
	}

	fmt.Println(getdb)

	// getChart("EKIZ.IS", "1d", "100d")

	// for _, stock := range stocks {
	// 	fmt.Println("Stock:", stock.Code)
	// }

	engine := django.New("./web/views", ".django")

	// AddFunc adds a function to the template's global function map.
	engine.AddFunc("greet", func(name string) string {
		return "Hello, " + name + "!"
	})

	config := configs.FiberConfig(engine)

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(config)

	routes.PublicRoutes(app) // Register a public routes for app.

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"stocks": []models.Stock{{Code: "asd2"}, {Code: "asd"}},
		})
	})

	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
