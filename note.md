```
docker run -d --name server-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass-server-mysql -e MYSQL_DATABASE=go_fast_food_db mysql:8.0
docker run --name server-postgres -e POSTGRES_PASSWORD=pass-server-postgres -p 5432:5432 -d postgres
```

```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

Create Go.mod
```
go mod init example.com/m
```

Create Go.sum
```
go mod tidy
```

Create .ENV
```
go get github.com/joho/godotenv
```