package demo

import (
	"fmt"
	"sort"
)

// Student 学生成绩表
type Student struct {
	Name  string
	Score int64
	Age   int64
}

// OrderBy 排序入口
type OrderBy func(p, q *Student) bool

// StudentWrapper 学生成绩表排序
type StudentWrapper struct {
	StudentSlice []Student
	OrderBy      func(p, q *Student) bool
}

// StudentOrder 学生成绩排序条件
func StudentOrder(student []Student, orderBy OrderBy) {
	sort.Sort(StudentWrapper{student, orderBy})
}

// StudentGroup 学生成绩表分类
type StudentGroup struct {
	Age  int64
	Info []Student
}

func (s StudentWrapper) Len() int {
	return len(s.StudentSlice)
}

func (s StudentWrapper) Swap(i, j int) {
	s.StudentSlice[i], s.StudentSlice[j] = s.StudentSlice[j], s.StudentSlice[i]
}

func (s StudentWrapper) Less(i, j int) bool {
	return s.OrderBy(&s.StudentSlice[i], &s.StudentSlice[j])
}

// GroupByAge 按年龄分组
func GroupByAge(sw StudentWrapper) (data []StudentGroup) {
	data = make([]StudentGroup, 0)
	studentData := sw.StudentSlice
	studentOrder := sw.OrderBy
	StudentOrder(studentData, studentOrder)
	var i int = 0
	var j int = 0

	dataLen := sw.Len()
	dataGroupList := make([][]Student, 0)

	for i < dataLen {
		j = i + 1
		if j < dataLen && studentData[i].Age == studentData[j].Age {
			j++
		}
		dataGroupList = append(dataGroupList, studentData[i:j])
		i = j
	}
	
	for _, dataGroup := range dataGroupList {
		var mainData StudentGroup
		mainData.Info = make([]Student, 0)
		for _, data := range dataGroup {
			mainData.Age = data.Age
			mainData.Info = append(mainData.Info, data)
		}
		data = append(data, mainData)
	}
	return
}

// Demo2 简单排序
func Demo2() {
	var sw StudentWrapper

	student := []Student{
		{"张三", 75, 15},
		{"李四", 85, 16},
		{"王五", 100, 14},
		{"赵六", 77, 15},
	}

	sw.StudentSlice = student
	sw.OrderBy = func(p, q *Student) bool { return p.Age < q.Age }
	// 相同年龄一组
	data := GroupByAge(sw)
	fmt.Println(data)
}
