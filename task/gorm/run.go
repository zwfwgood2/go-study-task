package gorm

import (
	"fmt"
)

//students 增删改查

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
