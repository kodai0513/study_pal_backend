// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-pal-backend/ent/descriptionproblem"
	"study-pal-backend/ent/selectionproblem"
	"study-pal-backend/ent/trueorfalseproblem"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookcategory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookCategoryCreate is the builder for creating a WorkbookCategory entity.
type WorkbookCategoryCreate struct {
	config
	mutation *WorkbookCategoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wcc *WorkbookCategoryCreate) SetCreatedAt(t time.Time) *WorkbookCategoryCreate {
	wcc.mutation.SetCreatedAt(t)
	return wcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wcc *WorkbookCategoryCreate) SetNillableCreatedAt(t *time.Time) *WorkbookCategoryCreate {
	if t != nil {
		wcc.SetCreatedAt(*t)
	}
	return wcc
}

// SetUpdatedAt sets the "updated_at" field.
func (wcc *WorkbookCategoryCreate) SetUpdatedAt(t time.Time) *WorkbookCategoryCreate {
	wcc.mutation.SetUpdatedAt(t)
	return wcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wcc *WorkbookCategoryCreate) SetNillableUpdatedAt(t *time.Time) *WorkbookCategoryCreate {
	if t != nil {
		wcc.SetUpdatedAt(*t)
	}
	return wcc
}

// SetName sets the "name" field.
func (wcc *WorkbookCategoryCreate) SetName(s string) *WorkbookCategoryCreate {
	wcc.mutation.SetName(s)
	return wcc
}

// SetWorkbookID sets the "workbook_id" field.
func (wcc *WorkbookCategoryCreate) SetWorkbookID(u uuid.UUID) *WorkbookCategoryCreate {
	wcc.mutation.SetWorkbookID(u)
	return wcc
}

// SetID sets the "id" field.
func (wcc *WorkbookCategoryCreate) SetID(u uuid.UUID) *WorkbookCategoryCreate {
	wcc.mutation.SetID(u)
	return wcc
}

// AddDescriptionProblemIDs adds the "description_problems" edge to the DescriptionProblem entity by IDs.
func (wcc *WorkbookCategoryCreate) AddDescriptionProblemIDs(ids ...uuid.UUID) *WorkbookCategoryCreate {
	wcc.mutation.AddDescriptionProblemIDs(ids...)
	return wcc
}

// AddDescriptionProblems adds the "description_problems" edges to the DescriptionProblem entity.
func (wcc *WorkbookCategoryCreate) AddDescriptionProblems(d ...*DescriptionProblem) *WorkbookCategoryCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return wcc.AddDescriptionProblemIDs(ids...)
}

// AddSelectionProblemIDs adds the "selection_problems" edge to the SelectionProblem entity by IDs.
func (wcc *WorkbookCategoryCreate) AddSelectionProblemIDs(ids ...uuid.UUID) *WorkbookCategoryCreate {
	wcc.mutation.AddSelectionProblemIDs(ids...)
	return wcc
}

// AddSelectionProblems adds the "selection_problems" edges to the SelectionProblem entity.
func (wcc *WorkbookCategoryCreate) AddSelectionProblems(s ...*SelectionProblem) *WorkbookCategoryCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wcc.AddSelectionProblemIDs(ids...)
}

// AddTrueOrFalseProblemIDs adds the "true_or_false_problems" edge to the TrueOrFalseProblem entity by IDs.
func (wcc *WorkbookCategoryCreate) AddTrueOrFalseProblemIDs(ids ...uuid.UUID) *WorkbookCategoryCreate {
	wcc.mutation.AddTrueOrFalseProblemIDs(ids...)
	return wcc
}

// AddTrueOrFalseProblems adds the "true_or_false_problems" edges to the TrueOrFalseProblem entity.
func (wcc *WorkbookCategoryCreate) AddTrueOrFalseProblems(t ...*TrueOrFalseProblem) *WorkbookCategoryCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wcc.AddTrueOrFalseProblemIDs(ids...)
}

// SetWorkbook sets the "workbook" edge to the Workbook entity.
func (wcc *WorkbookCategoryCreate) SetWorkbook(w *Workbook) *WorkbookCategoryCreate {
	return wcc.SetWorkbookID(w.ID)
}

