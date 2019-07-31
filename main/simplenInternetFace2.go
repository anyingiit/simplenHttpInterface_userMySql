package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func WebSerBase() {
	fmt.Println("This is webserver base!")
	http.HandleFunc("/login",loginTask)//接口名,对应的处理方法

	err:=http.ListenAndServe("127.0.0.1:8080",nil)//服务器要监听的主机地址和端口号,如果有返回错误信息则返回给err

	if err!=nil {//如果err存在则输出错误信息
		fmt.Println("ListenAndSer error:",err.Error())
	}
}
func loginTask(w http.ResponseWriter, req *http.Request) {//登录用方法
	fmt.Println("login Task is running...")
	time.Sleep(time.Second*2)//等待两秒
	req.ParseForm()//ParseForm方法用来解析表单提供的数据
	param_username,found:=req.Form["username"]
	param_password,found2:=req.Form["password"]

	if !(found && found2) {
		fmt.Fprintf(w,"请勿非法访问")
		return
	}
	result:=NewBaseJsonBean()
	userName:= param_username[0]
	password:=param_password[0]

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "testuser" && password == "123456" {
		result.Code=100
		result.Message="登录成功"
	}else {
		result.Code=101
		result.Message="用户名或密码不正确"
	}
	bytes,_:=json.Marshal(result)
	fmt.Fprintf(w,string(bytes))
}


func main() {
	WebSerBase()
}
