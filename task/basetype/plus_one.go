package basetype

import (
	"fmt"
	"strconv"
)

func RunPlusOne(digits []int64) {
	fmt.Printf("大整数数组加一后的数组为:%v\n", plusOne(digits))
}
func plusOne(digits []int64) []int64 {
	digitsStr := ""
	for _, digit := range digits {
		digitsStr += strconv.Itoa(int(digit))
	}
	num, _ := strconv.Atoi(digitsStr)
	num += 1
	outDigits := []int64{}
	for _, s := range strconv.Itoa(num) {
		v, _ := strconv.Atoi(string(byte(s)))
		outDigits = append(outDigits, int64(v))
	}
	return outDigits

}
