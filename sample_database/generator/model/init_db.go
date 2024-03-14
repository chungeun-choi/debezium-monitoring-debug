package model

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// DBType 데이터베이스 유형을 나타내는 타입
type DBType string

const (
	MySQL      DBType = "mysql"
	PostgreSQL DBType = "postgresql"
)

// makeDSN 함수는 주어진 데이터베이스 유형에 따라 DSN을 생성합니다.
func makeDSN(dbType DBType) string {

	// 환경 변수에서 데이터베이스 연결 정보 로드
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		// .env 파일에서 환경 변수 로드
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
		dbUsername = os.Getenv("DB_USERNAME")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 데이터베이스 유형에 따라 DSN 포맷 변경
	var dsn string
	switch dbType {
	case MySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	case PostgreSQL:
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul", dbHost, dbPort, dbUsername, dbPassword, dbName)
	default:
		panic("Unsupported database type")
	}

	return dsn
}

// InitDB 함수는 주어진 데이터베이스 유형에 따라 데이터베이스 연결을 초기화하고 마이그레이션합니다.
func InitDB() *gorm.DB {
	dbType := DBType(os.Getenv("DB_TYPE"))

	dsn := makeDSN(dbType)
	var db *gorm.DB
	var err error

	switch dbType {
	case MySQL:
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case PostgreSQL:
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		panic("Unsupported database type")
	}

	if err != nil {
		fmt.Println(dsn)
		panic("failed to connect database")
	}

	// 테이블 생성
	err = db.AutoMigrate(&Customer{}, &Order{}, &Product{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
