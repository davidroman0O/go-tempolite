// Code generated by ent, DO NOT EDIT.

package node

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeHandlerTask holds the string denoting the handler_task edge name in mutations.
	EdgeHandlerTask = "handler_task"
	// EdgeSagaStepTask holds the string denoting the saga_step_task edge name in mutations.
	EdgeSagaStepTask = "saga_step_task"
	// EdgeSideEffectTask holds the string denoting the side_effect_task edge name in mutations.
	EdgeSideEffectTask = "side_effect_task"
	// EdgeCompensationTask holds the string denoting the compensation_task edge name in mutations.
	EdgeCompensationTask = "compensation_task"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "nodes"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "node_children"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "nodes"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "node_children"
	// HandlerTaskTable is the table that holds the handler_task relation/edge.
	HandlerTaskTable = "handler_tasks"
	// HandlerTaskInverseTable is the table name for the HandlerTask entity.
	// It exists in this package in order to avoid circular dependency with the "handlertask" package.
	HandlerTaskInverseTable = "handler_tasks"
	// HandlerTaskColumn is the table column denoting the handler_task relation/edge.
	HandlerTaskColumn = "node_handler_task"
	// SagaStepTaskTable is the table that holds the saga_step_task relation/edge.
	SagaStepTaskTable = "saga_tasks"
	// SagaStepTaskInverseTable is the table name for the SagaTask entity.
	// It exists in this package in order to avoid circular dependency with the "sagatask" package.
	SagaStepTaskInverseTable = "saga_tasks"
	// SagaStepTaskColumn is the table column denoting the saga_step_task relation/edge.
	SagaStepTaskColumn = "node_saga_step_task"
	// SideEffectTaskTable is the table that holds the side_effect_task relation/edge.
	SideEffectTaskTable = "side_effect_tasks"
	// SideEffectTaskInverseTable is the table name for the SideEffectTask entity.
	// It exists in this package in order to avoid circular dependency with the "sideeffecttask" package.
	SideEffectTaskInverseTable = "side_effect_tasks"
	// SideEffectTaskColumn is the table column denoting the side_effect_task relation/edge.
	SideEffectTaskColumn = "node_side_effect_task"
	// CompensationTaskTable is the table that holds the compensation_task relation/edge.
	CompensationTaskTable = "compensation_tasks"
	// CompensationTaskInverseTable is the table name for the CompensationTask entity.
	// It exists in this package in order to avoid circular dependency with the "compensationtask" package.
	CompensationTaskInverseTable = "compensation_tasks"
	// CompensationTaskColumn is the table column denoting the compensation_task relation/edge.
	CompensationTaskColumn = "node_compensation_task"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldIndex,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "nodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"node_children",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Node queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByHandlerTaskField orders the results by handler_task field.
func ByHandlerTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHandlerTaskStep(), sql.OrderByField(field, opts...))
	}
}

// BySagaStepTaskField orders the results by saga_step_task field.
func BySagaStepTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSagaStepTaskStep(), sql.OrderByField(field, opts...))
	}
}

// BySideEffectTaskField orders the results by side_effect_task field.
func BySideEffectTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSideEffectTaskStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompensationTaskField orders the results by compensation_task field.
func ByCompensationTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompensationTaskStep(), sql.OrderByField(field, opts...))
	}
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newHandlerTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HandlerTaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, HandlerTaskTable, HandlerTaskColumn),
	)
}
func newSagaStepTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SagaStepTaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SagaStepTaskTable, SagaStepTaskColumn),
	)
}
func newSideEffectTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SideEffectTaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SideEffectTaskTable, SideEffectTaskColumn),
	)
}
func newCompensationTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompensationTaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CompensationTaskTable, CompensationTaskColumn),
	)
}
