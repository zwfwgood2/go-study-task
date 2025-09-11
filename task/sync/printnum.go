package sync

import (
	"fmt"
	"sync"
)

func RunPrintNum() {
	//printNum()
	syncPrintNum()
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

func syncPrintNum() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oc := make(chan int)
	ec := make(chan int)
	go func() {
		//fmt.Printf("偶数值=")
		for _, num := range nums {
			if num%2 == 0 {
				<-oc
				fmt.Printf("%v,", num)
				//if i != len(nums)-1 {
				//	ec <- num
				//}
				ec <- num
			}
		}
		fmt.Printf("\n")
		defer wg.Done()
	}()

	go func() {
		//fmt.Printf("奇数值=")
		for _, num := range nums {
			if num%2 != 0 {
				<-ec
				fmt.Printf("%v,", num)
				oc <- num
			}
		}
		<-ec
		defer wg.Done()
	}()
	ec <- 1
	wg.Wait()
}
