package gorm

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	name  *string
	age   int
	grade string
}
