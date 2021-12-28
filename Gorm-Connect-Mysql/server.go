package main

import (
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
	initGormV1()
}

func initGormV1() {

	var (
		dsn = "user:pass@tcp(0.0.0.0:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println(*db.Config)
	fmt.Println(db.ConnPool)

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100, Create: time.Now()})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

}

func initGormV2() {
	var (
		dsn = "user:pass@tcp(0.0.0.0:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:                 false,
		PrepareStmt:            true,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println(*db.Config)
	fmt.Println(db.ConnPool)

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100, Create: time.Now()})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
}
