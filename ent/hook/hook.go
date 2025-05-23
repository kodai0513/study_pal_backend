// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"study-pal-backend/ent"
)

// The ArticleFunc type is an adapter to allow the use of ordinary
// function as Article mutator.
type ArticleFunc func(context.Context, *ent.ArticleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArticleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArticleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArticleMutation", m)
}

// The ArticleLikeFunc type is an adapter to allow the use of ordinary
// function as ArticleLike mutator.
type ArticleLikeFunc func(context.Context, *ent.ArticleLikeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArticleLikeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArticleLikeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArticleLikeMutation", m)
}

// The DescriptionProblemFunc type is an adapter to allow the use of ordinary
// function as DescriptionProblem mutator.
type DescriptionProblemFunc func(context.Context, *ent.DescriptionProblemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DescriptionProblemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DescriptionProblemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DescriptionProblemMutation", m)
}

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *ent.PermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PermissionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionMutation", m)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
}

// The SelectionProblemFunc type is an adapter to allow the use of ordinary
// function as SelectionProblem mutator.
type SelectionProblemFunc func(context.Context, *ent.SelectionProblemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SelectionProblemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SelectionProblemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SelectionProblemMutation", m)
}

// The SelectionProblemAnswerFunc type is an adapter to allow the use of ordinary
// function as SelectionProblemAnswer mutator.
type SelectionProblemAnswerFunc func(context.Context, *ent.SelectionProblemAnswerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SelectionProblemAnswerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SelectionProblemAnswerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SelectionProblemAnswerMutation", m)
}

// The TrueOrFalseProblemFunc type is an adapter to allow the use of ordinary
// function as TrueOrFalseProblem mutator.
type TrueOrFalseProblemFunc func(context.Context, *ent.TrueOrFalseProblemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TrueOrFalseProblemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TrueOrFalseProblemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TrueOrFalseProblemMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The WorkbookFunc type is an adapter to allow the use of ordinary
// function as Workbook mutator.
type WorkbookFunc func(context.Context, *ent.WorkbookMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkbookFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkbookMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkbookMutation", m)
}

// The WorkbookCategoryFunc type is an adapter to allow the use of ordinary
// function as WorkbookCategory mutator.
type WorkbookCategoryFunc func(context.Context, *ent.WorkbookCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkbookCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkbookCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkbookCategoryMutation", m)
}

// The WorkbookCategoryClosureFunc type is an adapter to allow the use of ordinary
// function as WorkbookCategoryClosure mutator.
type WorkbookCategoryClosureFunc func(context.Context, *ent.WorkbookCategoryClosureMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkbookCategoryClosureFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkbookCategoryClosureMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkbookCategoryClosureMutation", m)
}

// The WorkbookInvitationMemberFunc type is an adapter to allow the use of ordinary
// function as WorkbookInvitationMember mutator.
type WorkbookInvitationMemberFunc func(context.Context, *ent.WorkbookInvitationMemberMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkbookInvitationMemberFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkbookInvitationMemberMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkbookInvitationMemberMutation", m)
}

// The WorkbookMemberFunc type is an adapter to allow the use of ordinary
// function as WorkbookMember mutator.
type WorkbookMemberFunc func(context.Context, *ent.WorkbookMemberMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkbookMemberFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkbookMemberMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkbookMemberMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
