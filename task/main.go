package main

import "github.com/zwfwgood2/go-study-task/task/basetype"

func main() {
	basetype.RunSingleNumber(basetype.BuildNums())
	basetype.RunIsPalindrame(101)
	basetype.RunIsValid("{()}")
	basetype.RunLongestCommonPrefix([]string{"bahao", "ni", "nihaoa"})
	basetype.RunPlusOne([]int64{1, 9, 9})
}
