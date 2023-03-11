package services

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	signKey = []byte("dangao")
)

type CustomerClaims struct {
	Data Data `json:"data"`
	jwt.RegisteredClaims
}
type Data struct {
	Id int64 `json:"id"`
}

// CreateToken 创建token
func CreateToken(data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"nbf":  time.Date(2022, 12, 27, 12, 12, 12, 0, time.UTC).Unix(),
	})
	return token.SignedString(signKey)
}

func ParseToken(strToken string) map[string]any {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil || !token.Valid {
		return make(map[string]any)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return make(map[string]any)
	}
	return claims
}
