package util

import (
	"androidHomeworkApi/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id int, username string) (gin.H, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour*24 * time.Duration( setting.JwtTokenEpx ) )

	claims := Claims{
		id,
		username,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return gin.H{
		"expireTime": expireTime,
		"token": token,
	}, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
