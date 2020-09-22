package demo

import (
	"go-demo/parallel"
	"go-demo/tools"
	"net/http"
)

func run(t string) interface{} {
	data := make(map[string]string)
	data["province"] = "北京市"
	data["city"] = ""
	data["town"] = ""

	headers := make(http.Header)
	headers.Add("Content-Type", "application/json; charset=UTF-8")
	x := tools.Minireq.PostBody(t, headers, data)
	return x
}

// Demo1 并发示例1 发送海量 HTTP 请求
func Demo1() {
	defer tools.TimeCtl.RunTime()()
	url := "http://127.0.0.1:5000/student/list"
	taskList := make([]string, 0)
	for i := 0; i < 1000; i++ {
		taskList = append(taskList, url)
	}
	parallel.Run(run, taskList, 2)
}
