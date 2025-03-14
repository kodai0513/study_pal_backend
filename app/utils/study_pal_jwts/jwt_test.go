package study_pal_jwts

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var secretKey string = "secretKey"

func TestCreateAccessTokens_正常にアクセストークンを生成できるか(t *testing.T) {
	token := CreateAccessToken(secretKey, 1)
	if len(token) <= 0 {
		t.Fatal("token not empty")
	}
}

func TestCreateRefreshToken_正常にリフレッシュトークンを生成できるか(t *testing.T) {
	token := CreateAccessToken(secretKey, 1)
	if len(token) <= 0 {
		t.Fatal("token not empty")
	}
}

func TestVerifyToken_有効なトークンの時に正しく判定できるか(t *testing.T) {
	token := CreateAccessToken(secretKey, 1)
	if len(token) <= 0 {
		t.Fatal("token not empty")
	}

	userId, err := VerifyToken(secretKey, token)
	assert.NoError(t, err)
	assert.Equal(t, userId, 1)
}

func TestVerifyToken_有効期限切れのトークンの時に正しく判定できるか(t *testing.T) {
	token := createToken(time.Now().Add(time.Hour*-1).Unix(), secretKey, 1)
	userId, err := VerifyToken(secretKey, token)
	assert.Equal(t, 0, userId)
	assert.Error(t, err)
}

func TestVerifyToken_無効なトークンの時に正しく判定できるか(t *testing.T) {
	invalidToken := "invalidToken"
	userId, err := VerifyToken(secretKey, invalidToken)
	assert.Equal(t, 0, userId)
	assert.Error(t, err)
}
