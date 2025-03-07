package value_object_commons

import (
	"errors"
)

type Id struct {
	value int
}

func NewId(value int) (*Id, error) {

	if value <= 0 {
		return nil, errors.New("id is greater than or equal to 1")
	}

	return &Id{value: value}, nil
}

func (p *Id) Value() int {
	return p.value
}
