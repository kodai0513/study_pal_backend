package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"
	"study-pal-backend/ent/workbookcategory"
	"study-pal-backend/ent/workbookmember"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Workbook holds the schema definition for the Workbook entity.
type Workbook struct {
	ent.Schema
}

func (Workbook) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the Workbook.
func (Workbook) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("created_id", uuid.UUID{}),
		field.String("description").MaxLen(400).Nillable(),
		field.Bool("is_public").Default(false),
		field.String("title").MaxLen(255).NotEmpty(),
	}
}

// Edges of the Workbook.
func (Workbook) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(problem.Table, Problem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(workbookcategory.Table, WorkbookCategory.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(workbookmember.Table, WorkbookMember.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
