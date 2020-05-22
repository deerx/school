package mypkg

import (
	// "mypkg/utils"
	// D:\sample\src\GoRedis\mypkg\utils
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

const (
	countTotalRoomSQL             = "select count(*) from room"
	countOccupyRoomSQL            = "select count(*) from room where type =0  "
	findRoomSQL                   = "select * from room where  type =0"
	updateRoomSQL                 = "update room set type = $1 where id = $2"
	insertOrderSQL                = "insert into orders (user_name,room_id,type,timestr,end_time) values($1,$2,$3,$4,$5) returning id"
	updateOrderSQL                = "update orders set type = 0 where user_name  = $1 and type = 1"
	findOrderSQL                  = "select room_id from orders where type = 1 and user_name = $1"
	clearOrderSQL                 = "update orders set type= 0 where end_time <now() and type = 1 returning room_id"
	findRoomByTypeAndStudentIDSQL = "select id,timestr from orders where user_name=$1 and type = 1"
	// 查找对象语句
	selectSQL = "select * from student where user_name = $1 and password = $2"
)

var (
	sessionMgr *SessionMgr = nil //session管理器

)

func init() {
	sessionMgr = NewSessionMgr("TestCookieName", 3600)
}

// HTTPTest 启动web项目
func HTTPTest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/index", addpHanderFunc(index))
	http.HandleFunc("/success", addpHanderFunc(success))
	http.HandleFunc("/exitRoom", addpHanderFunc(exitRoom))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		Mylog.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	var (
		viewEntity View1
		Entity     View2
	)
	loginUser := getLoginUser(w, r)
	DB.QueryRow(findRoomByTypeAndStudentIDSQL, loginUser.UserName).Scan(&Entity.Number, &Entity.TimeStr)
	if Entity.Number != "" && Entity.TimeStr != "" {
		tmpl, err := template.ParseFiles("view/success.html")
		if err != nil {
			Mylog.Println(err)
		}
		tmpl.Execute(w, Entity)
		return
	}
	Mylog.Println("登陆成功这是登陆用户", getLoginUser(w, r))
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		Mylog.Println(err)
	}
	viewEntity.Count = Count()
	Mylog.Println(viewEntity)
	t.Execute(w, viewEntity)
}

// 处理登录
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Mylog.Println("请求到了登陆页面")
		t, _ := template.ParseFiles("view/login.html")
		t.Execute(w, nil)

	} else if r.Method == "POST" {
		Mylog.Println("请求到了验证")
		// 请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()

		// 可以使用template.HTMLEscapeString()来避免用户进行js注入
		var student Student
		studentid := r.FormValue("studentid")
		password := r.FormValue("password")
		Mylog.Println("前端传过来的：学号" + studentid + "密码" + password)
		DB.QueryRow(selectSQL, studentid, password).Scan(&student.Name, &student.UserName, &student.Password)
		Mylog.Println("数据库查询到的user", student)

		// studentRedis := getStructToHash("students", username)
		// userRow := DB.QueryRow(username, password)
		// userRow.Scan(&userID)

		// TODO:判断用户名和密码
		if student.Password != "" && student.Password == password {
			// 创建客户端对应cookie以及在服务器中进行记录
			var sessionID = sessionMgr.StartSession(w, r)

			// 踢除重复登录的
			remRepeat(student.UserName)

			// 设置变量值
			sessionMgr.SetSessionVal(sessionID, "UserInfo", student)

			// TODO 设置其它数据

			// TODO 转向成功页面
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		} else {
			Mylog.Println("账号或密码错误")
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}

// 验证用户是否已经登陆没登陆重定向到登录页面
func testToken(w http.ResponseWriter, r *http.Request) bool {
	var sessionID = sessionMgr.CheckCookieValid(w, r)
	if sessionID == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}
	return true
}

// 获取当前登陆对象
func getLoginUser(w http.ResponseWriter, r *http.Request) Student {
	var sessionID = sessionMgr.CheckCookieValid(w, r)
	// Mylog.Println("这是sessionId", sessionID)
	userInfo, flags := sessionMgr.GetSessionVal(sessionID, "UserInfo")
	if flags {
		// Mylog.Println("返回了正确的对象")
		return userInfo.(Student)
	}
	return Student{}
}

// 将验证用户是否登陆包装到路由中
func addpHanderFunc(a http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if testToken(w, r) {
			a(w, r)
		}
	}
}
func remRepeat(loginUserID string) {
	// 踢除重复登录的
	var onlineSessionIDList = sessionMgr.GetSessionIDList()
	for _, onlineSessionID := range onlineSessionIDList {
		if userInfo, ok := sessionMgr.GetSessionVal(onlineSessionID, "UserInfo"); ok {
			if value, ok := userInfo.(Student); ok {
				if value.UserName == loginUserID {
					sessionMgr.EndSessionBy(onlineSessionID)
				}
			}
		}
	}
}

