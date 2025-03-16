package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// WorkbookCategoryClosure holds the schema definition for the WorkbookCategoryClosure entity.
type WorkbookCategoryClosure struct {
	ent.Schema
}

// Fields of the WorkbookCategoryClosure.
func (WorkbookCategoryClosure) Fields() []ent.Field {
	return []ent.Field{
		field.Int("child_id"),
		field.Int("parent_id"),
	}
}

// Edges of the WorkbookCategoryClosure.
func (WorkbookCategoryClosure) Edges() []ent.Edge {
	return nil
}
