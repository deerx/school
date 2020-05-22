package mypkg

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

//string类型数据操作
//redis命令：set key val 设置一个字符串的key和value
func set(key, val string) {
	//有效期为0表示不设置有效期，非0表示经过该时间后键值对失效
	result, err := MyClient.Set(key, val, 0).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：get key
func get(key string) {
	val, err := MyClient.Get(key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：mset key1 val1 key2 val2 key3 val3 ...  设置多个字符串数据
func mset(key1, val1, key2, val2, key3, val3 string) {
	//以下三种方式都可以，习惯于对象操作的我优先选择第三种
	//result,err := MyClient.MSet(key1,val1,key2,val2,key3,val3).Result()
	//result,err := MyClient.MSet([]string{key1,val1,key2,val2,key3,val3}).Result()
	result, err := MyClient.MSet(map[string]interface{}{key1: val1, key2: val2, key3: val3}).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：mget key1 key2 key3 ...  获取多个字符串数据
func mget(key1, key2, key3 string) {
	vals, err := MyClient.MGet(key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：del key1 key2 key3 ... 删除指定数据
func del(key1, key2, key3 string) {
	result, err := MyClient.Del(key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：getrange key start end 返回字符串指定字节数据
func getrange(key string, start, end int64) {
	val, err := MyClient.GetRange(key, start, end).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：strlen key 返回字符串的长度
func strlen(key string) {
	len, err := MyClient.StrLen(key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：setex key time val 设置值并制定时间
func setex(key, val string, expire int) {
	//time.Duration其实也是int64，不过是int64的别名罢了，但这里如果expire使用int64也无法与time.Second运算，
	//因为int64和Duration虽然本质一样，但表面上属于不同类型，go语言中不同类型是无法做运算的
	result, err := MyClient.Set(key, val, time.Duration(expire)*time.Second).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：append key val 将数据追加到指定数据的尾部
func redisAppend(key, val string) {
	//将val插入key对应值的末尾，并返回新串长度
	len, err := MyClient.Append(key, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：exists key  判断该数据是否存在
func exists(key string) {
	//返回1表示存在，0表示不存在
	isExists, err := MyClient.Exists(key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isExists)
}

//hash类型数据操作
//redis命令：hset hashTable key val  设置指定key的值
func hset(hashTable, key, val string) {
	isSetSuccessful, err := MyClient.HSet(hashTable, key, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	//如果键存在这返回false，如果键不存在则返回true
	fmt.Println(isSetSuccessful)
}

//redis命令：hget hashTable key  获取指定key的值
func hget(hashTable, key string) {
	val, err := MyClient.HGet(hashTable, key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：hmset hashTable key1 val1 key2 val2 key3 val3 ...
//该函数本身有问题，只插入一个键值对的话相当于hset，可以成功
//如果插入一个以上的键值对则会报错：ERR wrong number of arguments for 'hset' command
//且go-redis官方本身也不推荐是用该函数
//func hmset(hashTable,key1,val1,key2,val2,key3,val3 string){
//	_,err := client.HMSet(hashTable,key1,val1,key2,val2,key3,val3).Result()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//redis命令：hmget hashTable key1 key2 key3 ...  获取hash中多个数据
func hmget(hashTable, key1, key2, key3 string) {
	vals, err := MyClient.HMGet(hashTable, key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：hdel hashTable key1 key2 key3 ...  删除hash中的数据
func hdel(hashTable, key1, key2, key3 string) {
	//返回1表示删除成功，返回0表示删除失败
	//只要至少有一个被删除则返回1（不存在的键不管），一个都没删除则返回0（不存在的则也算没删除）
	n, err := MyClient.Del(hashTable, key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：hgetall hashTable  返回所有的键值
func hgetall(hashTable string) {
	vals, err := MyClient.HGetAll(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：hexists hashTable key  判断key是否存在
func hexists(hashTable, key string) {
	isExists, err := MyClient.HExists(hashTable, key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isExists)
}

//redis命令：hlen hashTable  返回hash的长度
func hlen(hashTable string) {
	len, err := MyClient.HLen(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：hkeys hashTable  返回hash里的所有键
func hkeys(hashTable string) {
	keys, err := MyClient.HKeys(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range keys {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：hvals hashTable  返回hash里的所有值
func hvals(hashTable string) {
	vals, err := MyClient.HVals(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//list类型数据操作
//redis命令：lpush mylist val1 val2 val3 ...  将数据添加到左边
func lpush(mylist, val1, val2, val3 string) {
	//返回列表的总长度（即有多少个元素在列表中）
	n, err := MyClient.LPush(mylist, val1, val2, val3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//添加一条数据到hash--------------
func setStructToHash(myhash string, userInfo Student) {
	userInfoJSON, _ := json.Marshal(userInfo)
	isSetSuccessful, err := MyClient.HSet(myhash, userInfo.UserName, userInfoJSON).Result()
	if err != nil {
		log.Fatal(err)
	}
	//如果键存在这返回false，如果键不存在则返回true
	fmt.Println(isSetSuccessful)
}

//从hash获取一条数据--------------
func getStructToHash(myhash, username string) Student {
	var (
		student Student
	)
	val, err := MyClient.HGet(myhash, username).Result()
	if err != nil {
		Mylog.Println(err)
	}
	err = json.Unmarshal([]byte(val), &student)
	if err != nil {
		Mylog.Println(err)
	}

	return student
}

//存储结构体到redislist
func setStructToRedis(mylist string, persons []Person) {
	for _, V := range persons {
		data, _ := json.Marshal(V)
		n, err := MyClient.LPush(mylist, data).Result()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n)
	}

	// //json序列化
	// datas, _ := json.Marshal(testStruct)
	// //缓存数据
	// conn.Do("set", "struct3", datas)
	// //读取数据
	// rebytes, _ := redis.Bytes(conn.Do("get", "struct3"))
	// //json反序列化
	// object := &TestStruct{}
	// json.Unmarshal(rebytes, object)

}

//获取list内容并转换成结构体数组
func getStructToRedis(mylist string) {
	var (
		persons []Person
		person  Person
	)
	vals, err := MyClient.LRange(mylist, 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range vals {
		json.Unmarshal([]byte(v), &person)
		persons = append(persons, person)
	}

	fmt.Println(persons)
}

//redis命令：rpush mylist val1 val2 val3 ...  将数据添加到右边
func rpush(mylist, val1, val2, val3 string) {
	//返回列表的总长度（即有多少个元素在列表中）
	n, err := MyClient.RPush(mylist, val1, val2, val3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：lpop mylist 删除左边指定第一个值并返回
func lpop(mylist string) {
	//返回被删除的值
	val, err := MyClient.LPop(mylist).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：rpop mylist 删除右边指定的值并返回
func rpop(mylist string) {
	//返回被删除的值
	val, err := MyClient.RPop(mylist).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：lrem mylist count val  删除指定count个元素
func lrem(mylist, val string, count int64) {
	//返回成功删除的val的数量
	n, err := MyClient.LRem(mylist, count, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：ltrim mylist start end 修剪集合指定内容
func ltrim(mylist string, start, end int64) {
	//返回状态（OK）
	status, err := MyClient.LTrim(mylist, start, end).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

//redis命令：lset mylist index val  设置指定下表元素
func lset(mylist, val string, index int64) {
	status, err := MyClient.LSet(mylist, index, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

//redis命令：lindex mylist index 通过下标获取元素
func lindex(mylist string, index int64) {
	//通过索引查找字符串
	val, err := MyClient.LIndex(mylist, index).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：lrange mylist start end 遍历集合
func lrange(mylist string, start, end int64) {
	vals, err := MyClient.LRange(mylist, start, end).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：llen mylist  返回长度
func llen(mylist string) {
	len, err := MyClient.LLen(mylist).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//无序集合set类型数据操作
//redis命令：sadd myset val1 val2 val3 ...  向set中添加数据
func sadd(myset, val1, val2, val3 string) {
	n, err := MyClient.SAdd(myset, val1, val2, val3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：srem myset val  删除set中的值并返回索引
func srem(myset, val string) {
	//删除集合中的值并返回其索引
	index, err := MyClient.SRem(myset, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(index)
}

//redis命令：spop myset 随机删除一个值并返回
func spop(myset string) {
	//随机删除一个值并返回
	val, err := MyClient.SPop(myset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：smembers myset  随机遍历set
func smembers(myset string) {
	vals, err := MyClient.SMembers(myset).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：scard myset 返回set长度
func scard(myset string) {
	len, err := MyClient.SCard(myset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：sismember myset val   判断set中是否有该成员
func sismember(myset, val string) {
	//判断值是否为集合中的成员
	isMember, err := MyClient.SIsMember(myset, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isMember)
}

//redis命令：srandmember myset count  随机返回count个数值
func srandmembers(myset string, count int64) {
	vals, err := MyClient.SRandMemberN(myset, count).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//该函数是上一个函数在只随机取一个元素的情况
func srandmember(myset string) {
	val, err := MyClient.SRandMember(myset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：smove myset myset2 val 将第一个集合中指定的元素转移到第二个集合中去
func smove(myset, myset2, val string) {
	isSuccessful, err := MyClient.SMove(myset, myset2, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isSuccessful)
}

//redis命令：sunion myset myset2 ... 返回多个集合的共同元素
func sunion(myset, myset2 string) {
	vals, err := MyClient.SUnion(myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sunionstore desset myset myset2 ...  返回多个集合中的共同元素存储到指定集合中去
func sunionstore(desset, myset, myset2 string) {
	//返回新集合的长度
	n, err := MyClient.SUnionStore(desset, myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：sinter myset myset2 ...  返回第一个集合中其他集合共同包括的元素
func sinter(myset, myset2 string) {
	vals, err := MyClient.SInter(myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sinterstore desset myset myset2 ...  返回第一个集合中其他集合共同包括的元素并存储到指定集合当中去
func sinterstore(desset, myset, myset2 string) {
	n, err := MyClient.SInterStore(desset, myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：sdiff myset myset2 ...  返回第一个集合中其他集合不包括的元素
func sdiff(myset, myset2 string) {
	vals, err := MyClient.SDiff(myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sdiffstore desset myset myset2 ...  返回第一个集合中其他集合不包括的元素并添加到指定集合中去
func sdiffstore(desset, myset, myset2 string) {
	n, err := MyClient.SDiffStore(desset, myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//有序集合zset类型数据操作
//redis命令：zadd myzset score1 val1 score2 val2 score3 val3 ...  添加数据并指定下标  你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
func zadd(myzset, val1, val2, val3 string, score1, score2, score3 float64) {
	member1 := &redis.Z{score1, val1}
	member2 := &redis.Z{score2, val2}
	member3 := &redis.Z{score3, val3}

	n, err := MyClient.ZAdd(myzset, member1, member2, member3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//添加任意数量的数据到有序集合
func zaddN(myzset string, vals []string, scores []float64) {
	for i, val := range vals {
		fmt.Println(i, val)
		member := &redis.Z{scores[i], val}
		MyClient.ZAdd(myzset, member)

	}
}

//redis命令：zrem myzset val1 val2 ...  删除zset内的指定数据
func zrem(myzset, val1, val2 string) {
	n, err := MyClient.ZRem(myzset, val1, val2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：srange myzset start end [withscores] 遍历数据从第几个到底几个 负数下标表示倒数 加上withscores 会额外返回数据对应的分数
func zrange(myzset string, start, end, flag int64) {

	if flag == 0 {
		//不加withscores
		vals, err := MyClient.ZRange(myzset, start, end).Result()

		if err != nil {
			log.Fatal(err)
		}

		for k, v := range vals {
			fmt.Printf("k = %v v = %s\n", k, v)
		}
	} else if flag == 1 {
		//加withscores
		svals, err := MyClient.ZRangeWithScores(myzset, start, end).Result()

		if err != nil {
			log.Fatal(err)
		}

		for k, v := range svals {
			fmt.Printf("k = %v v = %s s = %.2f\n", k, v.Member, v.Score)
		}
	}
}

//redis命令：srevrange myzset start end [withscores]  倒过来遍历
func zrevrange(myzset string, start, end, flag int64) {
	if flag == 0 {
		//不加withscores
		vals, err := MyClient.ZRevRange(myzset, start, end).Result()

		if err != nil {
			log.Fatal(err)
		}

		for k, v := range vals {
			fmt.Printf("k = %v v = %s\n", k, v)
		}
	} else if flag == 1 {
		//加withscores
		svals, err := MyClient.ZRevRangeWithScores(myzset, start, end).Result()

		if err != nil {
			log.Fatal(err)
		}

		for k, v := range svals {
			fmt.Printf("k = %v v = %s s = %.2f\n", k, v.Member, v.Score)
		}
	}
}

//redis命令：zrangebyscore myzset start end [withscores] 查询数据 大于start但是小于end分数的数据 查询所有的默认使用 -inf +inf 显示所有数据
func zrangebyscore(myzset, start, end string, flag int) {
	if flag == 0 {
		//不加withscores
		vals, err := MyClient.ZRangeByScore(myzset, &redis.ZRangeBy{Min: start, Max: end, Count: 0}).Result()

		if err != nil {
			log.Fatal(err)
		}

		for k, v := range vals {
			fmt.Printf("k = %v v = %s\n", k, v)
		}
	} else if flag == 1 {
		//加withscores
		svals, err := MyClient.ZRangeByScoreWithScores(myzset, &redis.ZRangeBy{Min: start, Max: end, Count: 0}).Result()

		if err != nil {
			log.Fatal(err)
		}

		for k, v := range svals {
			fmt.Printf("k = %v v = %s s = %.2f\n", k, v.Member, v.Score)
		}
	}
}

//redis命令：zcard myzset  返回zset数量
func zcard(myzset string) {
	len, err := MyClient.ZCard(myzset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：zcount myzset minscore maxscore   返回分数在min和max之间的数据个数
func zcount(myzset, minscore, maxscore string) {
	n, err := MyClient.ZCount(myzset, minscore, maxscore).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：zrank myzset val  返回指定字段的分数排名（从大到小）
func zrank(myzset, val string) {
	index, err := MyClient.ZRank(myzset, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(index)
}

//redis命令：zscore myzset val  返回指定字段的分数
func zscore(myzset, val string) {
	score, err := MyClient.ZScore(myzset, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)
}
