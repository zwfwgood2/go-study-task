package main

import (
	"time"

	"github.com/zwfwgood2/go-study-task/task/sync"
)

func main() {
	//basetype.Run()
	sync.Run()
	//gorm.Run()

	time.Sleep(10 * time.Minute)
}
