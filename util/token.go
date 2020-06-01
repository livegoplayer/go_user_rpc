package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	UserSession
	jwt.StandardClaims
}

type UserSession struct {
	Uid      int
	UserName string
}

func CreateToken(userSession *UserSession) (tokenStr string, err error) {

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		*userSession,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //过期时间是两个小时
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(mySigningKey)
	if err != nil {
		return
	}

	return
}

func ParseToken(tokenStr string) (claims MyCustomClaims, err error) {
	mySigningKey := []byte("AllYourBase")

	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return
	}

	if claimsP, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//检查jwt token是否过期
		if claimsP.IssuedAt+claimsP.ExpiresAt > time.Now().Unix() {
			err = errors.New("jwt过期")
			return
		}

		claims = *claimsP
		return
	} else {
		err = errors.New("jwt解析错误")
		return
	}
}
