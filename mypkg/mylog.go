package mypkg

import (
	"fmt"
	"log"
	"os"
)

var (
	//Mylog 共享log
	Mylog *log.Logger
)

func init() {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
	}
	Mylog = log.New(file, "[log]", log.LstdFlags|log.Llongfile)
}
