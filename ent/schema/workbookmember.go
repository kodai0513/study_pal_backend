package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/role"
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
		field.UUID(role.Label+"_id", uuid.UUID{}).Unique(),
		field.UUID("member_id", uuid.UUID{}).Unique(),
		field.UUID(workbook.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the WorkbookMember.
func (WorkbookMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(role.Label, Role.Type).Ref(workbookmember.Table).Unique().Required().Field(role.Label + "_id"),
		edge.From("member", User.Type).Ref(workbookmember.Table).Unique().Required().Field("member_id"),
		edge.From(workbook.Label, Workbook.Type).Ref(workbookmember.Table).Unique().Required().Field(workbook.Label + "_id"),
	}
}
