package ids

import (
	"github.com/google/uuid"
)

type Id struct {
	value uuid.UUID
}

func CreateId() Id {
	return Id{value: uuid.New()}
}

func NewId(value uuid.UUID) Id {
	return Id{value: value}
}

func (p *Id) Value() uuid.UUID {
	return p.value
}
