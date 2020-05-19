package mypkg

import (
	"fmt"
)

//Student 学生实体类
type Student struct {
	ID        int
	StudentID string
	Password  string
	Name      string
}

//Room 浴室实体类
type Room struct {
	ID   int
	Type string
}

//View1 传输到第一个页面的数据
type View1 struct {
	Count string
	Text  string
}

//View2 传输到第二个页面的数据
type View2 struct {
	Number    string
	TimeStr   string
	RoomID    string
	StudentID string
}

//Logstr 记录信息
type Logstr struct {
	ID        int
	StudentID string
	RoomID    string
	Type      string
	TimeStr   string
	EndTime   string
}

//Person 实体类
type Person struct {
	Name string
	Age  int
	Sex  string
}

type MyType = int

func TestType() {
	var t MyType
	t = 6
	fmt.Println(t)

}
