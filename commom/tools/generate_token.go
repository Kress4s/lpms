package tools

import (
	"lpms/constant"
	"time"

	"github.com/iris-contrib/middleware/jwt"
)

func Token(userID int64, username string) (string, int64) {
	exp := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 根据需求，可以存一些必要的数据
		"user_id":   userID,
		"user_name": username,
		// 设定过期时间，便于测试，设置24小时过期
		"exp": exp,
	})
	// 使用设置的秘钥，签名生成jwt字符串
	tokenString, _ := token.SignedString([]byte(constant.Salt))
	return tokenString, exp
}
