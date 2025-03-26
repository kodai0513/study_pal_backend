package schema

import (
	"study-pal-backend/ent/answerdescription"
	"study-pal-backend/ent/answermultichoices"
	"study-pal-backend/ent/answertruth"
	"study-pal-backend/ent/answertype"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookcategory"
	"study-pal-backend/ent/workbookcategorydetail"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

func (Problem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(answertype.Label+"_id", uuid.UUID{}).Unique(),
		field.String("statement").MaxLen(1000).NotEmpty(),
		field.UUID(workbook.Label+"_id", uuid.UUID{}).Unique(),
		field.UUID(workbookcategory.Label+"_id", uuid.UUID{}).Nillable().Optional().Unique(),
		field.UUID(workbookcategorydetail.Label+"_id", uuid.UUID{}).Nillable().Optional().Unique(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(answertype.Label, AnswerType.Type).Ref(problem.Table).Unique().Required().Field(answertype.Label + "_id"),
		edge.To(answerdescription.Table, AnswerDescription.Type).Unique().Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(answermultichoices.Table, AnswerMultiChoices.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(answertruth.Table, AnswerTruth.Type).Unique().Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From(workbook.Label, Workbook.Type).Ref(problem.Table).Unique().Required().Field(workbook.Label + "_id"),
		edge.From(workbookcategory.Label, WorkbookCategory.Type).Ref(problem.Table).Unique().Field(workbookcategory.Label + "_id"),
		edge.From(workbookcategorydetail.Label, WorkbookCategoryDetail.Type).Ref(problem.Table).Unique().Field(workbookcategorydetail.Label + "_id"),
	}
}
