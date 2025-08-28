package userApi

import (
	"fmt"
	"project/request"
	"project/utils"
)

type addUserInfoStruct struct {
	Account      string `json:"account"` // 添加的用户的账号91号码
	UserType     int8   `json:"userType"`
	PassWord     string `json:"password"`
	Remark       string `json:"remark"`
	RegisterType int8   `json:"registerType"`
}

// 添加用户
type AddUserStruct struct {
	AddUserList []addUserInfoStruct `json:"addUserList"`
	Random      int64               `json:"random"`
	Language    string              `json:"language"`
	Signature   string              `json:"signature"`
	Timestamp   int64               `json:"timestamp"`
}

// 发送添加用户的请求
func AddUserRequest() {
	//初始化这些添加用户的结构体
	api := "/api/Users/AddUsers"
	requesPayload := make(map[string]interface{})

	// 构建 addUserList 数组
	addUserList := []interface{}{
		map[string]interface{}{
			"account":      "917131997121",
			"userType":     0,
			"password":     "qwer1234",
			"remark":       "",
			"registerType": 1,
		},
	}

	// 填充 payload
	requesPayload["addUserList"] = addUserList
	requesPayload["random"] = request.RandmoNie()
	requesPayload["language"] = "zh"
	requesPayload["signature"] = ""
	requesPayload["timestamp"] = request.GetNowTime()

	// 给请求投添加token
	headPayload, err := AddHeaderToken()
	if err != nil {
		fmt.Println("添加头部token失败", err)
		return
	}

	responBody, err := request.PostRequest(requesPayload, api, headPayload)
	if err != nil {
		fmt.Println("添加用户失败", err)
		return
	}
	result := utils.Unmarshal(string(responBody))

	if len(result["data"].([]interface{})) == 0 {
		fmt.Printf("添加成功%v", result["data"])
	} else {
		fmt.Printf("添加失败%v", result["data"])
	}

}
