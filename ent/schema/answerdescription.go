package schema

import (
	"study-pal-backend/ent/answerdescription"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/problem"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AnswerDescription holds the schema definition for the AnswerDescription entity.
type AnswerDescription struct {
	ent.Schema
}

func (AnswerDescription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the AnswerDescription.
func (AnswerDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
		field.UUID(problem.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the AnswerDescription.
func (AnswerDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(problem.Label, Problem.Type).Ref(answerdescription.Table).Unique().Required().Field(problem.Label + "_id"),
	}
}
