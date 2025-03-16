package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AnswerDescription holds the schema definition for the AnswerDescription entity.
type AnswerDescription struct {
	ent.Schema
}

// Fields of the AnswerDescription.
func (AnswerDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty(),
	}
}

// Edges of the AnswerDescription.
func (AnswerDescription) Edges() []ent.Edge {
	return nil
}
