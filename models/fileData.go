package models

import (
	"time"

	"gorm.io/gorm"
)

type FileData struct {
	gorm.Model
	FileName   string    `gorm:"column:filename"`
	FileAuthor string    `gorm:"column:fileauthor"`
	FileURL    string    `gorm:"column:fileurl"`
	Expire     time.Time `gorm:"column:expire;type:date"`
}
