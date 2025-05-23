// Code generated by ent, DO NOT EDIT.

package trueorfalseproblem

import (
	"study-pal-backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsCorrect applies equality check predicate on the "is_correct" field. It's identical to IsCorrectEQ.
func IsCorrect(v bool) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldIsCorrect, v))
}

// Statement applies equality check predicate on the "statement" field. It's identical to StatementEQ.
func Statement(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldStatement, v))
}

// WorkbookID applies equality check predicate on the "workbook_id" field. It's identical to WorkbookIDEQ.
func WorkbookID(v uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldWorkbookID, v))
}

// WorkbookCategoryID applies equality check predicate on the "workbook_category_id" field. It's identical to WorkbookCategoryIDEQ.
func WorkbookCategoryID(v uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldWorkbookCategoryID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsCorrectEQ applies the EQ predicate on the "is_correct" field.
func IsCorrectEQ(v bool) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldIsCorrect, v))
}

// IsCorrectNEQ applies the NEQ predicate on the "is_correct" field.
func IsCorrectNEQ(v bool) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldIsCorrect, v))
}

// StatementEQ applies the EQ predicate on the "statement" field.
func StatementEQ(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldStatement, v))
}

// StatementNEQ applies the NEQ predicate on the "statement" field.
func StatementNEQ(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldStatement, v))
}

// StatementIn applies the In predicate on the "statement" field.
func StatementIn(vs ...string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIn(FieldStatement, vs...))
}

// StatementNotIn applies the NotIn predicate on the "statement" field.
func StatementNotIn(vs ...string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotIn(FieldStatement, vs...))
}

// StatementGT applies the GT predicate on the "statement" field.
func StatementGT(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGT(FieldStatement, v))
}

// StatementGTE applies the GTE predicate on the "statement" field.
func StatementGTE(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldGTE(FieldStatement, v))
}

// StatementLT applies the LT predicate on the "statement" field.
func StatementLT(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLT(FieldStatement, v))
}

// StatementLTE applies the LTE predicate on the "statement" field.
func StatementLTE(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldLTE(FieldStatement, v))
}

// StatementContains applies the Contains predicate on the "statement" field.
func StatementContains(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldContains(FieldStatement, v))
}

// StatementHasPrefix applies the HasPrefix predicate on the "statement" field.
func StatementHasPrefix(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldHasPrefix(FieldStatement, v))
}

// StatementHasSuffix applies the HasSuffix predicate on the "statement" field.
func StatementHasSuffix(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldHasSuffix(FieldStatement, v))
}

// StatementEqualFold applies the EqualFold predicate on the "statement" field.
func StatementEqualFold(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEqualFold(FieldStatement, v))
}

// StatementContainsFold applies the ContainsFold predicate on the "statement" field.
func StatementContainsFold(v string) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldContainsFold(FieldStatement, v))
}

// WorkbookIDEQ applies the EQ predicate on the "workbook_id" field.
func WorkbookIDEQ(v uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldWorkbookID, v))
}

// WorkbookIDNEQ applies the NEQ predicate on the "workbook_id" field.
func WorkbookIDNEQ(v uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldWorkbookID, v))
}

// WorkbookIDIn applies the In predicate on the "workbook_id" field.
func WorkbookIDIn(vs ...uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIn(FieldWorkbookID, vs...))
}

// WorkbookIDNotIn applies the NotIn predicate on the "workbook_id" field.
func WorkbookIDNotIn(vs ...uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotIn(FieldWorkbookID, vs...))
}

// WorkbookCategoryIDEQ applies the EQ predicate on the "workbook_category_id" field.
func WorkbookCategoryIDEQ(v uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldEQ(FieldWorkbookCategoryID, v))
}

// WorkbookCategoryIDNEQ applies the NEQ predicate on the "workbook_category_id" field.
func WorkbookCategoryIDNEQ(v uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNEQ(FieldWorkbookCategoryID, v))
}

// WorkbookCategoryIDIn applies the In predicate on the "workbook_category_id" field.
func WorkbookCategoryIDIn(vs ...uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIn(FieldWorkbookCategoryID, vs...))
}

// WorkbookCategoryIDNotIn applies the NotIn predicate on the "workbook_category_id" field.
func WorkbookCategoryIDNotIn(vs ...uuid.UUID) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotIn(FieldWorkbookCategoryID, vs...))
}

// WorkbookCategoryIDIsNil applies the IsNil predicate on the "workbook_category_id" field.
func WorkbookCategoryIDIsNil() predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldIsNull(FieldWorkbookCategoryID))
}

// WorkbookCategoryIDNotNil applies the NotNil predicate on the "workbook_category_id" field.
func WorkbookCategoryIDNotNil() predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.FieldNotNull(FieldWorkbookCategoryID))
}

// HasWorkbook applies the HasEdge predicate on the "workbook" edge.
func HasWorkbook() predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkbookTable, WorkbookColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkbookWith applies the HasEdge predicate on the "workbook" edge with a given conditions (other predicates).
func HasWorkbookWith(preds ...predicate.Workbook) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(func(s *sql.Selector) {
		step := newWorkbookStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkbookCategory applies the HasEdge predicate on the "workbook_category" edge.
func HasWorkbookCategory() predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkbookCategoryTable, WorkbookCategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkbookCategoryWith applies the HasEdge predicate on the "workbook_category" edge with a given conditions (other predicates).
func HasWorkbookCategoryWith(preds ...predicate.WorkbookCategory) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(func(s *sql.Selector) {
		step := newWorkbookCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TrueOrFalseProblem) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TrueOrFalseProblem) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TrueOrFalseProblem) predicate.TrueOrFalseProblem {
	return predicate.TrueOrFalseProblem(sql.NotPredicates(p))
}
