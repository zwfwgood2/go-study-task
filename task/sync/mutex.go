package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func RunMutexTest() {
	fmt.Printf("对num并发增加10000次后num=%v\n", mutexTest())
}
func mutexTest() int32 {
	wg := sync.WaitGroup{}
	wg.Add(10)
	//lock := sync.Mutex{}
	num := atomic.Int32{}
	num.Store(0)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				wg.Done()
				//lock.Unlock()
			}()
			//lock.Lock()
			for j := 0; j < 1000; j++ {
				num.Add(1)
			}
		}()
	}
	wg.Wait()
	return num.Load()
}
