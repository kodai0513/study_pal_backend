package timelines

import (
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type TimelineDto struct {
	Id           uuid.UUID
	Description  string
	UserId       uuid.UUID
	UserName     string
	UserNickName string
}

type TimelineQueryService interface {
	Fetch(page *app_types.Page) ([]*TimelineDto, *app_types.Page, usecase_error.UsecaseErrorGroup)
}
