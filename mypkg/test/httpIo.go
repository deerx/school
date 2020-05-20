package mypkg

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// TestIo 是应用程序的入口
func TestIo() {
	var b bytes.Buffer

	// 将字符串写入 Buffer
	b.Write([]byte("Hello"))

	// 使用 Fprintf 将字符串拼接到 Buffer
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容写到 Stdout
	io.Copy(os.Stdout, &b)
}

// type notify

func (user *UserInfo) notify() {
	fmt.Println(user.Name, user.Email)
}

//UserInfo 测试
type UserInfo struct {
	Name  string
	Email string
}

//TestInfo 测试
func TestInfo() {
	user := UserInfo{Name: "zhangsan", Email: "123456@qq.com"}
	user.notify()
}
