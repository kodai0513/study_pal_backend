package schema

import (
	"study-pal-backend/ent/descriptionproblem"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/selectionproblem"
	"study-pal-backend/ent/trueorfalseproblem"
	"study-pal-backend/ent/workbookcategory"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookCategoryDetail holds the schema definition for the WorkbookCategoryDetail entity.
type WorkbookCategoryDetail struct {
	ent.Schema
}

func (WorkbookCategoryDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookCategoryDetail.
func (WorkbookCategoryDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
		field.UUID(workbookcategory.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the WorkbookCategoryDetail.
func (WorkbookCategoryDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(descriptionproblem.Table, DescriptionProblem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(selectionproblem.Table, SelectionProblem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(trueorfalseproblem.Table, TrueOrFalseProblem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
