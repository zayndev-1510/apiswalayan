package setup

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbconn *gorm.DB

func ConnectDB() *gorm.DB {

	errenv := godotenv.Load()
	if errenv != nil {
		fmt.Println("File env not found")
	}
	var err error
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT:=os.Getenv("DB_PORT")
	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":"+DB_PORT+")/" + DB_NAME + "?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		fmt.Println("Koneksi Gagal", &err)
	}

	dbconn = database
	return dbconn
}
