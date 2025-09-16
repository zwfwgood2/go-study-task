package gorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func Run() {
	//OpStudents()
	//insterAccounts([]Accounts{{Balance: 100.0}, {Balance: 100.0}}, db)
	//err := trans(Transactions{FromAccountId: 5, ToAccountId: 6, Amount: 10})
	//fmt.Printf("error=%v\n", err)

	//insterBooks([]Books{{Title: "book1", Author: "zw", Price: 50}, {Title: "book1", Author: "zw", Price: 50}}, db)
	//insterEmployees([]Employees{{Name: "zw", Department: "技术部", Salary: 500001}, {Name: "zw1", Department: "技术部", Salary: 40001}, {Name: "zw2", Department: "技术部1", Salary: 2}}, db)
	//queryBookAndEmployees()
	queryUsersByCode()
}

// students 增删改查
func OpStudents() {
	addStudent(&Students{Name: "zhangsan", Age: 10, Grade: "4年级"})

	updateStudent(&Students{ID: 1, Name: "lisi", Grade: "六年级"})
	student, err := queryStudent(&Students{Age: 9})
	if err == nil {
		fmt.Printf("query result=%v\n", student)
		//fmt.Printf("query reusult=%v\n", student[0].Age)
	}
	deleteStudent(&Students{Name: "zhangsan", Age: 16})
}

func trans(trans Transactions) error {
	if trans.FromAccountId == 0 {
		return errors.New("fromAccountId is Required")
	}
	if trans.ToAccountId == 0 {
		return errors.New("toAccountId is Required")
	}
	if trans.Amount == 0 || trans.Amount < 0 {
		return errors.New("amount need gt 0 or not equal 0")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		//检验账户余额
		account, err := selectAccountByAccountId(trans.FromAccountId, tx)
		if err != nil {
			return err
		}
		if account.Balance < trans.Amount {
			return errors.New("balance deficiency")
		}
		err = updateAccount(Accounts{Model: gorm.Model{ID: trans.FromAccountId}, Balance: -trans.Amount}, tx)
		if err != nil {
			return err
		}
		err = updateAccount(Accounts{Model: gorm.Model{ID: trans.ToAccountId}, Balance: trans.Amount}, tx)
		if err != nil {
			return err
		}
		record, err := insterTransactionRecord(trans, tx)
		if err == nil {
			return errors.New("transaction rollback")
		}
		fmt.Printf("交易成功，交易记录=%v\n", record)
		return nil
	})
}

func queryBookAndEmployees() {
	//query books
	result := queryBooks(Books{Price: 50})
	fmt.Printf("books query result=%v\n", result)

	//query gourp by department having salary=max(salary)
	mr := queryEmployees(map[string]interface{}{"department": "技术部"})
	fmt.Printf("employees query result=%v\n", mr)

	mrp, count := queryEmployeesByPage(map[string]interface{}{"department": "技术部"}, map[string]interface{}{"page": 0, "pageSize": 10})
	fmt.Printf("employees query result=%v,总条数=%v,len=%v,cap=%v\n", mrp, count, len(mrp), cap(mrp))
}
func queryUsersByCode() {

	//saveUsers(User{Code: 1, Name: "lisi", Age: 10,
	//	Posts: []Post{{Title: "post1", Content: "content1",
	//		Comments: []Comment{{Content: "comment1"}, {Content: "comment2"}}}, {Title: "post2", Content: "content2",
	//		Comments: []Comment{{Content: "comment3"}, {Content: "comment4"}}}}})

	result, _ := selectUserByCode(1)
	fmt.Printf("users query result=%v\n", result)

	post1 := selectPostsByMaxComments()
	fmt.Printf("posts query result=%v\n", post1)
}
