package server

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sy-vendor/public"
	"net/http"
	"time"
)

func signTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	fmt.Println("\nSignTask is running...")

	//模拟延时
	time.Sleep(time.Second * 2)

	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_uid, found1 := req.Form["uid"]
	param_rid, found2 := req.Form["rid"]
	param_timestamp, found3 := req.Form["timestamp"]
	param_sign, found4 := req.Form["sign"]

	if !(found1 && found2 && found3 && found4) {
		fmt.Fprint(w, "请勿非法访问, 缺失参数")
		return
	}

	result := NewBaseJsonBean()
	uid := param_uid[0]
	rid := param_rid[0]
	sign := param_sign[0]
	timestamp := param_timestamp[0]

	s := "uid:" + uid + ",rid:" + rid + ",timestamp:" + timestamp + ",sign:" + sign
	fmt.Println(s)

	_, sign_str := public.MakeParams(uid, rid, timestamp)
	signRemote := public.MakeSign(sign_str)

	if sign == signRemote {
		result.Code = 0
		result.Data = "success"
		result.Message = "认证成功"
	} else {
		result.Code = 1
		result.Data = "fail"
		result.Message = "Sign认证失败"
	}

	userName := selectQuery()
	if userName != "err" {
		result.Data = userName
	}
	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func selectQuery() string {
	// Prepare statement for reading data
	db := public.GetDbConnetion()
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT user_name FROM user WHERE user_id = ?")
	if err != nil {
		fmt.Println("select error: ", err.Error())
		return "err"
	}
	defer stmtOut.Close()

	var userName string // we "scan" the result in here

	// Query the square-number of 1
	err = stmtOut.QueryRow(1).Scan(&userName) // WHERE number = 1
	if err != nil {
		fmt.Println("select error: ", err.Error())
		return "err"
	}
	fmt.Println("The user_name of 1 is: ", userName)
	return userName
}
