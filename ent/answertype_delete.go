// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"study-pal-backend/ent/answertype"
	"study-pal-backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnswerTypeDelete is the builder for deleting a AnswerType entity.
type AnswerTypeDelete struct {
	config
	hooks    []Hook
	mutation *AnswerTypeMutation
}

// Where appends a list predicates to the AnswerTypeDelete builder.
func (atd *AnswerTypeDelete) Where(ps ...predicate.AnswerType) *AnswerTypeDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AnswerTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, atd.sqlExec, atd.mutation, atd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AnswerTypeDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AnswerTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(answertype.Table, sqlgraph.NewFieldSpec(answertype.FieldID, field.TypeUUID))
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	atd.mutation.done = true
	return affected, err
}

// AnswerTypeDeleteOne is the builder for deleting a single AnswerType entity.
type AnswerTypeDeleteOne struct {
	atd *AnswerTypeDelete
}

// Where appends a list predicates to the AnswerTypeDelete builder.
func (atdo *AnswerTypeDeleteOne) Where(ps ...predicate.AnswerType) *AnswerTypeDeleteOne {
	atdo.atd.mutation.Where(ps...)
	return atdo
}

// Exec executes the deletion query.
func (atdo *AnswerTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{answertype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AnswerTypeDeleteOne) ExecX(ctx context.Context) {
	if err := atdo.Exec(ctx); err != nil {
		panic(err)
	}
}
