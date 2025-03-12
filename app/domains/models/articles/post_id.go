package articles

import (
	"study-pal-backend/app/domains/models/shared"
)

type PostId struct {
	value int
}

func NewPostId(value int) (PostId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return PostId{value: value}, err
	}
	return PostId{value: id.Value()}, nil
}

func (p *PostId) Value() int {
	return p.value
}
