package app_types

import (
	"study-pal-backend/ent"
)

type AppData struct {
	client *ent.Client
}

func NewAppData(client *ent.Client) *AppData {
	return &AppData{
		client: client,
	}
}

func (a *AppData) Client() *ent.Client {
	return a.client
}
