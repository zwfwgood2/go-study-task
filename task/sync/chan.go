package sync

import (
	"fmt"
	"time"
)

func RunchanOut() {
	chanOut()
}
func chanOut() {
	//c := make(chan int)
	c := make(chan int, 1)
	go func() {
		for i := 1; i < 11; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for {
			select {
			case o, ok := <-c:
				if !ok {
					return
				}
				time.Sleep(1 * time.Second)
				fmt.Printf("收到数字%v,", o)
			}
		}
	}()

	time.Sleep(12 * time.Second)
}
