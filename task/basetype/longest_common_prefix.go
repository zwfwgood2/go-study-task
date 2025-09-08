package basetype

import (
	"fmt"
)

func RunLongestCommonPrefix(strs []string) {
	fmt.Printf("字符数组:%v中最大公共前缀字符串为:%v\n", strs, longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	row := strs[0]
	for i, c := range row {
		for _, str := range strs {
			if len(str) == i || str[i] != byte(c) {
				return row[:i]
			}
		}
	}
	return row
}
