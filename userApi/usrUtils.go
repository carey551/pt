package userApi

import (
	"fmt"
	"project/utils"
)

var (
	Token_addr_yaml_local string = "C:\\Users\\Lenovo\\Desktop\\gowork\\src\\project\\yaml\\token.yaml"
	Token_addr_yaml_ar    string = "C:\\Users\\ar_wensi\\Desktop\\go\\project\\yaml\\token.yaml"
)

// 获取yaml中的token
type ConfigToken struct {
	Token string
}

func GetToken() (string, error) {
	// 获取token
	var config ConfigToken
	err := utils.ReadYAML(Token_addr_yaml_local, &config)
	if err != nil {
		return "22222", fmt.Errorf("读取失败%v", err)
	}
	// fmt.Printf("读取的内容%v", config.Token)
	n := 0
	for {
		if n > 3 || len(config.Token) > 0 {
			return config.Token, nil
		}
		if len(config.Token) == 0 {
			//读取的内容是空的，就发送登录请求
			err := Login()
			if err != nil {
				// fmt.Println(err)
				return "1111", fmt.Errorf("Log的错误", err)
			}
			return config.Token, nil
		}

		fmt.Printf("n的值==%v", n)
		n++
	}

}

// 为请求头添加token
func AddHeaderToken() (map[string]interface{}, error) {
	headPayload := make(map[string]interface{})
	token, err := GetToken()
	if err != nil {
		fmt.Printf("token获取失败%v", err)
		return nil, err
	}
	headPayload["Authorization"] = "Bearer " + token
	return headPayload, nil
}
