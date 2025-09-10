package sync

import "fmt"

func RunAddNum(num int) {
	p := &num
	b := addNum(p)
	fmt.Printf("传入num:%v;b=%v; pointer为:%v;操作后的num=%v", num, b, p, *p)
}
func addNum(num *int) (result int) {
	*num = *num + 1
	return *num
}
