package gorm

import (
	"gorm.io/gorm"
	"time"
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

type Accounts struct {
	gorm.Model
	Balance float64
}

type Transactions struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

type Employees struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Name        string
	Department  string
	Salary      int32
}

type Books struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Title       string
	Author      string
	Price       float64
}
type User struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Name        string
	Age         int
	Code        uint   `gorm:"uniqueIndex"`
	Posts       []Post `gorm:"foreignKey:Code1;references:Code"`
}

type Post struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Title       string
	Content     string
	Code1       uint
	User        User `gorm:"foreignKey:Code1;references:Code"`
	Comments    []Comment
}

type Comment struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	DeleteFlag  gorm.DeletedAt
	Content     string
	PostID      uint
	UserID      uint
	User        User
	Post        Post
}
