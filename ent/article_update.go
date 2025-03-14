// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-pal-backend/ent/article"
	"study-pal-backend/ent/predicate"
	"study-pal-backend/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ArticleUpdate) SetCreatedAt(t time.Time) *ArticleUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCreatedAt(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ArticleUpdate) SetUpdatedAt(t time.Time) *ArticleUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableUpdatedAt(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// SetDescription sets the "description" field.
func (au *ArticleUpdate) SetDescription(s string) *ArticleUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableDescription(s *string) *ArticleUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// SetPostID sets the "post_id" field.
func (au *ArticleUpdate) SetPostID(i int) *ArticleUpdate {
	au.mutation.SetPostID(i)
	return au
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (au *ArticleUpdate) SetNillablePostID(i *int) *ArticleUpdate {
	if i != nil {
		au.SetPostID(*i)
	}
	return au
}

// ClearPostID clears the value of the "post_id" field.
func (au *ArticleUpdate) ClearPostID() *ArticleUpdate {
	au.mutation.ClearPostID()
	return au
}

// SetPost sets the "post" edge to the User entity.
func (au *ArticleUpdate) SetPost(u *User) *ArticleUpdate {
	return au.SetPostID(u.ID)
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// ClearPost clears the "post" edge to the User entity.
func (au *ArticleUpdate) ClearPost() *ArticleUpdate {
	au.mutation.ClearPost()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArticleUpdate) check() error {
	if v, ok := au.mutation.Description(); ok {
		if err := article.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Article.description": %w`, err)}
		}
	}
	return nil
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
	}
	if au.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.PostTable,
			Columns: []string{article.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.PostTable,
			Columns: []string{article.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetCreatedAt sets the "created_at" field.
func (auo *ArticleUpdateOne) SetCreatedAt(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCreatedAt(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ArticleUpdateOne) SetUpdatedAt(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableUpdatedAt(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// SetDescription sets the "description" field.
func (auo *ArticleUpdateOne) SetDescription(s string) *ArticleUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableDescription(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// SetPostID sets the "post_id" field.
func (auo *ArticleUpdateOne) SetPostID(i int) *ArticleUpdateOne {
	auo.mutation.SetPostID(i)
	return auo
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillablePostID(i *int) *ArticleUpdateOne {
	if i != nil {
		auo.SetPostID(*i)
	}
	return auo
}

// ClearPostID clears the value of the "post_id" field.
func (auo *ArticleUpdateOne) ClearPostID() *ArticleUpdateOne {
	auo.mutation.ClearPostID()
	return auo
}

// SetPost sets the "post" edge to the User entity.
func (auo *ArticleUpdateOne) SetPost(u *User) *ArticleUpdateOne {
	return auo.SetPostID(u.ID)
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// ClearPost clears the "post" edge to the User entity.
func (auo *ArticleUpdateOne) ClearPost() *ArticleUpdateOne {
	auo.mutation.ClearPost()
	return auo
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArticleUpdateOne) check() error {
	if v, ok := auo.mutation.Description(); ok {
		if err := article.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Article.description": %w`, err)}
		}
	}
	return nil
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
	}
	if auo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.PostTable,
			Columns: []string{article.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.PostTable,
			Columns: []string{article.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
