package main

import (
	"fmt"
	"github.com/zwfwgood2/go-study-task/task/basetype"
)

func main() {
	//输入切片
	nums := make([]int, 21)
	//nums := []int{}
	for i := 0; i < 21; i++ {
		nums[i] = i
		if i > 9 {
			nums[i] = i - 10
		}
	}

	fmt.Printf("此切片%v中只出现一次的数字是：%v", nums, basetype.SingleNumber(nums))
}
