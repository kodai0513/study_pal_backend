package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/permission"
	"study-pal-backend/ent/role"
	"study-pal-backend/ent/workbookinvitationmember"
	"study-pal-backend/ent/workbookmember"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(workbookmember.Table, WorkbookMember.Type).Annotations(entsql.OnDelete(entsql.Restrict)),
		edge.To(workbookinvitationmember.Table, WorkbookInvitationMember.Type).Annotations(entsql.OnDelete(entsql.Restrict)),
		edge.From(permission.Table, Permission.Type).Ref(role.Table).Required(),
	}
}
