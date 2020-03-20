package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/validations"
	"os"
	"shorten-url/models"
)

type GormManager struct {
	DB *gorm.DB
}

var DBManager GormManager

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load dotenv")
	}
	dbType := os.Getenv("db_type")
	dbSource := os.Getenv("db_source")
	db, err := gorm.Open(dbType, dbSource)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func gormConnectForTest() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func InitDB() {
	DBManager.DB = gormConnect()
	migrate(DBManager.DB)
	validations.RegisterCallbacks(DBManager.DB)
}

func InitDBForTest() {
	DBManager.DB = gormConnectForTest()
	migrate(DBManager.DB)
	validations.RegisterCallbacks(DBManager.DB)
}
