package main

import (
	"fmt"
	"os"
	"time"
)

type Task1 struct {
	id int
	name string
	status int
	startTime *time.Time
	endTime *time.Time
	user string
}

func time2str1(time *time.Time) string {
	//当time传入的是一个空指针时，则返回一个空字符串
	if time == nil {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}

func main()  {
	now := time.Now()
	end := now.Add(time.Hour*24)
	task := &Task1{
		id: 1,
		name: "整理课程笔记",
		status: 0,
		startTime: &now,
		endTime: &end,
		user: "daquan",
	}

	task2 := &Task1{
		id: 2,
		name: "吃饭",
		status: 0,
		startTime: &now,
		endTime: &end,
		user: "daquan",
	}
	//定义一个一个存储Task指针类型的切片,并将task和task2放入切片中
	tasks := []*Task1{task,task2}

	//创建一个文件类型指针，err忽略
	file, _ := os.Create("task1.txt")
	defer file.Close()

	//遍历tasks获取task
	for _,task := range tasks {
		//使用Fprintf()将字符串写入文件,此函数第一个参数是io.Writer
		fmt.Fprintf(file, "%d,%s,%d,%s,%s,%s\n",
			task.id, task.name, task.status,
			time2str1(task.startTime),
			time2str1(task.endTime),
			task.user)
	}
}

