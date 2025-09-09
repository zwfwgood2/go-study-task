package main

import "github.com/zwfwgood2/go-study-task/task/basetype"

func main() {
	basetype.RunSingleNumber(basetype.BuildNums())
	basetype.RunIsPalindrame(101)
	basetype.RunIsValid("{()}")
	basetype.RunLongestCommonPrefix([]string{"bahao", "ni", "nihaoa"})
	basetype.RunPlusOne([]int64{1, 9, 9})
	basetype.RunRemoveDuplicates([]int{1, 1, 3, 5, 5, 6, 7})
}
