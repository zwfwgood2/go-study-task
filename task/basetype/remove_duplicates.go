package basetype

import (
	"fmt"
	"sort"
)

func RunRemoveDuplicates(nums []int) {
	size, result := removeDuplicates(nums)
	fmt.Printf("nums=%v去除重复数字后nums的长度=%v;重复元素数组nums=%v\n", nums, size, result)
}
func removeDuplicates(nums []int) (int, []int) {
	dm := map[int]int{}
	for i, num := range nums {
		if dm[num] == 0 {
			dm[num] = num
		} else {
			nums[i] = 0
		}
	}
	sort.Ints(nums)
	nums = nums[len(nums)-len(dm):]
	return len(nums), nums
}
