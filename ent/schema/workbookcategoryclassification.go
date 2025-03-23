package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"
	"study-pal-backend/ent/workbookcategory"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookCategoryClassification holds the schema definition for the WorkbookCategoryClassification entity.
type WorkbookCategoryClassification struct {
	ent.Schema
}

func (WorkbookCategoryClassification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookCategoryClassification.
func (WorkbookCategoryClassification) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
		field.UUID(workbookcategory.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the WorkbookCategoryClassification.
func (WorkbookCategoryClassification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(problem.Table, Problem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
