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