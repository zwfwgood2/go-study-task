package blog

import (
	"time"

	"gorm.io/gorm"
)

type BUser struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	//DeleteFlag  gorm.DeletedAt
	Name  string
	Age   int
	Code  uint    `gorm:"uniqueIndex"`
	Posts []BPost `gorm:"foreignKey:Code1;references:Code"`
}

type BPost struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Title       string
	Content     string
	Code1       uint
	User        BUser `gorm:"foreignKey:Code1;references:Code"`
	Comments    []BComment
}

type BComment struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Content     string
	PostID      uint
	UserID      uint
	Post        BPost
}
