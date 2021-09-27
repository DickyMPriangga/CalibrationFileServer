package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load("env/rdsDB.env")

	if err != nil {
		log.Fatal("Error loading rdsDB.env file", err)
	}

	postgreSQLInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err = gorm.Open(postgres.Open(postgreSQLInfo), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&FileData{}, &LogData{})
}

func InsertFileData(file ...FileData) {
	for _, val := range file {
		db.Create(&val)
	}
}

func GetAllFileDataList() []FileData {
	result := []FileData{}
	db.Find(&result)
	return result
}

func GetFileData(filename string) FileData {
	result := &FileData{}
	db.Where("filename = ?", filename).First(result)
	return *result
}

func InsertLogData(log ...LogData) {
	for _, val := range log {
		db.Create(&val)
	}
}

func GetLogDataList(filename string) []LogData {
	result := []LogData{}
	db.Where("filename = ?", filename).Find(&result)
	return result
}

func GetAllLogDataList() []LogData {
	result := []LogData{}
	db.Find(&result)
	return result
}

func DeleteAll() {
	db.Unscoped().Where("1 = 1").Delete(&LogData{})
	db.Unscoped().Where("1 = 1").Delete(&FileData{})
}

func DropTable() {
	db.Exec("DROP TABLE log_data")
}
