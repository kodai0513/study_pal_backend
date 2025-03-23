package study_pal_jwts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func createToken(exp int64, jwtSecretKey string, userId uuid.UUID) string {
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

func CreateAccessToken(jwtSecretKey string, userId uuid.UUID) string {
	return createToken(time.Now().Add(time.Hour*1).Unix(), jwtSecretKey, userId)
}

func CreateRefreshToken(jwtSecretKey string, userId uuid.UUID) string {
	return createToken(time.Now().Add(time.Hour*24*30).Unix(), jwtSecretKey, userId)
}

func VerifyToken(jwtSecretKey string, token string) (uuid.UUID, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(fmt.Errorf("unexpected signing method: %v", t.Header["alg"]))
		}

		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	return uuid.MustParse(claims["user_id"].(string)), nil
}
