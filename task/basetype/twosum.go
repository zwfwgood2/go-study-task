package basetype

import (
	"fmt"
)

func RunTowSum(nums []int, target int) {
	fmt.Printf("输入数字数组:%v中等于target:%v的两个数字在nums中的索引是:%v\n", nums, target, twoSum(nums, target))
}
func twoSum(nums []int, target int) (index []int) {
	//存储匹配到的数值和数值对应的索引
	nim := map[int]int{}
	for i, n := range nums {
		if v, ok := nim[target-n]; !ok {
			nim[n] = i
		} else {
			return append(index, v, i)
		}
	}
	return
}
