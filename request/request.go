package request

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"project/utils"
	"time"
)

// post请求
// args 一般用来传请求头，比如token
func PostRequest(payload map[string]interface{}, api string, args ...map[string]interface{}) ([]byte, error) {
	url := "https://sit-tenantadmin-3003.mggametransit.com" + api
	// 判断传进来的paylaod是否有签名，没有就添加上
	_, exists := payload["signature"]
	if !exists {
		payload["signature"] = ""
	}
	verfiyp := ""
	signature := utils.GetSignature(payload, &verfiyp)
	if signature == "" {
		fmt.Println("utils的签名是空的", signature)
	}

	payload["signature"] = signature
	// fmt.Printf("请求的payload%v\n", payload)
	//将请求数据转换成json
	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf(" json 编码失败:%v", err)
	}
	// 发送post请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("请求失败：%v", err)
	}
	// 设置请求头
	if len(args) > 0 {
		// for key, value := range args {
		// 	strValue, ok := value.(string)
		// 	if ok {
		// 		req.Header.Set(string(key), strValue)
		// 	} else {
		// 		return nil, fmt.Errorf("header %s 的值不是字符串类型，实际类型: %T", key, value)
		// 	}
		// }
		// 直接访问 token
		if authorization, ok := args[0]["Authorization"].(string); ok {
			req.Header.Set("Authorization", authorization)
		} else {
			fmt.Println("错误: Authorization 不存在或不是字符串")
		}
	}
	req.Header.Set(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
	)
	req.Header.Set("Domainurl", "https://sit-tenantadmin-3003.mggametransit.com")
	req.Header.Set("Referer", "https://sit-tenantadmin-3003.mggametransit.com")
	req.Header.Set("Origin", "https://sit-tenantadmin-3003.mggametransit.com")
	req.Header.Set("Content-Type", "application/json")
	//本次请求的请求头
	// for key, values := range req.Header {
	// 	fmt.Printf("本次请求的请求头%s: %v\n", key, values)
	// }

	client := checkHttp2()
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("请求失败： %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		//获取相应的内容
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("读取响应失败：%v", err)
		}
		return respBody, nil

	} else if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		//
		return nil, nil
	} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		//需要身份验证,或者需要
		return nil, nil
	} else {
		err := errors.New("状态码不是200~~或者是服务器错误~~~")
		return nil, err
	}

}

// get请求
func GetRequest() {

}

// 验证请求的协议是不是http/2
func checkHttp2() *http.Client {
	client := &http.Client{
		// 检查 确保使用http/2
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
		},
	}
	return client
}

// 获取微妙级的时间戳并返回
func GetNowTime() int64 {
	now := time.Now()
	timestampMicro := now.Unix()
	fmt.Println("当前时间戳", timestampMicro)
	return timestampMicro
}

func RandmoNie() int64 {
	//生成9位的随机数
	max_number := big.NewInt(900000000000)
	n, err := rand.Int(rand.Reader, max_number)
	if err != nil {
		fmt.Printf("生成随机数失败：%v", err)
		return -1
	}
	random_number := n.Int64() + 100000000000
	fmt.Println("生成的随机数", random_number)
	return random_number
}
