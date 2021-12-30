- go mod init gorm-transaction-manual
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


```
