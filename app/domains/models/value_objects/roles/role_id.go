package roles

import (
	"study-pal-backend/app/master_datas/master_roles"

	"github.com/google/uuid"
)

type RoleId struct {
	value uuid.UUID
}

func AdminRoleId() RoleId {
	return RoleId{value: master_roles.Admin}
}

func EditorRoleId() RoleId {
	return RoleId{value: master_roles.Editor}
}

func ReaderRoleId() RoleId {
	return RoleId{value: master_roles.Reader}
}

func NewRoleId(value uuid.UUID) RoleId {
	return RoleId{value: value}
}

func (r *RoleId) Value() uuid.UUID {
	return r.value
}
