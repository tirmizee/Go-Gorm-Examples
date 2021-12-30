- go mod init gorm-error
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

```

        record not found
        record not found
