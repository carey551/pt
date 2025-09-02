package common

// 商户后台的账号和密码
type AdminUserName struct {
	UserName string // 账号
	Pwd      string // 密码
}

func (admin *AdminUserName) AdminUserInit() *AdminUserName {
	admin.UserName = "carey3003"
	admin.Pwd = "qwer1234"
	return admin
}

var cofingURL CofingURL

// url配置
type CofingURL struct {
	ADMIN_URL string // 后台地址
	H5_URL    string // 前台地址
	BET_URL   string // 投注地址
	Iss_URL   string // 获取期号的地址
}

func (config *CofingURL) ConfigUrlInit() *CofingURL {
	config.ADMIN_URL = "https://sit-tenantadmin-3003.mggametransit.com"
	config.H5_URL = "https://sit-webapi.mggametransit.com"
	config.BET_URL = "https://sit-lotteryh5.wmgametransit.com"
	config.Iss_URL = "https://h5.wmgametransit.com"
	return config
}

// 后台的请求头设置不带token
type AdminHeaderConfig struct {
	Domainurl string
	Referer   string
	Origin    string
}

// 结构体的自我初始化
func NewAdminHeaderConfig() *AdminHeaderConfig {
	return &AdminHeaderConfig{
		Domainurl: "Domainurl",
		Referer:   "Referer",
		Origin:    "Origin",
	}
}

func (header *AdminHeaderConfig) AdminHeaderConfigFunc() map[string]interface{} {
	header = NewAdminHeaderConfig()

	url := cofingURL.ConfigUrlInit().ADMIN_URL
	return map[string]interface{}{
		header.Domainurl: url,
		header.Referer:   url,
		header.Origin:    url,
	}
}

// 后台请求头设置带token
type AdminHeaderAuthorizationConfig struct {
	Authorization string // token
	AdminHeaderConfig
}

func newAdminHeaderAuthorizationConfig() *AdminHeaderAuthorizationConfig {
	header := NewAdminHeaderConfig()
	return &AdminHeaderAuthorizationConfig{
		Authorization:     "Authorization",
		AdminHeaderConfig: *header,
	}
}

// token 把登录后的token
func (author *AdminHeaderAuthorizationConfig) AdminHeaderAuthorizationFunc(token string) map[string]interface{} {
	author = newAdminHeaderAuthorizationConfig()
	url := cofingURL.ADMIN_URL
	return map[string]interface{}{
		author.Authorization: "Bearer " + token,
		author.Domainurl:     url,
		author.Referer:       url,
		author.Origin:        url,
	}
}

// 前台的请求头设置
type DeskHeaderConfig struct {
	tenantId string
	AdminHeaderConfig
}

func NewDeskHeaderConfig() *DeskHeaderConfig {
	header := NewAdminHeaderConfig()
	return &DeskHeaderConfig{
		tenantId:          "tenantId",
		AdminHeaderConfig: *header,
	}
}

func (desk *DeskHeaderConfig) DeskHeaderConfigFunc() map[string]interface{} {
	return map[string]interface{}{
		desk.tenantId:  "3003",
		desk.Domainurl: "https://sit-plath5-y1.mggametransit.com",
		desk.Referer:   "https://sit-plath5-y1.mggametransit.com",
		desk.Origin:    "https://sit-plath5-y1.mggametransit.com",
	}
}

// 下注的请求头设置
type BetHeaderConfig struct {
	Referer       string
	Origin        string
	Authorization string
	Sec           string
	SecCh         string
	SecUa         string
	SecFetch      string
	SecFetchMode  string
	SecFetchDest  string
}

func NewBetHeaderConfig() *BetHeaderConfig {
	return &BetHeaderConfig{
		Referer:       "Referer",
		Origin:        "Origin",
		Authorization: "Authorization",
		Sec:           "sec-ch-ua-platform",
		SecCh:         "sec-ch-ua",
		SecUa:         "sec-ch-ua-mobile",
		SecFetch:      "Sec-Fetch-Site",
		SecFetchMode:  "Sec-Fetch-Mode",
		SecFetchDest:  "Sec-Fetch-Dest",
	}
}

func (bet *BetHeaderConfig) BetHeaderConfigFunc(token string) map[string]interface{} {
	bet = NewBetHeaderConfig()
	return map[string]interface{}{
		bet.Origin:        "https://h5.wmgametransit.com",
		bet.Referer:       "https://h5.wmgametransit.com",
		bet.Authorization: "Bearer " + token,
		bet.Sec:           "Windows",
		bet.SecCh:         `"Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"`,
		bet.SecUa:         "?0",
		bet.SecFetch:      "same-site",
		bet.SecFetchMode:  "cors",
		bet.SecFetchDest:  "empty",
	}
}

// 获取期号的请求头设置
type GetIssNunmberHeaderConfig struct {
	Referer string
}

func newGetIssNunmberHeaderConfig() *GetIssNunmberHeaderConfig {
	return &GetIssNunmberHeaderConfig{
		Referer: "Referer",
	}
}

/*
token
betType 投注的方式 wingo 30s  wingo1min  wingo 3min  wing 5min
**/
func (iss *GetIssNunmberHeaderConfig) GetIssNunmberHeaderFunc(token, betType string) map[string]interface{} {
	result := "https://h5.wmgametransit.com/WinGo/"
	iss = newGetIssNunmberHeaderConfig()
	if token == "" {
		//游客的方式
		result = result + betType
	} else {
		// token有值的情况
		r1 := "?Lang=en&Skin=Classic&SkinColor=Default&Token="
		r2 := "&RedirectUrl=https%3A%2F%2Fsit-plath5-y1.mggametransit.com%2Fgame%2Fcategory%3FcategoryCode%3DC202505280608510046&Beck=0"
		result = result + betType + r1 + token + r2
	}
	return map[string]interface{}{
		iss.Referer: result,
	}
}
