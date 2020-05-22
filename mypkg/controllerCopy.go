package mypkg

import (
	"fmt"
	"strconv"
)

func GetRoomBegin(userName string) (bool, string, string, string) {
	var (
		// rooms      []Room
		returnRoom Room
		number     int
	)
	// 开启事物锁
	tx, err1 := DB.Begin()
	if err1 != nil {
		fmt.Println("初始化事物失败")
		fmt.Println(err1)
		return false, "", "", ""
	} else {
		fmt.Println("初始化事物成功")
	}
	err := tx.QueryRow(findRoomSQL).Scan(&returnRoom.ID, &returnRoom.Type)
	if err != nil {
		fmt.Println("没有房间", err)
		clearLogAndRoom()
		return false, "", "", ""
	}
	timestr, endtime := GetTime()
	if returnRoom.ID != 0 {
		_, err5 := tx.Exec(updateRoomSQL, 1, returnRoom.ID)
		if err5 != nil {
			fmt.Println("修改房间状态失败执行回滚")
			fmt.Println(err5)
			tx.Rollback()
			return false, "", "", ""
		} else {
			fmt.Println("修改房间状态为占用", returnRoom)
		}

		// err4 := tx.QueryRow(insertLogSQL, strconv.Itoa(userNumber)+"测试", returnRoom.ID, "1", timestr, endtime).Scan(&number)
		err4 := tx.QueryRow(insertOrderSQL, userName, returnRoom.ID, "1", timestr, endtime).Scan(&number)
		if err4 != nil {
			fmt.Println("插入log订单失败执行回滚", err4)
			tx.Rollback()
			return false, "", "", ""
		}
	}
	tx.Commit()
	fmt.Println(userName, "预约成功执行Commit，获取房间为", returnRoom)
	return true, strconv.Itoa(returnRoom.ID), strconv.Itoa(number), timestr
}
