package main

import (
	"fmt"
	"project/betApi"
	_ "project/paymoneyapi"
	_ "project/userApi"
)

func run() {

	// userAmount := "918281997445"  // 需要添加的用户账号
	// userApi.Login() // 商户后台登录
	// userApi.AddUserRequest(userAmount) // 添加用户
	// userid := userApi.GetUserApi(userAmount) // 获取用户id
	// if userid == -1 {
	// 	return
	// }
	// paymoneyapi.ManualRecharge(userid, 667, 0) // 用户充值
	// userApi.UserloginY1(userAmount, "qwer1234")  // 前台登录
	isBet, result := betApi.IsBet()
	if isBet && result != "-1" {
		// 可以投注
		fmt.Println("可以投注")
	} else {
		return
	}
}

func main() {
	run()
}
