package vo

type LoginResponse struct {
	// token
	AccessToken string `json:"access_token"`
	// 认证类型
	TokenType string `json:"token_type"`
	// token 到期时间 默认两小时
	Expiry int64 `json:"expiry"`
	// 是否是管理员
	IsAdmin bool `json:"is_admin"`
}
