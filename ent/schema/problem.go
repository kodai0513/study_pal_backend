package schema

import (
	"study-pal-backend/ent/answerdescription"
	"study-pal-backend/ent/answermultichoices"
	"study-pal-backend/ent/answertruth"
	"study-pal-backend/ent/answertype"
	"study-pal-backend/ent/problem"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int(answertype.Label + "_id"),
		field.String("statement").MaxLen(1000).NotEmpty(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(answertype.Label, AnswerType.Type).Ref(problem.Table).Unique().Required().Field(answertype.Label + "_id"),
		edge.To(answerdescription.Table, AnswerDescription.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(answermultichoices.Table, AnswerMultiChoices.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To(answertruth.Table, AnswerTruth.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
