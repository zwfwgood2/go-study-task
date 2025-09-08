package basetype

import (
	"fmt"
	"strconv"
)

func RunIsPalindrame(num int32) {
	palindrome, err := IsPalindrome(num)
	if err != nil {
		return
	}
	fmt.Printf("数字num:%v是否为回数%v\n", num, palindrome)
}
func IsPalindrome(num int32) (bool, error) {
	//TODO 检验数值合法性
	if num < 0 {
		return false, nil
	}
	rawValue := strconv.Itoa(int(num))
	runes := []rune(rawValue)
	reverseRunes := make([]rune, len(runes))
	for i, _ := range runes {
		reverseRunes[i] = runes[len(runes)-1-i]
	}
	reverseValue := string(reverseRunes)
	fmt.Printf("runes=%v;reverseRunes=%v;rawValue=%v;reverseValue=%v\n", runes, reverseRunes, rawValue, reverseValue)
	if rawValue == reverseValue {
		return true, nil
	}
	return false, nil
}
