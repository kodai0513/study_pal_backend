// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-pal-backend/ent/article"
	"study-pal-backend/ent/predicate"
	"study-pal-backend/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeArticle = "Article"
	TypeUser    = "User"
)

// ArticleMutation represents an operation that mutates the Article nodes in the graph.
type ArticleMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	updated_at    *time.Time
	description   *string
	clearedFields map[string]struct{}
	post          *int
	clearedpost   bool
	done          bool
	oldValue      func(context.Context) (*Article, error)
	predicates    []predicate.Article
}

var _ ent.Mutation = (*ArticleMutation)(nil)

// articleOption allows management of the mutation configuration using functional options.
type articleOption func(*ArticleMutation)

// newArticleMutation creates new mutation for the Article entity.
func newArticleMutation(c config, op Op, opts ...articleOption) *ArticleMutation {
	m := &ArticleMutation{
		config:        c,
		op:            op,
		typ:           TypeArticle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArticleID sets the ID field of the mutation.
func withArticleID(id int) articleOption {
	return func(m *ArticleMutation) {
		var (
			err   error
			once  sync.Once
			value *Article
		)
		m.oldValue = func(ctx context.Context) (*Article, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Article.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArticle sets the old Article of the mutation.
func withArticle(node *Article) articleOption {
	return func(m *ArticleMutation) {
		m.oldValue = func(context.Context) (*Article, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArticleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArticleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArticleMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArticleMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Article.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *ArticleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ArticleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ArticleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ArticleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ArticleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ArticleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetDescription sets the "description" field.
func (m *ArticleMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *ArticleMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *ArticleMutation) ResetDescription() {
	m.description = nil
}

// SetPostID sets the "post_id" field.
func (m *ArticleMutation) SetPostID(i int) {
	m.post = &i
}

// PostID returns the value of the "post_id" field in the mutation.
func (m *ArticleMutation) PostID() (r int, exists bool) {
	v := m.post
	if v == nil {
		return
	}
	return *v, true
}

// OldPostID returns the old "post_id" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldPostID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPostID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPostID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPostID: %w", err)
	}
	return oldValue.PostID, nil
}

// ClearPostID clears the value of the "post_id" field.
func (m *ArticleMutation) ClearPostID() {
	m.post = nil
	m.clearedFields[article.FieldPostID] = struct{}{}
}

// PostIDCleared returns if the "post_id" field was cleared in this mutation.
func (m *ArticleMutation) PostIDCleared() bool {
	_, ok := m.clearedFields[article.FieldPostID]
	return ok
}

// ResetPostID resets all changes to the "post_id" field.
func (m *ArticleMutation) ResetPostID() {
	m.post = nil
	delete(m.clearedFields, article.FieldPostID)
}

// ClearPost clears the "post" edge to the User entity.
func (m *ArticleMutation) ClearPost() {
	m.clearedpost = true
	m.clearedFields[article.FieldPostID] = struct{}{}
}

// PostCleared reports if the "post" edge to the User entity was cleared.
func (m *ArticleMutation) PostCleared() bool {
	return m.PostIDCleared() || m.clearedpost
}

// PostIDs returns the "post" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// PostID instead. It exists only for internal usage by the builders.
func (m *ArticleMutation) PostIDs() (ids []int) {
	if id := m.post; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetPost resets all changes to the "post" edge.
func (m *ArticleMutation) ResetPost() {
	m.post = nil
	m.clearedpost = false
}

// Where appends a list predicates to the ArticleMutation builder.
func (m *ArticleMutation) Where(ps ...predicate.Article) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ArticleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ArticleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Article, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ArticleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ArticleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Article).
func (m *ArticleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArticleMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, article.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, article.FieldUpdatedAt)
	}
	if m.description != nil {
		fields = append(fields, article.FieldDescription)
	}
	if m.post != nil {
		fields = append(fields, article.FieldPostID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArticleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case article.FieldCreatedAt:
		return m.CreatedAt()
	case article.FieldUpdatedAt:
		return m.UpdatedAt()
	case article.FieldDescription:
		return m.Description()
	case article.FieldPostID:
		return m.PostID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArticleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case article.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case article.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case article.FieldDescription:
		return m.OldDescription(ctx)
	case article.FieldPostID:
		return m.OldPostID(ctx)
	}
	return nil, fmt.Errorf("unknown Article field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case article.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case article.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case article.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case article.FieldPostID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPostID(v)
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArticleMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArticleMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Article numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArticleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(article.FieldPostID) {
		fields = append(fields, article.FieldPostID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArticleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArticleMutation) ClearField(name string) error {
	switch name {
	case article.FieldPostID:
		m.ClearPostID()
		return nil
	}
	return fmt.Errorf("unknown Article nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArticleMutation) ResetField(name string) error {
	switch name {
	case article.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case article.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case article.FieldDescription:
		m.ResetDescription()
		return nil
	case article.FieldPostID:
		m.ResetPostID()
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArticleMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.post != nil {
		edges = append(edges, article.EdgePost)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArticleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case article.EdgePost:
		if id := m.post; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArticleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArticleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArticleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedpost {
		edges = append(edges, article.EdgePost)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArticleMutation) EdgeCleared(name string) bool {
	switch name {
	case article.EdgePost:
		return m.clearedpost
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArticleMutation) ClearEdge(name string) error {
	switch name {
	case article.EdgePost:
		m.ClearPost()
		return nil
	}
	return fmt.Errorf("unknown Article unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArticleMutation) ResetEdge(name string) error {
	switch name {
	case article.EdgePost:
		m.ResetPost()
		return nil
	}
	return fmt.Errorf("unknown Article edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op              Op
	typ             string
	id              *int
	created_at      *time.Time
	updated_at      *time.Time
	email           *string
	name            *string
	nick_name       *string
	password        *string
	clearedFields   map[string]struct{}
	articles        map[int]struct{}
	removedarticles map[int]struct{}
	clearedarticles bool
	done            bool
	oldValue        func(context.Context) (*User, error)
	predicates      []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetNickName sets the "nick_name" field.
func (m *UserMutation) SetNickName(s string) {
	m.nick_name = &s
}

// NickName returns the value of the "nick_name" field in the mutation.
func (m *UserMutation) NickName() (r string, exists bool) {
	v := m.nick_name
	if v == nil {
		return
	}
	return *v, true
}

// OldNickName returns the old "nick_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickName: %w", err)
	}
	return oldValue.NickName, nil
}

// ResetNickName resets all changes to the "nick_name" field.
func (m *UserMutation) ResetNickName() {
	m.nick_name = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// AddArticleIDs adds the "articles" edge to the Article entity by ids.
func (m *UserMutation) AddArticleIDs(ids ...int) {
	if m.articles == nil {
		m.articles = make(map[int]struct{})
	}
	for i := range ids {
		m.articles[ids[i]] = struct{}{}
	}
}

// ClearArticles clears the "articles" edge to the Article entity.
func (m *UserMutation) ClearArticles() {
	m.clearedarticles = true
}

// ArticlesCleared reports if the "articles" edge to the Article entity was cleared.
func (m *UserMutation) ArticlesCleared() bool {
	return m.clearedarticles
}

// RemoveArticleIDs removes the "articles" edge to the Article entity by IDs.
func (m *UserMutation) RemoveArticleIDs(ids ...int) {
	if m.removedarticles == nil {
		m.removedarticles = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.articles, ids[i])
		m.removedarticles[ids[i]] = struct{}{}
	}
}

// RemovedArticles returns the removed IDs of the "articles" edge to the Article entity.
func (m *UserMutation) RemovedArticlesIDs() (ids []int) {
	for id := range m.removedarticles {
		ids = append(ids, id)
	}
	return
}

// ArticlesIDs returns the "articles" edge IDs in the mutation.
func (m *UserMutation) ArticlesIDs() (ids []int) {
	for id := range m.articles {
		ids = append(ids, id)
	}
	return
}

// ResetArticles resets all changes to the "articles" edge.
func (m *UserMutation) ResetArticles() {
	m.articles = nil
	m.clearedarticles = false
	m.removedarticles = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.nick_name != nil {
		fields = append(fields, user.FieldNickName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldEmail:
		return m.Email()
	case user.FieldName:
		return m.Name()
	case user.FieldNickName:
		return m.NickName()
	case user.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldNickName:
		return m.OldNickName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldNickName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldNickName:
		m.ResetNickName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.articles != nil {
		edges = append(edges, user.EdgeArticles)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeArticles:
		ids := make([]ent.Value, 0, len(m.articles))
		for id := range m.articles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedarticles != nil {
		edges = append(edges, user.EdgeArticles)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeArticles:
		ids := make([]ent.Value, 0, len(m.removedarticles))
		for id := range m.removedarticles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedarticles {
		edges = append(edges, user.EdgeArticles)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeArticles:
		return m.clearedarticles
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeArticles:
		m.ResetArticles()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
