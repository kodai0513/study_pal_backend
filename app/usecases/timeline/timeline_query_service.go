package timelines

import (
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type TimelineDto struct {
	id           uuid.UUID
	description  string
	postId       uuid.UUID
	postName     string
	postNickName string
}

func NewTimelineDto(id uuid.UUID, description string, postId uuid.UUID, postName string, postNickName string) *TimelineDto {
	return &TimelineDto{
		id:           id,
		description:  description,
		postId:       postId,
		postName:     postName,
		postNickName: postNickName,
	}
}

func (t *TimelineDto) Id() uuid.UUID {
	return t.id
}

func (t *TimelineDto) Description() string {
	return t.description
}

func (t *TimelineDto) PostId() uuid.UUID {
	return t.postId
}

func (t *TimelineDto) PostName() string {
	return t.postName
}

func (t *TimelineDto) PostNickName() string {
	return t.postNickName
}

type TimelineQueryService interface {
	Fetch(page *app_types.Page) ([]*TimelineDto, *app_types.Page, usecase_error.UsecaseErrorGroup)
}
