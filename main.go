package main

import (
	"fmt"

	"time"

	"github.com/jasonlvhit/gocron"
)

func task() {

	fmt.Println("task 被执行一次", time.Now().Format("2001-01-02 03:04:05"))

}

func main() {

	// 每五秒运行一次 task

	// 申请一个调度器

	s := gocron.NewScheduler()

	s.Every(1).Seconds().Do(task)

	<-s.Start()

}
