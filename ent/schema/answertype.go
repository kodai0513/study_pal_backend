package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AnswerType holds the schema definition for the AnswerType entity.
type AnswerType struct {
	ent.Schema
}

func (AnswerType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AnswerType.
func (AnswerType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).Unique().NotEmpty(),
	}
}

// Edges of the AnswerType.
func (AnswerType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(problem.Table, Problem.Type).Annotations(entsql.OnDelete(entsql.Restrict)),
	}
}
