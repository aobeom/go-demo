package demo

import (
	"fmt"
	"go-demo/parallel"
	"go-demo/tools"

	"time"
)

func run(count interface{}) interface{} {
	c := count.(int)
	c++
	time.Sleep(time.Duration(2) * time.Second)
	return c
}

// Demo1 并发示例1 发送海量 HTTP 请求
func Demo1() {
	defer tools.TimeSuite.RunTime()()
	taskList := make([]interface{}, 0)
	for i := 0; i < 10; i++ {
		taskList = append(taskList, i)
	}
	results := parallel.TaskBoard(run, taskList, 8)
	fmt.Println(len(results))
}
