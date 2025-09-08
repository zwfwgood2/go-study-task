package basetype

import (
	"errors"
	"fmt"
)

func RunSingleNumber(nums []int) {
	fmt.Printf("此切片%v中只出现一次的数字是：%v\n", nums, SingleNumber(nums))
}

func BuildNums() []int {
	//输入切片
	nums := make([]int, 21)
	//nums := []int{}
	for i := 0; i < 21; i++ {
		nums[i] = i
		if i > 9 {
			nums[i] = i - 10
		}
	}
	return nums
}

/**
 * 136. 只出现一次的数字
 */
func SingleNumber(nums []int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic类型是：%v,字面值是：%v\n", r.(error), r)
		}
	}()
	_, err := validate(nums)
	if err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			panic(err)
		}
	}
	singleNum := 0
	numsGroup := make(map[int]int, 10)
	//按照要统计的数字分组
	for _, value := range nums {
		//不存在此key
		if _, exist := numsGroup[value]; !exist {
			numsGroup[value] = 1
		} else {
			numsGroup[value] += 1
		}
	}
	//找到组内数量为1的key
	for key, value := range numsGroup {
		if value == 1 {
			singleNum = key
			return singleNum
		}
	}
	defer func() {
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	return singleNum
}

func validate(nums []int) (int, error) {
	if nums == nil || len(nums) < 1 || len(nums) > 21 {
		return 0, errors.ErrUnsupported
	}
	return len(nums), nil
}
