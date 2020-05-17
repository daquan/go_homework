package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	id int
	name string
	status int
	startTime *time.Time
	endTime *time.Time
	user string
}

func time2str(time *time.Time) string {
	//当time传入的是一个空指针时，则返回一个空字符串
	if time == nil {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}

func main()  {
	now := time.Now()
	end := now.Add(time.Hour*24)
	task := &Task{
		id: 1,
		name: "整理课程笔记",
		status: 0,
		startTime: &now,
		endTime: &end,
		user: "daquan",
	}

	task2 := &Task{
		id: 2,
		name: "吃饭",
		status: 0,
		startTime: &now,
		endTime: &end,
		user: "daquan",
	}
	//定义一个一个存储Task指针类型的切片,并将task和task2放入切片中
	tasks := []*Task{task,task2}

	//创建一个文件类型指针，err忽略
	file, _ := os.Create("task.txt")
	defer file.Close()

	//遍历tasks获取task
	for _,task := range tasks {
		//可以使用Sprintf()来拼接字符串并返回字符串，并写入到文件中，注意每个task要换行
		file.WriteString(
			fmt.Sprintf("%d,%s,%d,%s,%s,%s\n",
				task.id,task.name,task.status,
				time2str(task.startTime),
				time2str(task.endTime),
				task.user),
			)
	}
}
