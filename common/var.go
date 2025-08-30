package common

// url配置
type CofingURL struct {
	ADMIN_URL string // 后台地址
	H5_URL    string // 前台地址
	BET_URL   string // 投注地址
}

func (config *CofingURL) ConfigFile() *CofingURL {
	config.ADMIN_URL = "https://sit-tenantadmin-3003.mggametransit.com"
	config.H5_URL = "https://sit-webapi.mggametransit.com"
	config.BET_URL = "sit-lotteryh5.wmgametransit.com"
	return config
}

// 后台的请求头设置
type AdminHeaderConfig struct {
	Domainurl string
	Referer   string
	Origin    string
}

func (header *AdminHeaderConfig) AdminHeaderConfigFunc() map[string]interface{} {
	return map[string]interface{}{
		header.Domainurl: "https://sit-tenantadmin-3003.mggametransit.com",
		header.Referer:   "https://sit-tenantadmin-3003.mggametransit.com",
		header.Origin:    "https://sit-tenantadmin-3003.mggametransit.com",
	}
}

// 前台的请求头设置
type DeskHeaderConfig struct {
	tenantId string
	AdminHeaderConfig
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
}

func (bet *BetHeaderConfig) BetHeaderConfigFunc(token string) map[string]interface{} {
	return map[string]interface{}{
		bet.Origin:        " https://h5.wmgametransit.com",
		bet.Referer:       "https://h5.wmgametransit.com",
		bet.Authorization: "Bearer " + token,
	}
}
