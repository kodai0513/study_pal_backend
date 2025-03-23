package schema

import (
	"study-pal-backend/ent/answermultichoices"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AnswerMultiChoices holds the schema definition for the AnswerMultiChoices entity.
type AnswerMultiChoices struct {
	ent.Schema
}

func (AnswerMultiChoices) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the AnswerMultiChoices.
func (AnswerMultiChoices) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
		field.Bool("is_correct"),
		field.UUID(problem.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the AnswerMultiChoices.
func (AnswerMultiChoices) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(problem.Label, Problem.Type).Ref(answermultichoices.Table).Unique().Required().Field(problem.Label + "_id"),
	}
}
