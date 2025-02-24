package articles

import (
	"study-pal-backend/app/utils/application_errors"
)

type PostId struct {
	value int
}

func NewPostId(value int) (*PostId, application_errors.ApplicationError) {
	return &PostId{value: value}, nil
}

func (p *PostId) Value() int {
	return p.value
}
