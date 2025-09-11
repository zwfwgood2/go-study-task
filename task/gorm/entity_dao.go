package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
}
func addStudent(student *Students) bool {
	tx := db.Create(student)
	return tx.Error == nil
}

func queryStudent(student *Students) (result []Students) {
	db.Model(student).Where("age>?", 18).Find(result)
	return
}

func deleteStudent(student *Students) bool {
	tx := db.Where("age<?", 15).Delete(student)
	return tx.Error == nil
}

func updateStudent(student *Students) bool {
	tx := db.Model(student).Update("grade", "四年级")
	return tx.Error == nil
}
