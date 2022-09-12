package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"pd-go-server/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type CustomClaims struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateToken(account, password string) (token string, err error) {
	err = nil
	currentTime := time.Now()
	expireTime := currentTime.Add(72 * time.Hour)

	claims := CustomClaims{
		account,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "pd-blog",
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString(jwtSecret)

	return
}

func ParseToken(token string) (c *CustomClaims, err error) {
	tokenClaims, tokenError := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenError != nil {
		err = tokenError
		return
	}

	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		c = claims
		return
	}

	return
}
