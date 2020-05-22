package mypkg

import (
	"time"
)

//GetTime 获取所需的时间
func GetTime() (string, string) {
	var (
		timestr string
	)
	t := time.Now().Format("01-02 15:04")
	t1 := time.Now() //.Add(time.Second * 1)
	// t1 := time.Now().Add(time.Minute * 30)
	timestr = t + " ~ " + t1.Format("01-02 15:04")
	return timestr, TimeToString(t1)
}

// TimeToString 将时间格式换成字符串
func TimeToString(timet time.Time) string {
	t := timet.Format("2006-01-02 15:04:05")
	return t
}

//StringToTime 将字符串格式换成时间
func StringToTime(times string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", times, time.Local)
	return t
}
