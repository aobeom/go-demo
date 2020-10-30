package demo

import (
	"go-demo/parallel"
	"go-demo/tools"
)

func run(url string) interface{} {
	data := tools.MiniJSONData{
		"province": "北京市",
		"city": "",
		"town": "",
	}
	x := tools.Minireq.Post(url, data)
	return x
}

// Demo1 并发示例1 发送海量 HTTP 请求
func Demo1() {
	defer tools.TimeSuite.RunTime()()
	url := "http://127.0.0.1:5000/student/list"
	taskList := make([]string, 0)
	for i := 0; i < 1000; i++ {
		taskList = append(taskList, url)
	}
	parallel.Run(run, taskList, 2)
}
