package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code   string
	Price  uint
	Create time.Time
}

func main() {

	var (
		dsn = "user:pass@tcp(0.0.0.0:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Clean data
	db.Where("1 = 1").Delete(&Product{})

	// Create record
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	if err := db.Where("code = ?", "D41").First(&product).Error; err != nil {
		// error handling...
		fmt.Println(err)
	}

	product = Product{}
	err = db.Where("code = ?", "D41").First(&product).Error

	// error handling...
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrInvalidDB) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrEmptySlice) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrInvalidDB) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrInvalidData) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrInvalidTransaction) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrInvalidValue) {
		fmt.Println(err.Error())
	} else if errors.Is(err, gorm.ErrModelValueRequired) {
		fmt.Println(err.Error())
	}

}
