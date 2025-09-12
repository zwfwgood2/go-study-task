package gorm

import (
	"time"

	"gorm.io/gorm"
)

type Students struct {
	//gorm.Model
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Name        string
	Age         int
	Grade       string
}
