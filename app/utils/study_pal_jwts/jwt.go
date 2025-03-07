package study_pal_jwts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(exp int64, jwtSecretKey string, userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = exp

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func CreateAccessToken(jwtSecretKey string, userId int) (string, error) {
	token, err := createToken(time.Now().Add(time.Hour*1).Unix(), jwtSecretKey, userId)
	return token, err
}

func CreateRefreshToken(jwtSecretKey string, userId int) (string, error) {
	token, err := createToken(time.Now().Add(time.Hour*24*30).Unix(), jwtSecretKey, userId)
	return token, err
}

func VerifyToken(jwtSecretKey string, token string) (bool, int, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return false, 0, err
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	return parseToken.Valid, int(claims["user_id"].(float64)), nil
}
