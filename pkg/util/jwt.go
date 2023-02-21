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

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

func GenerateEmailToken(uid uint, email, password string, operation_type uint) (string, error) {
	issTime := time.Now()
	expTime := issTime.Add(time.Hour * 24)
	claims := EmailClaims{
		UserID:        uid,
		Email:         email,
		Password:      password,
		OperationType: operation_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			IssuedAt:  issTime.Unix(),
			Issuer:    "E-COMMERCE",
		},
	}
	emailtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	emailtokenString, err := emailtoken.SignedString(jwtSecret)
	return emailtokenString, err
}

func ParseEmailToken(emailtokenString string) (*EmailClaims, error) {
	token, err := jwt.ParseWithClaims(emailtokenString, &EmailClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*EmailClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
