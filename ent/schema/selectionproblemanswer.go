package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/selectionproblem"
	"study-pal-backend/ent/selectionproblemanswer"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SelectionProblemAnswer holds the schema definition for the SelectionProblemAnswer entity.
type SelectionProblemAnswer struct {
	ent.Schema
}

func (SelectionProblemAnswer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SelectionProblemAnswer.
func (SelectionProblemAnswer) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_correct"),
		field.UUID("selection_problem_id", uuid.UUID{}).Unique(),
		field.String("statement").MaxLen(255).NotEmpty(),
	}
}

// Edges of the SelectionProblemAnswer.
func (SelectionProblemAnswer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(selectionproblem.Label, SelectionProblem.Type).Ref(selectionproblemanswer.Table).Unique().Required().Field("selection_problem_id"),
	}
}
