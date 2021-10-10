package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	//err := godotenv.Load("dev.env")

	//if err != nil {
	//	log.Fatal("Error loading rdsDB.env file", err)
	//}

	var err error

	postgreSQLInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	//postgreSQLInfo := "host=database-file-server.cfvrwyehu9g1.ap-southeast-1.rds.amazonaws.com user=dickymp password=8jC9emPuLJqdFG6 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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
