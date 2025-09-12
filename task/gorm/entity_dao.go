package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		db.AutoMigrate(&Students{})
	}
}
func addStudent(student *Students) bool {
	tx := db.Debug().Create(student)
	return tx.Error == nil
}

func queryStudent(student *Students) (result []Students, e error) {
	fmt.Printf("rsult=%v\n", result)
	db.Debug().Model(student).Where("age>?", student.Age).Find(&result)
	return result, db.Error
}

func deleteStudent(student *Students) bool {
	tx := db.Debug().Unscoped().Where("age<?", student.Age).Delete(student)
	return tx.Error == nil
}

func updateStudent(student *Students) bool {
	tx := db.Debug().Model(student).Update("grade", student.Grade)
	return tx.Error == nil
}
