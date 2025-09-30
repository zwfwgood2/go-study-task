package sync

import (
	"fmt"
	"time"
)

func RunTaskExcutor() {
	//定义任务
	tasks := []func(i int) (b int){}
	tasks = append(tasks, func(i int) (b int) {
		s := time.Now()
		time.Sleep(5 * time.Millisecond)
		e := time.Now()
		duration := e.Sub(s)
		b = int(duration.Nanoseconds())
		// 将纳秒转换为秒输出
		seconds := float64(b) / 1000000000.0
		fmt.Printf("任务编号=%v;执行时间=%v秒\n", i, seconds)
		return
	}, func(i int) (b int) {
		s := time.Now()
		time.Sleep(5 * time.Millisecond)
		e := time.Now()
		duration := e.Sub(s)
		b = int(duration.Nanoseconds())
		// 将纳秒转换为秒输出
		seconds := float64(b) / 1000000000.0
		fmt.Printf("任务编号=%v;执行时间=%v秒\n", i, seconds)
		return
	})
	taskExecutor(tasks)
	time.Sleep(20 * time.Millisecond)
}
func taskExecutor(tasks []func(i int) (b int)) {
	for i, task := range tasks {
		go task(i)
	}
}
