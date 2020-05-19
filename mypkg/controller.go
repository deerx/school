package mypkg

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	sessionMgr *SessionMgr = nil //session管理器
)

func init() {
	sessionMgr = NewSessionMgr("TestCookieName", 3600)
}

func HTTPTest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/index", addpHanderFunc(index))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("登陆成功这是登陆用户", getLoginUser(w, r))
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		Mylog.Println(err)
	}
	t.Execute(w, nil)
}

//处理登录
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("请求到了登陆页面")
		t, _ := template.ParseFiles("view/login.html")
		t.Execute(w, nil)

	} else if r.Method == "POST" {
		fmt.Println("请求到了验证")
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()

		//可以使用template.HTMLEscapeString()来避免用户进行js注入
		username := r.FormValue("username")
		password := r.FormValue("password")

		studentRedis := getStructToHash("students", username)

		// userRow := DB.QueryRow(username, password)
		// userRow.Scan(&userID)

		//TODO:判断用户名和密码
		if studentRedis.Password != "" && studentRedis.Password == password {
			//创建客户端对应cookie以及在服务器中进行记录
			var sessionID = sessionMgr.StartSession(w, r)

			//踢除重复登录的
			remRepeat(studentRedis.StudentID)

			//设置变量值
			sessionMgr.SetSessionVal(sessionID, "UserInfo", studentRedis)

			//TODO 设置其它数据

			//TODO 转向成功页面
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		} else {
			fmt.Println("账号或密码错误")
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}

//验证用户是否已经登陆没登陆重定向到登录页面
func testToken(w http.ResponseWriter, r *http.Request) bool {
	var sessionID = sessionMgr.CheckCookieValid(w, r)
	if sessionID == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}
	return true
}

//获取当前登陆对象
func getLoginUser(w http.ResponseWriter, r *http.Request) Student {
	var sessionID = sessionMgr.CheckCookieValid(w, r)
	fmt.Println("这是sessionId", sessionID)
	userInfo, flags := sessionMgr.GetSessionVal(sessionID, "UserInfo")
	if flags {
		// fmt.Println("返回了正确的对象")
		return userInfo.(Student)
	}
	return Student{}
}

//将验证用户是否登陆包装到路由中
func addpHanderFunc(a http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if testToken(w, r) {
			a(w, r)
		}
	}
}
func remRepeat(loginUserID string) {
	//踢除重复登录的
	var onlineSessionIDList = sessionMgr.GetSessionIDList()
	for _, onlineSessionID := range onlineSessionIDList {
		if userInfo, ok := sessionMgr.GetSessionVal(onlineSessionID, "UserInfo"); ok {
			if value, ok := userInfo.(Student); ok {
				if value.StudentID == loginUserID {
					sessionMgr.EndSessionBy(onlineSessionID)
				}
			}
		}
	}
}
