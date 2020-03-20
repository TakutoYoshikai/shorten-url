package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"shorten-url/models"
)

func initDB() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load dotenv")
	}
	db, err := gorm.Open("sqlite3", os.Getenv("db_source"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})
}
func main() {
	initDB()
}
