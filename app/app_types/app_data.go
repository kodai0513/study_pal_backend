package app_types

import (
	"study-pal-backend/ent"
)

type AppData struct {
	client       *ent.Client
	jwtSecretKey string
}

func NewAppData(client *ent.Client, jwtSecretKey string) *AppData {
	return &AppData{
		client: client,
	}
}

func (a *AppData) Client() *ent.Client {
	return a.client
}

func (a *AppData) JwtSecretKey() string {
	return a.jwtSecretKey
}
