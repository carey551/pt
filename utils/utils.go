package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// Md5Info 计算 MD5 哈希值
func Md5Info(data string, uppercase bool) string {
	hash := md5.New()
	hash.Write([]byte(data))
	result := hex.EncodeToString(hash.Sum(nil))
	if uppercase {
		return strings.ToUpper(result)
	}
	return strings.ToLower(result)
}

// GetSignature 生成签名
func GetSignature(body map[string]interface{}, verifyPwd *string) string {
	// 过滤字段
	filteredObj := make(map[string]interface{})
	keys := make([]string, 0, len(body))
	for key := range body {
		keys = append(keys, key)
	}
	sort.Strings(keys) // 按键排序

	for _, key := range keys {
		value := body[key]
		// 检查 value 不为 nil 且不为空字符串，且 key 不在排除列表中，且 value 不是数组
		if value != nil && value != "" && key != "signature" && key != "timestamp" && key != "track" {
			// 确保 value 不是切片（相当于 Python 的 list）
			if _, ok := value.([]interface{}); !ok {
				filteredObj[key] = value
			}
		}
	}

	// 转换为 JSON 字符串
	jsonData, err := json.Marshal(filteredObj)
	if err != nil {
		return "" // 错误处理，可根据需求调整
	}

	encoder := string(jsonData)
	if verifyPwd != nil {
		encoder += *verifyPwd
	}

	// 计算 MD5
	return Md5Info(encoder, true)
}

// 用来解析请求
func Unmarshal(strResbody string) (result map[string]interface{}) {
	//是一个字符串
	error := json.Unmarshal([]byte(strResbody), &result)
	if error != nil {
		log.Fatalf("解析响应失败~~:%v", error)
	}
	return

}

// 读取yaml
// ReadYAML 从指定路径读取 YAML 文件并解析到结构体
func ReadYAML(filePath string, result interface{}) error {
	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取 YAML 文件失败: %w", err)
	}

	// 解析 YAML 数据到结构体
	if err := yaml.Unmarshal(data, result); err != nil {
		return fmt.Errorf("解析 YAML 数据失败: %w", err)
	}

	return err
}

// 写入yaml 以覆盖的方式
// WriteYAML 以覆盖方式将结构体写入 YAML 文件
func WriteYAML(filePath string, data interface{}) error {
	// 将结构体编码为 YAML 数据
	yamlData, err := yaml.Marshal(data)
	// fmt.Printf("文件路径%s,写入的文件内容%v", filePath, yamlData)
	if err != nil {
		return fmt.Errorf("编码 YAML 数据失败: %w", err)
	}

	// 以覆盖方式写入文件（os.Create 会创建或覆盖文件）
	if err := os.WriteFile(filePath, yamlData, 0644); err != nil {
		return fmt.Errorf("写入 YAML 文件失败: %w", err)
	}

	return nil
}

// 处理两层map[string]interface{}
func HandlerMap(strResbody string, str string) (string, error) {
	var result map[string]interface{}
	result = Unmarshal(strResbody)
	innerMap, ok := result["data"].(map[string]interface{})
	if !ok {
		// fmt.Println("data 不存在")
		err := errors.New("data 不存在")
		return "第一层data不存在", err
	}

	// 访问内层 map 的 token
	token, ok := innerMap[str].(string)
	if !ok {
		// fmt.Println("token 不是字符串或不存在")
		err := errors.New("token 不是字符串或不存在")
		return "token 不是字符串或不存在", err
	}
	return token, nil
}

// 生成随机浏览器指纹
func GenerateCryptoRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	fmt.Println("本次指纹", string(bytes))
	return string(bytes)
}
