package betApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"project/request"
	"project/utils"
	"time"
)

// BetRequest 定义请求体的结构体
type BetRequest2 struct {
	GameCode    string `json:"gameCode"`
	IssueNumber string `json:"issueNumber"`
	Amount      int    `json:"amount"`
	BetMultiple int    `json:"betMultiple"`
	BetContent  string `json:"betContent"`
	Language    string `json:"language"`
	Random      int64  `json:"random"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

func BetWingo2(gameCode string, amount int64, betContent, issueNumber, token string) {
	fmt.Println("请求了新的请求BetWingo2")
	url := "https://sit-lotteryh5.wmgametransit.com/api/Lottery/WinGoBet"

	// 创建请求体
	bet := BetRequest{
		GameCode:    gameCode,
		IssueNumber: issueNumber,
		Amount:      int(amount),
		BetMultiple: 1,
		BetContent:  betContent,
		Language:    "en",
		Random:      request.RandmoNie(),  // 确保 RandmoNie 实现正确
		Timestamp:   request.GetNowTime(), // 确保 GetNowTime 返回秒级时间戳
		Signature:   "",
	}

	// 计算签名
	verfiy := "" // 替换为实际密钥
	signalVal := utils.GetSignature2(bet, &verfiy)
	bet2 := BetRequest{
		GameCode:    gameCode,
		IssueNumber: issueNumber,
		Amount:      int(amount),
		BetMultiple: 1,
		BetContent:  betContent,
		Language:    "en",
		Random:      bet.Random,
		Signature:   signalVal,
		Timestamp:   bet.Timestamp,
	}

	// 序列化为 JSON
	body, err := json.Marshal(bet2)
	if err != nil {
		fmt.Println("序列化请求体失败:", err)
		return
	}

	// 调试：打印请求体和签名
	fmt.Println("请求体:", string(body))
	fmt.Println("签名:", signalVal)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json") // 改为 application/json
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-ch-ua", `"Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Origin", "https://h5.wmgametransit.com")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://h5.wmgametransit.com/")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")

	// 调试：打印完整请求
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("打印请求失败:", err)
		return
	}
	fmt.Println("完整请求:", string(dump))

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 30 * time.Second, // 增加超时时间
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应状态
	fmt.Println("响应状态:", resp.Status)

	// 打印响应体
	var respBody bytes.Buffer
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}
	fmt.Println("响应体:", respBody.String())
}