// Count 查询全部和剩余浴室
func Count() string {
	var (
		countOccupy int
		countTotal  int
	)
	DB.QueryRow(countOccupyRoomSQL).Scan(&countOccupy)
	DB.QueryRow(countTotalRoomSQL).Scan(&countTotal)
	return strconv.Itoa(countOccupy) + " / " + strconv.Itoa(countTotal)
}

func success(w http.ResponseWriter, r *http.Request) {
	var (
		Entity View2
	)
	Mylog.Println("请求到了预约界面")
	loginUser := getLoginUser(w, r)

	DB.QueryRow(findRoomByTypeAndStudentIDSQL, loginUser.UserName).Scan(&Entity.Number, &Entity.TimeStr)
	if Entity.Number != "" && Entity.TimeStr != "" {
		tmpl, err := template.ParseFiles("view/success.html")
		if err != nil {
			Mylog.Println(err)
		}
		tmpl.Execute(w, Entity)
	} else {
		// ExitRoomAndInsert(studentid)
		flag, roomid, number, timestr := GetRoomBegin(loginUser.UserName)
		if flag {
			Entity.Number = number
			Entity.RoomID = roomid
			Entity.TimeStr = timestr
			tmpl, err := template.ParseFiles("view/success.html")
			if err != nil {
				Mylog.Println(err)
			}
			tmpl.Execute(w, Entity)

		} else {
			var viewEntity View1
			tmpl, err := template.ParseFiles("view/index.html")
			if err != nil {
				Mylog.Println(err)
			}
			viewEntity.Count = Count()
			viewEntity.Text = "暂无空闲浴室请稍后预约"
			tmpl.Execute(w, viewEntity)
		}
	}

}

var mu sync.Mutex

// 修改房间状态以及插入使用记录
func updateRoomAndInsetLog(studentID string) (string, bool, string, string) {
	mu.Lock()
	var (
		flag   bool = false
		number int
	)
	boolean, room := FindRoom()
	timestr, endtime := GetTime()
	if boolean {
		DB.QueryRow(updateRoomSQL, "1", room.ID)

		// 插入一条使用记录
		DB.QueryRow(insertOrderSQL, studentID, room.ID, "1", timestr, endtime).Scan(&number)
		flag = true
	} else {
		mu.Unlock()
		return "", false, "", ""
	}
	mu.Unlock()
	return strconv.Itoa(room.ID), flag, strconv.Itoa(number), timestr
}

// FindRoom 查找空闲浴室的方法
func FindRoom() (bool, Room) {
	var (
		rooms []Room
		room  Room
	)
	room1 := Room{ID: 0, Type: -1}
	rows, err := DB.Query(findRoomSQL)
	if err != nil {
		fmt.Println("没找到房间", err)
		return false, room1
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&room.ID, &room.Type)
		if err != nil {
			Mylog.Println(err)
			break
		} else {
			// Mylog.Print("刚才添加到rooms集合的数据是")
			// Mylog.Println(room)
			rooms = append(rooms, room)
		}
	}
	// 如果没有剩余房间，先清除过期预约腾出房间
	if len(rooms) < 1 {
		clearLogAndRoom()

		return false, room1
	}
	return true, rooms[len(rooms)-1]
}

// 清空预约过期的房间
func clearLogAndRoom() {
	fmt.Println("执行清理")
	var (
		roomIDs []int
		roomID  int
	)
	rows, err := DB.Query(clearOrderSQL)
	if err != nil {
		fmt.Println("清理失败", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&roomID)
		if err != nil {
			Mylog.Println(err)
			break
		} else {
			roomIDs = append(roomIDs, roomID)
		}
	}
	fmt.Println("过期的浴室号码")
	fmt.Println(roomIDs)
	for _, id := range roomIDs {
		DB.QueryRow(updateRoomSQL, 0, id)
	}
}

// 取消房间预约
func exitRoom(w http.ResponseWriter, r *http.Request) {
	Mylog.Println("取消预约")
	var (
		Entity View1
	)
	loginUser := getLoginUser(w, r)
	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		Mylog.Println(err)
	}

	if loginUser.UserName != "" {
		ExitRoomAndInsert(loginUser.UserName)
		Entity.Text = "取消成功"
	} else {
		Entity.Text = "取消失败"
	}
	Entity.Count = Count()
	tmpl.Execute(w, Entity)
}

// ExitRoomAndInsert 取消被预约的数据记录
func ExitRoomAndInsert(studentID string) {
	var (
		roomids []int
		roomid  int
	)
	rows, err := DB.Query(findOrderSQL, studentID)
	if err != nil {
		Mylog.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&roomid)
		if err != nil {
			Mylog.Println(err)
			break
		} else {
			roomids = append(roomids, roomid)
		}

	}

	for _, id := range roomids {
		DB.QueryRow(updateRoomSQL, 0, id)
	}
	DB.QueryRow(updateOrderSQL, studentID)
}
