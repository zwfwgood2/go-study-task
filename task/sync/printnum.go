package sync

import (
	"fmt"
	"sync"
)

func RunPrintNum() {
	printNum()
}
func printNum() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go func() {
		fmt.Printf("偶数值=")
		for _, num := range nums {
			if num%2 == 0 {
				fmt.Printf("%v,", num)
			}
		}
		fmt.Printf("\n")
		wg.Done()
	}()

	go func() {
		fmt.Printf("奇数值=")
		for _, num := range nums {
			if num%2 != 0 {
				fmt.Printf("%v,", num)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
