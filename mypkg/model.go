package mypkg

import (
	"fmt"
	"time"
)

// Student 学生实体类
type Student struct {
	ID       int
	UserName string
	Password string
	Name     string
}

	// Room 浴室实体类
	type Room struct {
		ID   int
		Type int
	}

// Logstr 记录信息
type Order struct {
	ID       int
	UserName string
	RoomID   string
	Type     int
	TimeStr  string
	EndTime  time.Time
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

//Person 实体类
type Person struct {
	Name string
	Age  int
	Sex  string
}

//MyType 测试属性
type MyType = int

//TestType 测试测试属性
func TestType() {
	var t MyType
	t = 6
	fmt.Println(t)

}
