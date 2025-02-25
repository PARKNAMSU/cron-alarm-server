package jwt_tool

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSignMethod = jwt.SigningMethodHS256
)

func GenerateToken[T any](data T, secretKey string, expired time.Duration) string {
	// token 에 저장할 데이터
	claim := jwt.MapClaims{
		"userData": data,
		"exp":      time.Now().Add(expired).Unix(),
	}

	// 토큰생성
	t := jwt.NewWithClaims(jwtSignMethod, claim)

	token, err := t.SignedString([]byte(secretKey))

	if err != nil {
		log.Println(err.Error())
	}

	return token
}

func GetData[T any](token string, secretKey string) (T, error) {
	var data T

	// 토큰 검증 진행
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		if !t.Valid {
			return nil, fmt.Errorf("Invalid token")
		}
		return secretKey, nil
	})

	if err != nil {
		return data, err
	}

	// 토큰 데이터 get
	if claims, ok :=
		parsedToken.Claims.(jwt.MapClaims); ok {
		expiredAt, _ := claims["exp"].(int64)
		if expiredAt < time.Now().Unix() {
			return data, errors.New("token expired")
		}
		data, _ = claims["userData"].(T)
	} else {
		return data, errors.New("Invalid token")
	}
	return data, nil
}
