package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/role"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookmember"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WorkbookMember holds the schema definition for the WorkbookMember entity.
type WorkbookMember struct {
	ent.Schema
}

func (WorkbookMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookMember.
func (WorkbookMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int(role.Label + "_id"),
		field.Int("member_id"),
		field.Int(workbook.Label + "_id"),
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
