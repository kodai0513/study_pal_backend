package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/trueorfalseproblem"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookcategory"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TrueOrFalseProblem holds the schema definition for the TrueOrFalseProblem entity.
type TrueOrFalseProblem struct {
	ent.Schema
}

func (TrueOrFalseProblem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the TrueOrFalseProblem.
func (TrueOrFalseProblem) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_correct"),
		field.String("statement").MaxLen(255).NotEmpty(),
		field.UUID("workbook_id", uuid.UUID{}).Unique(),
		field.UUID("workbook_category_id", uuid.UUID{}).Nillable().Optional().Unique(),
	}
}

// Edges of the TrueOrFalseProblem.
func (TrueOrFalseProblem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(workbook.Label, Workbook.Type).Ref(trueorfalseproblem.Table).Unique().Required().Field("workbook_id"),
		edge.From(workbookcategory.Label, WorkbookCategory.Type).Ref(trueorfalseproblem.Table).Unique().Field("workbook_category_id"),
	}
}
