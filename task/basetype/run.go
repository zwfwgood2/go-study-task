package basetype

func Run() {
	RunSingleNumber(BuildNums())
	RunIsPalindrame(101)
	RunIsValid("{()}")
	RunLongestCommonPrefix([]string{"bahao", "ni", "nihaoa"})
	RunPlusOne([]int64{1, 9, 9})
	RunRemoveDuplicates([]int{1, 1, 3, 5, 5, 6, 7})
	RunMerge([][]int{{1, 3}, {2, 4}, {1, 5}})
	RunTowSum([]int{1, 3, 5}, 5)
}
