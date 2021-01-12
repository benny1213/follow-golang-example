package util

import (
	"fmt"
	"time"

	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/benny1213/etLab_BE/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims ： Claims
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken ：生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "test",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 此处jwtstring只接收[]byte类型 不然会报错 key is invalid or of invalid type
	token, err := tokenClaims.SignedString([]byte(setting.JwtString))

	return token, err
}

// ParseToken : 校验并解析token
func ParseToken(token string) (*Claims, error) {
	token = token[len("Bearer "):]
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.JwtString), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
		logging.Info(fmt.Sprintf("鉴权失败 tokenClaims： %v", tokenClaims))
	}
	return nil, err
}
