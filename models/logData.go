package models

import (
	"time"

	"gorm.io/gorm"
)

type LogData struct {
	gorm.Model
	FileName string    `gorm:"column:filename"`
	Action   string    `gorm:"column:action"`
	User     string    `gorm:"column:user"`
	Time     time.Time `gorm:"column:time"`
}
