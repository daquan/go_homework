package main

import (
	"bufio"
	"fmt"
	"os"
)

//type Task2 struct {
//	id int
//	name string
//	status int
//	startTime *time.Time
//	endTime *time.Time
//	user string
//}
//
//func time2str2(time *time.Time) string {
//	//当time传入的是一个空指针时，则返回一个空字符串
//	if time == nil {
//		return ""
//	}
//	return time.Format("2006-01-02 15:04:05")
//}

func main()  {
	//now := time.Now()
	//end := now.Add(time.Hour*24)
	//task := &Task2{
	//	id: 1,
	//	name: "整理课程笔记",
	//	status: 0,
	//	startTime: &now,
	//	endTime: &end,
	//	user: "daquan",
	//}
	//
	//task2 := &Task2{
	//	id: 2,
	//	name: "吃饭",
	//	status: 0,
	//	startTime: &now,
	//	endTime: &end,
	//	user: "daquan",
	//}
	//定义一个一个存储Task指针类型的切片,并将task和task2放入切片中
	//tasks := []*Task2{task,task2}

	//创建一个文件类型指针，err忽略
	file, _ := os.Open("task1.txt")
	defer file.Close()

	//创建一个Scanner结构体类型指针,传入一个输入流
	scanner := bufio.NewScanner(file)
	//循环，扫描的时候到文件末尾时，scanner.Scan()会返回一个false
	for scanner.Scan() {
		//每一行都是一个task任务信息
		taskInfo := scanner.Text()
		fmt.Printf("%T\n",taskInfo)
		fmt.Println(taskInfo)
	}

}
/*
输出结果：
string
1,整理课程笔记,0,2020-05-17 16:54:16,2020-05-18 16:54:16,daquan
string
2,吃饭,0,2020-05-17 16:54:16,2020-05-18 16:54:16,daquan
 */

//思考: task任务是自定义的结构体类型数据，从文件读取出来后，目前这种写法得到的确是字符串类型数据。那么怎样才能从文件读取出来后生成的是struct类型数据呢？

