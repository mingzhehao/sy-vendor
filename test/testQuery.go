package main

import (
	"fmt"
	"github.com/go-simplejson"
	"github.com/sy-vendor/util"
	"time"
)

const (
	Uid     = 1
	UserAge = 100
	UserSex = 1
)

func main() {
	selectMethodOne()
	time.Sleep(time.Second * 2)
	selectMethodTwo()
	time.Sleep(time.Second * 2)
	updateMethod()
	time.Sleep(time.Second * 2)
	deleteMethod()
}

func selectMethodOne() {
	fmt.Println("\n测试query查询 GetUserById user_id = ", Uid)
	userInfo := util.GetUserByIdMethodOne(Uid)
	fmt.Println("Uid: ", userInfo.UserId)
	fmt.Println("UserName: ", userInfo.UserName)
	fmt.Println("UserAge: ", userInfo.UserAge)
}

func selectMethodTwo() {
	fmt.Println("\n测试query查询 GetUserById user_id = ", Uid)
	userInfoByte := util.GetUserByIdMethodTwo(Uid)

	/** json format --start-- **/
	json, err := simplejson.NewJson([]byte(userInfoByte))
	if err != nil {
		fmt.Println("json format error")
	}
	UserId, err := json.Get("user_id").Int()
	UserName, err := json.Get("user_name").String()
	UserAge, err := json.Get("user_age").Int()
	fmt.Println("json 解析后的信息 user_id:", UserId, " user_name:"+UserName+" user_age:", UserAge)
	/** json format --end-- **/

	userInfoString := string(userInfoByte)
	fmt.Println("解析为字符串后的数据 userInfo ", userInfoString)
}

func updateMethod() {
	fmt.Println("\n测试query 更新 UpdateUserInfoByUserId user_id = ", Uid)
	err := util.UpdateUserInfoByUserId(Uid, UserAge, UserSex)
	if err != nil {
		fmt.Println("更新存在错误 error: ", err)
	}
}

func deleteMethod() {
	fmt.Println("\n测试query 删除 DeleteUserInfoById user_id = ", Uid)
	err := util.DeleteUserInfoByUserId(Uid)
	if err != nil {
		fmt.Println("删除存在错误 error: ", err)
	}
}
