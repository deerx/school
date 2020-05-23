package mypkg

import (
	"database/sql"
	"fmt"

	//postgres包
	_ "github.com/lib/pq"
)

var (
	//DB 定义一个初始化连接池
	DB *sql.DB
)

const (
	//定义常量数据库连接
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "school"
)

//ConnentDB 数据库连接
func ConnentDB() *sql.DB {
	//创建连接池
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//open方法返回连接池和错误
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//尝试ping通
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successful connented")
	return db

}

//初始化连接池
func init() {
	DB = ConnentDB()
}

func GetDB() *sql.DB {
	return ConnentDB()
}
