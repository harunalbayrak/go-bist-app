package utils

import (
	"github.com/gocolly/colly"
	"github.com/harunalbayrak/go-finance/app/models"
)

func FindAllStocks() ([]models.Stock, error) {
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
