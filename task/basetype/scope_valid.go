package basetype

import "fmt"

func RunIsValid(s string) {
	fmt.Printf("字符串:%v是合法的括号:%v\n", s, isValid(s))
}

// 判断某字符串是否为有效的字符
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mappings := map[rune]rune{')': '(', ']': '[', '}': '{'}
	leftChar := []rune{}
	for _, char := range s {
		//left char
		if mappings[char] == 0 {
			leftChar = append(leftChar, char)
		} else {
			if leftChar[len(leftChar)-1] != mappings[char] || len(leftChar) == 0 {
				return false
			} else {
				leftChar = leftChar[:len(leftChar)-1]
			}
		}
	}
	return len(leftChar) == 0
}
