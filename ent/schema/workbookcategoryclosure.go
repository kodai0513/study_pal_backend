package schema

import (
	"study-pal-backend/ent/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookCategoryClosure holds the schema definition for the WorkbookCategoryClosure entity.
type WorkbookCategoryClosure struct {
	ent.Schema
}

func (WorkbookCategoryClosure) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookCategoryClosure.
func (WorkbookCategoryClosure) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("workbook_id", uuid.UUID{}),
		field.UUID("child_id", uuid.UUID{}),
		field.Bool("is_root"),
		field.UUID("parent_id", uuid.UUID{}),
		field.Int("position"),
		field.Int("level"),
	}
}

// Edges of the WorkbookCategoryClosure.
func (WorkbookCategoryClosure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_category", WorkbookCategory.Type).Required().Unique().Field("child_id"),
		edge.To("parent_category", WorkbookCategory.Type).Required().Unique().Field("parent_id"),
	}
}
