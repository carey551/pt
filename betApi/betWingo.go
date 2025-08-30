package betApi

import (
	"fmt"
	"project/request"
)

// 投注
type betApistruct struct {
	GameCode    string `json:"gameCode"`
	IssueNumber string `json:"issueNumber"`
	Amount      int64  `json:"amount"`
	BetMultiple int8   `json:"betMultiple"`
	BetContent  string `json:"betContent"`
	Language    string `json:"language"`
	Random      int64  `json:"random"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

/*
*
betType  彩票投注种类
amount 投注金额
betContent 投注盘口
issueNumber 期号
token token对象
*/
func BetWingo(gameCode string, amount int64, betContent, issueNumber string, token map[string]interface{}) {
	api := "/api/Lottery/WinGoBet"
	betapiData := betApistruct{
		GameCode:    gameCode,
		IssueNumber: issueNumber,
		Amount:      amount,
		BetMultiple: 1,
		BetContent:  betContent,
		Language:    "en",
		Random:      request.RandmoNie(),
		Signature:   "",
		Timestamp:   request.GetNowTime(),
	}

	betapiMap := map[string]interface{}{
		"gameCode":    betapiData.GameCode,
		"issueNumber": betapiData.IssueNumber,
		"amount":      betapiData.Amount,
		"betMultiple": betapiData.BetMultiple,
		"betContent":  betapiData.BetContent,
		"language":    betapiData.Language,
		"random":      betapiData.Random,
		"signature":   betapiData.Signature,
		"timestamp":   betapiData.Timestamp,
	}
	resp, err := request.PostRequestY1(betapiMap, api, token)
	if err != nil {
		fmt.Println("投注的post请求失败")
		return
	}
	fmt.Println(string(resp))
}
