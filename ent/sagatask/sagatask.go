// Code generated by ent, DO NOT EDIT.

package sagatask

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sagatask type in the database.
	Label = "saga_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// Table holds the table name of the sagatask in the database.
	Table = "saga_tasks"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "saga_tasks"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_saga_step_task"
)

// Columns holds all SQL columns for sagatask fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "saga_tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"node_saga_step_task",
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

// OrderOption defines the ordering options for the SagaTask queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, NodeTable, NodeColumn),
	)
}