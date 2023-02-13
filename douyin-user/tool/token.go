package tool

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaim struct {
	ID       int
	UserName string
	NickName string
	jwt.RegisteredClaims
}

func GetToken(id uint, username, nickname string) (string, error) {
	u := &UserClaim{
		ID:       int(id),
		UserName: username,
		NickName: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, u).SignedString([]byte("douyin"))
	return token, err
}

func ParseToken(token string) (*UserClaim, error) {
	u := &UserClaim{}
	claims, err := jwt.ParseWithClaims(token, u, func(token *jwt.Token) (interface{}, error) {
		return []byte("douyin"), nil
	})
	if claim, ok := claims.Claims.(*UserClaim); ok && claims.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}
