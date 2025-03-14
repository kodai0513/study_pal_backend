package study_pal_jwts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(exp int64, jwtSecretKey string, userId int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = exp

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func CreateAccessToken(jwtSecretKey string, userId int) string {
	return createToken(time.Now().Add(time.Hour*1).Unix(), jwtSecretKey, userId)
}

func CreateRefreshToken(jwtSecretKey string, userId int) string {
	return createToken(time.Now().Add(time.Hour*24*30).Unix(), jwtSecretKey, userId)
}

func VerifyToken(jwtSecretKey string, token string) (int, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(fmt.Errorf("unexpected signing method: %v", t.Header["alg"]))
		}

		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	return int(claims["user_id"].(float64)), nil
}
