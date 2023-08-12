package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/harunalbayrak/go-finance/app/db"
	"github.com/harunalbayrak/go-finance/app/models"
	"github.com/harunalbayrak/go-finance/pkg/configs"
	"github.com/harunalbayrak/go-finance/pkg/routes"
	"github.com/harunalbayrak/go-finance/pkg/utils"
	"github.com/joho/godotenv"
)

func GetStocks() ([]models.Stock, error) {
	stocks := make([]models.Stock, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("uzmanpara.milliyet.com.tr"),
	)

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		code := e.ChildText("b")
		if code != "" {
			stocks = append(stocks, models.Stock{Code: code})
		}
	})

	c.Visit("https://uzmanpara.milliyet.com.tr/canli-borsa/bist-TUM-hisseleri/")

	return stocks, nil
}

func main() {
	godotenv.Load(".env")

	// spy, _ := quote.NewQuoteFromYahoo("ekiz.is", "2023-01-01", "2023-08-18", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

	// 1. Configure the example database connection.
	database := db.CreateDB()

	// AutoMigrate for player table
	database.AutoMigrate(&models.Stock{})

	stocks, _ := GetStocks()

	db.CreateStocks(database, stocks)

	// deneme1()

	getdb, _ := db.GetAllStocks(database)

	fmt.Println(getdb)

	// getChart("EKIZ.IS", "1d", "100d")

	// for _, stock := range stocks {
	// 	fmt.Println("Stock:", stock.Code)
	// }

	// engine := html.New("./views", ".html")
	// engine.Reload(true)       // Optional. Default: false
	// engine.Debug(true)        // Optional. Default: false
	// engine.Layout("embed")    // Optional. Default: "embed"
	// engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	engine := django.New("./views", ".django")

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
