package sync

import "fmt"

func RunSlicedMultiplyTwo(nums []int) {
	fmt.Printf("nums=%v,result=%v\n", nums, slicedMultiplyTwo(nums))
}
func slicedMultiplyTwo(nums []int) (result []int) {
	for i, _ := range nums {
		nums[i] *= 2
	}
	result = nums
	return
}
