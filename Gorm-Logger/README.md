- go get -u gorm.io/gorm
- go get -u gorm.io/driver/mysql

### docker-compose.yaml

```yaml

version: '3.1'

services:
  db:
    image: mysql
    restart: always
    ports:
      - '3306:3306'
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: example
      MYSQL_USER: user
      MYSQL_PASSWORD: pass

```

### demo

```go

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type CustomLogger struct {
	logger.Interface
}

func (c CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Println(sql)
}

func main() {

	var (
		dsn = "user:pass@tcp(0.0.0.0:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	)

	gormConfig := gorm.Config{
		// Logger : logger.Default.LogMode(logger.Info),
		Logger: &CustomLogger{},
	}

	db, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	db.Where("1 = 1").Delete(&Product{})

	// Create record
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

}

```