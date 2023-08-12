package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDB() *gorm.DB {
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_CONNECTION") + ")/test?charset=utf8mb4"
	fmt.Println("Dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: log.Default.LogMode(log.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
