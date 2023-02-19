package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("wannima")

type Claims struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签发 token
func GenerateToken(id uint, account string, authority int) (string, error) {
	issTime := time.Now()
	expTime := issTime.Add(time.Hour * 24)
	claims := Claims{
		ID:        id,
		Account:   account,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			IssuedAt:  issTime.Unix(),
			Issuer:    "E-COMMERCE",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// 解析 token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
