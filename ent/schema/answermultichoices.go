package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AnswerMultiChoices holds the schema definition for the AnswerMultiChoices entity.
type AnswerMultiChoices struct {
	ent.Schema
}

// Fields of the AnswerMultiChoices.
func (AnswerMultiChoices) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
		field.Bool("is_correct"),
	}
}

// Edges of the AnswerMultiChoices.
func (AnswerMultiChoices) Edges() []ent.Edge {
	return nil
}
