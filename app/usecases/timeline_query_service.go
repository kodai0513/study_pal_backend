package usecases

import (
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/application_errors"
)

type TimelineDto struct {
	id           int
	description  string
	postId       int
	postName     string
	postNickName string
}

func NewTimelineDto(id int, description string, postId int, postName string, postNickName string) *TimelineDto {
	return &TimelineDto{
		id:           id,
		description:  description,
		postId:       postId,
		postName:     postName,
		postNickName: postNickName,
	}
}

func (t *TimelineDto) Id() int {
	return t.id
}

func (t *TimelineDto) Description() string {
	return t.description
}

func (t *TimelineDto) PostId() int {
	return t.postId
}

func (t *TimelineDto) PostName() string {
	return t.postName
}

func (t *TimelineDto) PostNickName() string {
	return t.postNickName
}

type TimelineQueryService interface {
	Fetch(page *app_types.Page) ([]*TimelineDto, *app_types.Page, application_errors.ApplicationError)
}
