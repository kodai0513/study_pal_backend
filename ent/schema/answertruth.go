package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AnswerTruth holds the schema definition for the AnswerTruth entity.
type AnswerTruth struct {
	ent.Schema
}

// Fields of the AnswerTruth.
func (AnswerTruth) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("truth"),
	}
}

// Edges of the AnswerTruth.
func (AnswerTruth) Edges() []ent.Edge {
	return nil
}
