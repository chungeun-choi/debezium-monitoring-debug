package main

import (
	"debezium-test-data-generator/generate"
	"debezium-test-data-generator/model"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func initApp(db *gorm.DB) {
	generate.GenerateCustomer(db, 100)
	generate.GenerateProduct(db, 2000)
}

func setLog() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stdout)
}

func Run1Hour(db *gorm.DB) {
	// 6시간마다 반복 실행되는 Ticker 생성
	ticker := time.NewTicker(1 * time.Hour) // 올바른 시간 간격으로 수정

	defer ticker.Stop()
	for range ticker.C {
		generate.GenerateCustomer(db, 2)
		generate.GenerateProduct(db, 15)
		log.Println("Generated additional customers and products")
	}
}

func Run1Minute(db *gorm.DB) {
	// 1시간마다 반복 실행되는 Ticker 생성
	ticker := time.NewTicker(1 * time.Minute) // 올바른 시간 간격으로 수정

	defer ticker.Stop()
	for range ticker.C {
		var customers, products []uint64

		switch db_type := model.DBType(os.Getenv("DB_TYPE")); db_type {
		case model.MySQL:
			customers = generate.RandomCustomerIdMySQL(db)
			products = generate.RandomProductIdMySQL(db)
		case model.PostgreSQL:
			customers = generate.RandomCustomerIdPostgreSQL(db)
			products = generate.RandomProductIdPostgreSQL(db)
		}

		generate.GenerateOrders(db, customers, products, 10000)
		log.Println("Generated new orders")
	}
}

func main() {
	db := model.InitDB()
	initApp(db)
	setLog()
	log.Println("Data generation and saving to DB completed.")

	// 고루틴을 사용하여 동시에 두 함수 실행
	go Run1Hour(db)
	go Run1Minute(db)

	// 메인 함수가 종료되지 않도록 무한 루프
	select {}
}
