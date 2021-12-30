package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
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
	db.Create(&Product{Code: "D42", Price: 100, Create: time.Now()})

	var product Product
	db.First(&product, "code = ?", "D42")

	// begin a transaction
	tx := db.Begin()

	tx.Model(&product).Where("code = ?", "D42").Update("price", 80000)

	isSomtingErr := true
	if isSomtingErr {
		// rollback the transaction in case of error
		tx.Rollback()
	} else {
		// commit the transaction in case all process success
		tx.Commit()
	}

}