// Mutation returns the WorkbookCategoryMutation object of the builder.
func (wcc *WorkbookCategoryCreate) Mutation() *WorkbookCategoryMutation {
	return wcc.mutation
}

// Save creates the WorkbookCategory in the database.
func (wcc *WorkbookCategoryCreate) Save(ctx context.Context) (*WorkbookCategory, error) {
	wcc.defaults()
	return withHooks(ctx, wcc.sqlSave, wcc.mutation, wcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wcc *WorkbookCategoryCreate) SaveX(ctx context.Context) *WorkbookCategory {
	v, err := wcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcc *WorkbookCategoryCreate) Exec(ctx context.Context) error {
	_, err := wcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcc *WorkbookCategoryCreate) ExecX(ctx context.Context) {
	if err := wcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcc *WorkbookCategoryCreate) defaults() {
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		v := workbookcategory.DefaultCreatedAt()
		wcc.mutation.SetCreatedAt(v)
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		v := workbookcategory.DefaultUpdatedAt()
		wcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcc *WorkbookCategoryCreate) check() error {
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WorkbookCategory.created_at"`)}
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WorkbookCategory.updated_at"`)}
	}
	if _, ok := wcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WorkbookCategory.name"`)}
	}
	if v, ok := wcc.mutation.Name(); ok {
		if err := workbookcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WorkbookCategory.name": %w`, err)}
		}
	}
	if _, ok := wcc.mutation.WorkbookID(); !ok {
		return &ValidationError{Name: "workbook_id", err: errors.New(`ent: missing required field "WorkbookCategory.workbook_id"`)}
	}
	if len(wcc.mutation.WorkbookIDs()) == 0 {
		return &ValidationError{Name: "workbook", err: errors.New(`ent: missing required edge "WorkbookCategory.workbook"`)}
	}
	return nil
}

func (wcc *WorkbookCategoryCreate) sqlSave(ctx context.Context) (*WorkbookCategory, error) {
	if err := wcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wcc.mutation.id = &_node.ID
	wcc.mutation.done = true
	return _node, nil
}

func (wcc *WorkbookCategoryCreate) createSpec() (*WorkbookCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkbookCategory{config: wcc.config}
		_spec = sqlgraph.NewCreateSpec(workbookcategory.Table, sqlgraph.NewFieldSpec(workbookcategory.FieldID, field.TypeUUID))
	)
	if id, ok := wcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wcc.mutation.CreatedAt(); ok {
		_spec.SetField(workbookcategory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wcc.mutation.UpdatedAt(); ok {
		_spec.SetField(workbookcategory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wcc.mutation.Name(); ok {
		_spec.SetField(workbookcategory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := wcc.mutation.DescriptionProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workbookcategory.DescriptionProblemsTable,
			Columns: []string{workbookcategory.DescriptionProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(descriptionproblem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wcc.mutation.SelectionProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workbookcategory.SelectionProblemsTable,
			Columns: []string{workbookcategory.SelectionProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(selectionproblem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wcc.mutation.TrueOrFalseProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workbookcategory.TrueOrFalseProblemsTable,
			Columns: []string{workbookcategory.TrueOrFalseProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trueorfalseproblem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wcc.mutation.WorkbookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workbookcategory.WorkbookTable,
			Columns: []string{workbookcategory.WorkbookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workbook.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WorkbookID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkbookCategoryCreateBulk is the builder for creating many WorkbookCategory entities in bulk.
type WorkbookCategoryCreateBulk struct {
	config
	err      error
	builders []*WorkbookCategoryCreate
}

// Save creates the WorkbookCategory entities in the database.
func (wccb *WorkbookCategoryCreateBulk) Save(ctx context.Context) ([]*WorkbookCategory, error) {
	if wccb.err != nil {
		return nil, wccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wccb.builders))
	nodes := make([]*WorkbookCategory, len(wccb.builders))
	mutators := make([]Mutator, len(wccb.builders))
	for i := range wccb.builders {
		func(i int, root context.Context) {
			builder := wccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkbookCategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wccb *WorkbookCategoryCreateBulk) SaveX(ctx context.Context) []*WorkbookCategory {
	v, err := wccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wccb *WorkbookCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := wccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wccb *WorkbookCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := wccb.Exec(ctx); err != nil {
		panic(err)
	}
}
