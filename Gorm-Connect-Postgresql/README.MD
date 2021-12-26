- go mod init gorm-postressql
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/postgres

### docker-compose.yaml

```yaml

version: '3.1'

services:

  db:
    image: postgres
    restart: always
    environment:
      POSTGRES_USER: usr
      POSTGRES_PASSWORD: pass
      POSTGRES_DB: test
    ports:
      - '5432:5432'

```


### demo

```go

  postgresConfig := postgres.Config{
		DSN: "host=0.0.0.0 user=usr password=pass dbname=test port=5432 sslmode=disable TimeZone=Asia/Bangkok",
	}

	gormConfig := gorm.Config{}

	db, err := gorm.Open(postgres.New(postgresConfig), &gormConfig)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, "code = ?", "D42") // find product with code D42

	fmt.Printf("%+v\n", product)

```