package main

import (
	"async/asyncq"
	"time"
)

func main() {

	var tasks []asyncq.Task

	for i:= 0; i<5; i++ {
		ht := asyncq.HelloTask{Name: "aaa"}
		ft := asyncq.ByeTask{Name: "bbb"}
		tasks = append(tasks, ht, ft)
	}

	asyncq.Dispatcher(10)
	for _, task := range tasks{
		asyncq.TaskQueue <- task
	}
	time.Sleep(1*time.Hour)
}
