package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/selectionproblem"
	"study-pal-backend/ent/selectionproblemanswer"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookcategory"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SelectionProblem holds the schema definition for the SelectionProblem entity.
type SelectionProblem struct {
	ent.Schema
}

func (SelectionProblem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SelectionProblem.
func (SelectionProblem) Fields() []ent.Field {
	return []ent.Field{
		field.String("statement").MaxLen(255).NotEmpty(),
		field.UUID(workbook.Label+"_id", uuid.UUID{}).Unique(),
		field.UUID(workbookcategory.Label+"_id", uuid.UUID{}).Nillable().Optional().Unique(),
	}
}

// Edges of the SelectionProblem.
func (SelectionProblem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(selectionproblemanswer.Table, SelectionProblemAnswer.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From(workbook.Label, Workbook.Type).Ref(selectionproblem.Table).Unique().Required().Field(workbook.Label + "_id"),
		edge.From(workbookcategory.Label, WorkbookCategory.Type).Ref(selectionproblem.Table).Unique().Field(workbookcategory.Label + "_id"),
	}
}
