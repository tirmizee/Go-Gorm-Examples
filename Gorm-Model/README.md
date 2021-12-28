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

```