package db

import (
	"github.com/harunalbayrak/go-finance/app/models"
	"gorm.io/gorm"
)

func CreateStock(db *gorm.DB, stock *models.Stock) error {
	if err := db.Where("code = ?", stock.Code).Updates(stock).FirstOrCreate(stock).Error; err != nil {
		return nil
	}

	return nil
}

func CreateStocks(db *gorm.DB, stocks []models.Stock) error {
	for _, stock := range stocks {
		CreateStock(db, &stock)
	}

	return nil
}
