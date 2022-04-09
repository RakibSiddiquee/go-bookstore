Install the following packages:
1. go get -u gorm.io/gorm
2. go get gorm.io/driver/mysql
3. go get github.com/gorilla/mux

Next, create a mysql database name "book_store"

Then run the following commands:
cd cmd/main
go build
go run main.go

After that, test using postman