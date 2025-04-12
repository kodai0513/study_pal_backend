package schema

import (
	"study-pal-backend/ent/descriptionproblem"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/selectionproblem"
	"study-pal-backend/ent/trueorfalseproblem"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookcategory"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookCategory holds the schema definition for the WorkbookCategory entity.
type WorkbookCategory struct {
	ent.Schema
}

func (WorkbookCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookCategory.
func (WorkbookCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
		field.UUID(workbook.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the WorkbookCategory.
func (WorkbookCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(descriptionproblem.Table, DescriptionProblem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(selectionproblem.Table, SelectionProblem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(trueorfalseproblem.Table, TrueOrFalseProblem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From(workbook.Label, Workbook.Type).Ref(workbookcategory.Table).Unique().Required().Field(workbook.Label + "_id"),
	}
}
