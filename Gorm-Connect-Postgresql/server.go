package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	postgresConfig := postgres.Config{
		DSN: "host=0.0.0.0 user=usr password=pass dbname=test port=5432 sslmode=disable TimeZone=Asia/Bangkok",
	}

	gormConfig := gorm.Config{}

	db, err := gorm.Open(postgres.New(postgresConfig), &gormConfig)
	if err != nil {
		panic(err)
	}

	if sqlDB, err := db.DB(); err == nil {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, "code = ?", "D42") // find product with code D42

	fmt.Printf("%+v\n", product)

}
