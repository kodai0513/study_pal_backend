// Code generated by ent, DO NOT EDIT.

package workbookcategoryclosure

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the workbookcategoryclosure type in the database.
	Label = "workbook_category_closure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldWorkbookID holds the string denoting the workbook_id field in the database.
	FieldWorkbookID = "workbook_id"
	// FieldChildID holds the string denoting the child_id field in the database.
	FieldChildID = "child_id"
	// FieldIsRoot holds the string denoting the is_root field in the database.
	FieldIsRoot = "is_root"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// EdgeChildCategory holds the string denoting the child_category edge name in mutations.
	EdgeChildCategory = "child_category"
	// EdgeParentCategory holds the string denoting the parent_category edge name in mutations.
	EdgeParentCategory = "parent_category"
	// Table holds the table name of the workbookcategoryclosure in the database.
	Table = "workbook_category_closures"
	// ChildCategoryTable is the table that holds the child_category relation/edge.
	ChildCategoryTable = "workbook_category_closures"
	// ChildCategoryInverseTable is the table name for the WorkbookCategory entity.
	// It exists in this package in order to avoid circular dependency with the "workbookcategory" package.
	ChildCategoryInverseTable = "workbook_categories"
	// ChildCategoryColumn is the table column denoting the child_category relation/edge.
	ChildCategoryColumn = "child_id"
	// ParentCategoryTable is the table that holds the parent_category relation/edge.
	ParentCategoryTable = "workbook_category_closures"
	// ParentCategoryInverseTable is the table name for the WorkbookCategory entity.
	// It exists in this package in order to avoid circular dependency with the "workbookcategory" package.
	ParentCategoryInverseTable = "workbook_categories"
	// ParentCategoryColumn is the table column denoting the parent_category relation/edge.
	ParentCategoryColumn = "parent_id"
)

// Columns holds all SQL columns for workbookcategoryclosure fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldWorkbookID,
	FieldChildID,
	FieldIsRoot,
	FieldParentID,
	FieldPosition,
	FieldLevel,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the WorkbookCategoryClosure queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByWorkbookID orders the results by the workbook_id field.
func ByWorkbookID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkbookID, opts...).ToFunc()
}

// ByChildID orders the results by the child_id field.
func ByChildID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChildID, opts...).ToFunc()
}

// ByIsRoot orders the results by the is_root field.
func ByIsRoot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRoot, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByPosition orders the results by the position field.
func ByPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPosition, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByChildCategoryField orders the results by child_category field.
func ByChildCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentCategoryField orders the results by parent_category field.
func ByParentCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentCategoryStep(), sql.OrderByField(field, opts...))
	}
}
func newChildCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChildCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ChildCategoryTable, ChildCategoryColumn),
	)
}
func newParentCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParentCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ParentCategoryTable, ParentCategoryColumn),
	)
}
