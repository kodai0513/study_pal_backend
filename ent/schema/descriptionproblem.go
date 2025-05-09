package schema

import (
	"study-pal-backend/ent/descriptionproblem"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookcategory"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DescriptionProblem holds the schema definition for the DescriptionProblem entity.
type DescriptionProblem struct {
	ent.Schema
}

func (DescriptionProblem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the DescriptionProblem.
func (DescriptionProblem) Fields() []ent.Field {
	return []ent.Field{
		field.String("correct_statement").MaxLen(255).NotEmpty(),
		field.String("statement").MaxLen(1000).NotEmpty(),
		field.UUID("workbook_id", uuid.UUID{}).Unique(),
		field.UUID("workbook_category_id", uuid.UUID{}).Nillable().Optional().Unique(),
	}
}

// Edges of the DescriptionProblem.
func (DescriptionProblem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(workbook.Label, Workbook.Type).Ref(descriptionproblem.Table).Unique().Required().Field("workbook_id"),
		edge.From(workbookcategory.Label, WorkbookCategory.Type).Ref(descriptionproblem.Table).Unique().Field("workbook_category_id"),
	}
}
