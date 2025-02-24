package timelines

import (
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/application_errors"
)

type TimeLineDto struct {
	Id           int
	Description  string
	PostId       int
	PostName     string
	PostNickName string
}

type TimeLineQueryService interface {
	Fetch(page *app_types.Page) ([]*TimeLineDto, *app_types.Page, application_errors.ApplicationError)
}
