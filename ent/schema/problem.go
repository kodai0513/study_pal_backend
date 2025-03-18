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

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

func (Problem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int(answertype.Label + "_id"),
		field.String("statement").MaxLen(1000).NotEmpty(),
		field.Int(workbook.Label + "_id"),
		field.Int(workbookcategory.Label + "_id"),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(answertype.Label, AnswerType.Type).Ref(problem.Table).Unique().Required().Field(answertype.Label + "_id"),
		edge.To(answerdescription.Table, AnswerDescription.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(answermultichoices.Table, AnswerMultiChoices.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(answertruth.Table, AnswerTruth.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From(workbook.Label, Workbook.Type).Ref(problem.Table).Unique().Required().Field(workbook.Label + "_id"),
		edge.From(workbookcategory.Label, WorkbookCategory.Type).Ref(problem.Table).Unique().Required().Field(workbookcategory.Label + "_id"),
	}
}
