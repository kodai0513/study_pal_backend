// Code generated by ent, DO NOT EDIT.

package selectionproblemanswer

import (
	"study-pal-backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsCorrect applies equality check predicate on the "is_correct" field. It's identical to IsCorrectEQ.
func IsCorrect(v bool) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldIsCorrect, v))
}

// SelectionProblemID applies equality check predicate on the "selection_problem_id" field. It's identical to SelectionProblemIDEQ.
func SelectionProblemID(v uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldSelectionProblemID, v))
}

// Statement applies equality check predicate on the "statement" field. It's identical to StatementEQ.
func Statement(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldStatement, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsCorrectEQ applies the EQ predicate on the "is_correct" field.
func IsCorrectEQ(v bool) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldIsCorrect, v))
}

// IsCorrectNEQ applies the NEQ predicate on the "is_correct" field.
func IsCorrectNEQ(v bool) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNEQ(FieldIsCorrect, v))
}

// SelectionProblemIDEQ applies the EQ predicate on the "selection_problem_id" field.
func SelectionProblemIDEQ(v uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldSelectionProblemID, v))
}

// SelectionProblemIDNEQ applies the NEQ predicate on the "selection_problem_id" field.
func SelectionProblemIDNEQ(v uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNEQ(FieldSelectionProblemID, v))
}

// SelectionProblemIDIn applies the In predicate on the "selection_problem_id" field.
func SelectionProblemIDIn(vs ...uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldIn(FieldSelectionProblemID, vs...))
}

// SelectionProblemIDNotIn applies the NotIn predicate on the "selection_problem_id" field.
func SelectionProblemIDNotIn(vs ...uuid.UUID) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNotIn(FieldSelectionProblemID, vs...))
}

// StatementEQ applies the EQ predicate on the "statement" field.
func StatementEQ(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEQ(FieldStatement, v))
}

// StatementNEQ applies the NEQ predicate on the "statement" field.
func StatementNEQ(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNEQ(FieldStatement, v))
}

// StatementIn applies the In predicate on the "statement" field.
func StatementIn(vs ...string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldIn(FieldStatement, vs...))
}

// StatementNotIn applies the NotIn predicate on the "statement" field.
func StatementNotIn(vs ...string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldNotIn(FieldStatement, vs...))
}

// StatementGT applies the GT predicate on the "statement" field.
func StatementGT(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGT(FieldStatement, v))
}

// StatementGTE applies the GTE predicate on the "statement" field.
func StatementGTE(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldGTE(FieldStatement, v))
}

// StatementLT applies the LT predicate on the "statement" field.
func StatementLT(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLT(FieldStatement, v))
}

// StatementLTE applies the LTE predicate on the "statement" field.
func StatementLTE(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldLTE(FieldStatement, v))
}

// StatementContains applies the Contains predicate on the "statement" field.
func StatementContains(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldContains(FieldStatement, v))
}

// StatementHasPrefix applies the HasPrefix predicate on the "statement" field.
func StatementHasPrefix(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldHasPrefix(FieldStatement, v))
}

// StatementHasSuffix applies the HasSuffix predicate on the "statement" field.
func StatementHasSuffix(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldHasSuffix(FieldStatement, v))
}

// StatementEqualFold applies the EqualFold predicate on the "statement" field.
func StatementEqualFold(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldEqualFold(FieldStatement, v))
}

// StatementContainsFold applies the ContainsFold predicate on the "statement" field.
func StatementContainsFold(v string) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.FieldContainsFold(FieldStatement, v))
}

// HasSelectionProblem applies the HasEdge predicate on the "selection_problem" edge.
func HasSelectionProblem() predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SelectionProblemTable, SelectionProblemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSelectionProblemWith applies the HasEdge predicate on the "selection_problem" edge with a given conditions (other predicates).
func HasSelectionProblemWith(preds ...predicate.SelectionProblem) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(func(s *sql.Selector) {
		step := newSelectionProblemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SelectionProblemAnswer) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SelectionProblemAnswer) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SelectionProblemAnswer) predicate.SelectionProblemAnswer {
	return predicate.SelectionProblemAnswer(sql.NotPredicates(p))
}
