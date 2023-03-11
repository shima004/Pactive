package postgres

import (
	"fmt"
	"os"
	"time"

	"github.com/shima004/pactive/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresHandler struct {
	Conn *gorm.DB
}

var connectionPool *gorm.DB

func GetDB() (*gorm.DB, error) {
	if connectionPool == nil {
		return nil, fmt.Errorf("db is nil")
	}
	return connectionPool, nil
}

func InitDB() *gorm.DB {
	connection := sqlConnect()
	connection.AutoMigrate(&model.User{})
	connectionPool = connection
	return connectionPool
}

func sqlConnect() (database *gorm.DB) {
	HOST := os.Getenv("POSTGRES_HOST")
	PORT := os.Getenv("POSTGRES_PORT")
	USER := os.Getenv("POSTGRES_USER")
	PASS := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DB")
	TIME_ZONE := os.Getenv("POSTGRES_TIME_ZONE")
	SSL_MODE := os.Getenv("POSTGRES_SSL_MODE")

	CONNECT := "host=" + HOST + " user=" + USER + " password=" + PASS + " dbname=" + DBNAME + " port=" + PORT + " TimeZone=" + TIME_ZONE + " sslmode=" + SSL_MODE

	count := 0
	db, err := gorm.Open(postgres.Open(CONNECT), &gorm.Config{})
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("error")
				panic(err)
			}
			db, err = gorm.Open(postgres.Open(CONNECT), &gorm.Config{})
		}
	}

	DB, err := db.DB()
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour)
	return db
}
