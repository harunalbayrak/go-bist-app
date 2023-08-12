package db

import (
	"github.com/harunalbayrak/go-finance/app/models"
	"gorm.io/gorm"
)

func GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	stocks := []models.Stock{}
	db.Find(&stocks)

	return stocks, nil
}
