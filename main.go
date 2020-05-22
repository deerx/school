package main

import (
	"strconv"
	"time"

	"github.com/my/repo/mypkg"
)

func main() {

	// mypkg.RedisGetKey()
	// mypkg.RedisZsetTest()
	// mypkg.RedisTest()
	// mypkg.HTTPTest()
	// if kiana, err := mypkg.MyClient.Get("Kiana").Result(); err == nil {
	// 	fmt.Println(kiana)
	// }

	// mypkg.MyClient.Set("Kiana", "しんいさんのかのじょきやな", 0)

	// mypkg.TestType()

	// mypkg.TestIo()
	// mypkg.TestInfo()
	// mypkg.HTTPTest()
	// mypkg.ExitRooms()
	// mypkg.Getrooms()
	for i := 1; i <= 100; i++ {
		go mypkg.GetRoomBeginTest(strconv.Itoa(i))
	}
	time.Sleep(60 * time.Second)
}
