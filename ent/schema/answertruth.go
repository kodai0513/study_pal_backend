package schema

import (
	"study-pal-backend/ent/answertruth"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AnswerTruth holds the schema definition for the AnswerTruth entity.
type AnswerTruth struct {
	ent.Schema
}

func (AnswerTruth) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the AnswerTruth.
func (AnswerTruth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(problem.Label+"_id", uuid.UUID{}).Unique(),
		field.Bool("truth"),
	}
}

// Edges of the AnswerTruth.
func (AnswerTruth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(problem.Label, Problem.Type).Ref(answertruth.Table).Unique().Required().Field(problem.Label + "_id"),
	}
}
