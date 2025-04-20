package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/role"
	"study-pal-backend/ent/user"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookmember"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookMember holds the schema definition for the WorkbookMember entity.
type WorkbookMember struct {
	ent.Schema
}

func (WorkbookMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookMember.
func (WorkbookMember) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("role_id", uuid.UUID{}).Unique(),
		field.UUID("user_id", uuid.UUID{}).Unique(),
		field.UUID("workbook_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the WorkbookMember.
func (WorkbookMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(role.Label, Role.Type).Ref(workbookmember.Table).Unique().Required().Field("role_id"),
		edge.From(user.Label, User.Type).Ref(workbookmember.Table).Unique().Required().Field("user_id"),
		edge.From(workbook.Label, Workbook.Type).Ref(workbookmember.Table).Unique().Required().Field("workbook_id"),
	}
}
