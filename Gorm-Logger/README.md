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

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/go/pkg/mod/gorm.io/driver/mysql@v1.2.2/migrator.go:234
        [1.625ms] [rows:-] SELECT DATABASE()

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/go/pkg/mod/gorm.io/driver/mysql@v1.2.2/migrator.go:237
        [4.358ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'example%' ORDER BY SCHEMA_NAME='example' DESC limit 1

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:31
        [4.598ms] [rows:-] SELECT count(*) FROM information_schema.statistics WHERE table_schema = 'example' AND table_name = 'products' AND index_name = 'idx_products_deleted_at'

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:31
        [32.461ms] [rows:0] CREATE INDEX `idx_products_deleted_at` ON `products`(`deleted_at`)

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:33
        [7.522ms] [rows:0] UPDATE `products` SET `deleted_at`='2021-12-30 01:23:29.457' WHERE 1 = 1 AND `products`.`deleted_at` IS NULL

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:36
        [11.293ms] [rows:1] INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2021-12-30 01:23:29.465','2021-12-30 01:23:29.465',NULL,'D42',100)

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:40
        [3.657ms] [rows:1] SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:41
        [4.262ms] [rows:1] SELECT * FROM `products` WHERE code = 'D42' AND `products`.`deleted_at` IS NULL AND `products`.`id` = 1 ORDER BY `products`.`id` LIMIT 1

        2021/12/30 01:23:29 /Users/pratya.yeekhaday/Desktop/GO-Playground/Go-Gorm-Examples/Gorm-Logger/server.go:44
        [12.663ms] [rows:1] UPDATE `products` SET `price`=200,`updated_at`='2021-12-30 01:23:29.485' WHERE `id` = 1