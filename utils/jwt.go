package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"pentag.kr/BuildinAuth/configs"
)

type AuthTokenClaims struct {
	UserID string `json:"id"` // 유저 ID
	jwt.RegisteredClaims
}

func CreateToken(userID string) (string, error) {
	var err error
	//Creating Access Token
	now := time.Now()
	at := AuthTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: now.Add(time.Duration(WasConfig.JWT_EXPIRE) * time.Second)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, at)
	tokenString, err := token.SignedString([]byte(WasConfig.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ParseToken(tokenString string) (*AuthTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(WasConfig.JWT_SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
		return &AuthTokenClaims{}, err
	}
	claims, ok := token.Claims.(*AuthTokenClaims)
	if !ok {
		return &AuthTokenClaims{}, err
	}
	return claims, nil
}

var WasConfig = configs.Config.WAS
