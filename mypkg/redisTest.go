package mypkg

import (
	"fmt"
)

func RedisGetKey() {

	val, err := MyClient.Get("key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// err = MyClient.Set("key2", "inininininininininin", 0).Err()
	// if err != nil {
	// 	Mylog.Println(err)
	// }
}

func RedisListTest() {
	var str = "emmmmm"
	redisAppend("key2", str)

}

func RedisZsetTest() {
	zrange("myzset1", 0, 10, 0)
}

func RedisTest() {
	// mset("name1", "老张", "name2", "小芳", "name3", "小丽")
	// mget("name1", "name2", "name3")
	// del("name1", "name2", "name3")
	// getrange("mykey", 0, 10)
	// strlen("mykey")
	// hset("mytable", "name", "莉")
	// hset("mytable", "sex", "女")
	// hset("mytable", "age", "15")
	// hget("mytable", "name")
	// hmget("mytable", "name", "age", "sex")
	// hlen("mytable")
	// hexists("mytable", "age")
	// hkeys("mytable")
	// hvals("mytable")
	// hgetall("mytable")
	// hdel("mytable", "name", "age", "sex")
	// lpush("mylist1", "1", "2", "3")
	// lrange("mylist1", 0, 100)
	// rpush("mylist1", "ppp", "ppp", "ppp")
	// lset("mylist1", "WWW", 0)
	// lindex("mylist1", 0)
	// sadd("myset1", "laoA", "laoB", "laoC")
	// scard("myset1")
	// smembers("myset1")
	// srem("myset1", "laohei")
	// srandmember("myset1")
	// srandmembers("myset1", 2)
	// sunion("myseyt1", "myset2")
	// sinter("myset1", "myset3")
	// sinterstore("myset4", "myset1", "myset2")
	// zadd("myzsetA", "xiaozhang", "xiaowang", "xiaohei", 1, 2, 3)

	// zrange("myzsetB", 0, -1, 1)
	// zrevrange("myzsetB", 0, -1, 1)
	// zrangebyscore("myzsetB", "-inf", "+inf", 1)
	// zcard("myzsetB")
	// zcount("myzsetB", "0", "1")
	// zrank("myzsetB", "你好")
	// zscore("myzsetB", "你好")

	// zaddN("myzsetC", []string{"李先生", "王先生", "张先生", "黑先生"}, []float64{77.5, 75.5, 88.5, 66.5})
	// zrangebyscore("myzsetC", "76", "90", 1)
	// zrevrange("myzsetC", 0, -1, 1)
	// zrank("myzsetC", "黑先生")

	// var (
	// 	persons []Person
	// )

	// persons = append(persons, Person{Name: "校长", Age: 50, Sex: "男"})
	// persons = append(persons, Person{Name: "张三", Age: 45, Sex: "男"})
	// persons = append(persons, Person{Name: "李四", Age: 45, Sex: "男"})
	// persons = append(persons, Person{Name: "王五", Age: 35, Sex: "男"})
	// persons = append(persons, Person{Name: "赵六", Age: 30, Sex: "男"})
	// persons = append(persons, Person{Name: "周七", Age: 25, Sex: "女"})

	// setStructToRedis("persons", persons)

	// getStructToRedis("persons")
	// Mylog.Println("giao")
	// fmt.Println("pos:%d:rs", 666)
	// strconv.FormatInt(time.Now().Add(-3 * time.Second).Unix())
	// fmt.Println(time.Second)
	// var student = Student{Name: "小张", ID: 1, StudentID: "XX140101", Password: "123456"}
	// setStructToHash("students", student)

	// fmt.Println(getStructToHash("students", "XX140101"))

}

// func DoGobEncodingStore() {
// 	var person1 Person
// 	person1.Name = "小红"
// 	person1.Age = 16
// 	person1.Sex = "女"
// 	//将数据进行gob序列化
// 	var buffer bytes.Buffer
// 	ecoder := gob.NewEncoder(&buffer)
// 	ecoder.Encode(person1)
// 	//reids缓存数据
// 	MyClient.Do("set", "person1", buffer.Bytes())
// 	//redis读取缓存
// 	rebytes, _ := redis.Bytes(MyClient.Do("get", "person1"))
// 	//进行gob序列化
// 	reader := bytes.NewReader(rebytes)
// 	dec := gob.NewDecoder(reader)
// 	object := &Person{}
// 	dec.Decode(object)
// }

// func DoHashStore() {
// 	//以hash类型保存
// 	MyClient.Do("hmset", redis.Args{"struct1"}.AddFlat(testStruct)...)
// 	//获取缓存
// 	value, _ := redis.Values(conn.Do("hgetall", "struct1"))
// 	//将values转成结构体
// 	object := &Student{}
// 	redis.ScanStruct(value, object)
// }
