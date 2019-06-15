package asyncq

import (
	"fmt"
	"time"
)

type Task interface {
	Perform()
}

type HelloTask struct {
	Name string
}

func (h HelloTask) Perform() {
	time.Sleep(3 * time.Second)
	fmt.Println("hello ", h.Name)
}

type ByeTask struct {
	Name string
}

func (h ByeTask) Perform() {
	time.Sleep(5 * time.Second)
	fmt.Println("bye ", h.Name)
}
