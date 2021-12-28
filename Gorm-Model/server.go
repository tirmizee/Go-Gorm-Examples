package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           // id is primaryKey
	Name         string         // name
	Email        *string        // email
	Age          uint8          // age
	Salary       *uint8         // salary
	Birthday     *time.Time     // birthday
	MemberNumber sql.NullString // member_number
	ActivatedAt  sql.NullTime   // activated_at
	CreatedAt    time.Time      // created_at
	UpdatedAt    time.Time      // updated_at
}

type Product struct {
	ProductId   uint         `gorm:"primaryKey;column:id"`
	ProductCode string       `gorm:"column:code"`
	ProductName string       `gorm:"column:name"`
	CreatedAt   sql.NullTime `gorm:"column:created"`
	UpdatedAt   sql.NullTime `gorm:"column:updated"`
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
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

	// delete
	db.Where("1 = 1").Delete(&User{})
	db.Where("1 = 1").Delete(&Product{})

	// Create
	db.Create(&User{Name: "Praty"})

	email := "zee_pratya@hotmail.com"
	db.Create(&User{Name: "Tirmizee", Email: &email})

}
