package vo

type LoginReq struct {
	// 用户名
	UserName string `json:"user_name"`
	// 密码
	Password string `json:"password"`
}
