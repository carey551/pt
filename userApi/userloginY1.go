package userApi

import (
	"encoding/json"
	"fmt"
	"project/request"
	"project/utils"
)

type userloginY1 struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	LoginType   string `json:"loginType"`
	DeviceId    string `json:"deviceId"`
	BrowserId   string `json:"browserId"`
	PackageName string `json:"packageName"`
	Language    string `json:"language"`
	Random      int64  `json:"random"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

// 账号，密码 登录
func UserloginY1(username, password string) {
	api := "/api/Home/Login"
	userloginInit := userloginY1{
		UserName:    username,
		Password:    password,
		LoginType:   "Mobile",
		DeviceId:    "",
		BrowserId:   utils.GenerateCryptoRandomString(32),
		PackageName: "",
		Language:    "en",
		Random:      request.RandmoNie(),
		Signature:   "",
		Timestamp:   request.GetNowTime(),
	}

	userloginMap := map[string]interface{}{
		"userName":    userloginInit.UserName,
		"password":    userloginInit.Password,
		"loginType":   userloginInit.LoginType,
		"deviceId":    userloginInit.DeviceId,
		"browserId":   userloginInit.BrowserId,
		"packageName": userloginInit.PackageName,
		"language":    userloginInit.Language,
		"random":      userloginInit.Random,
		"signature":   userloginInit.Signature,
		"timestamp":   userloginInit.Timestamp,
	}

	resp, err := request.PostRequestY1(userloginMap, api)
	if err != nil {
		fmt.Printf("输入账号和密码登录的post请求失败")
		return
	}

	strResbody := string(resp)
	var response Response
	error := json.Unmarshal([]byte(strResbody), &response)
	if error != nil {
		fmt.Println(error)
		return
	}
	// fmt.Printf("登录结果%v", response)
	token, err := utils.HandlerMap(strResbody, "token")
	if err != nil {
		return
	}
	fmt.Printf("token==>%v", token)
}
