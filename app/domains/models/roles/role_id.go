package roles

import "study-pal-backend/app/domains/models/shared"

type RoleId struct {
	value int
}

func NewRoleId(value int) (RoleId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return RoleId{value: 0}, err
	}

	return RoleId{value: id.Value()}, nil
}

func (r *RoleId) Value() int {
	return r.value
}
