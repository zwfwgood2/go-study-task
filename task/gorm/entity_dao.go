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
		db.AutoMigrate(&Students{}, &Accounts{}, &Transactions{}, &Books{}, &Employees{})
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

func selectAccountByAccountId(accountId uint, tx *gorm.DB) (accounts Accounts, error error) {
	accounts.ID = accountId
	tx.Debug().Find(&accounts)
	error = db.Error
	return
}
func insterAccounts(accounts []Accounts, tx *gorm.DB) ([]Accounts, error) {
	tx.Debug().Create(&accounts)
	return accounts, db.Error
}
func updateAccount(accounts Accounts, tx *gorm.DB) error {
	tx.Debug().Model(&accounts).Update("balance", gorm.Expr("balance+?", accounts.Balance))
	return tx.Error
}

func insterTransactionRecord(transactions Transactions, tx *gorm.DB) (Transactions, error) {
	tx.Debug().Create(&transactions)
	return transactions, db.Error
}

func queryBooks(book Books) (result []Books) {
	db.Debug().Model(book).Select("title").Where(book).Find(&result)
	return result
}

func queryEmployees(employees map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, 3)
	db.Debug().Model(Employees{}).Select("department,max(salary) as salary").Where(employees).Group("department").Find(result)
	return result
}

func queryEmployeesByPage(employees map[string]interface{}, page map[string]interface{}) ([]map[string]interface{}, int64) {
	result := []map[string]interface{}{}
	page1 := page["page"].(int)
	pageSize1 := page["pageSize"].(int)
	var count int64
	db.Debug().Model(Employees{}).Select("department,salary").Where(employees).Count(&count).Limit(pageSize1).Offset(page1).Find(&result)
	return result, count
}

func insterBooks(books []Books, tx *gorm.DB) ([]Books, error) {
	tx.Debug().Create(&books)
	return books, db.Error
}

func insterEmployees(employees []Employees, tx *gorm.DB) ([]Employees, error) {
	tx.Debug().Create(&employees)
	return employees, db.Error
}
