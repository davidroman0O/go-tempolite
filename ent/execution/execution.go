// Code generated by ent, DO NOT EDIT.

package execution

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the execution type in the database.
	Label = "execution"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDag holds the string denoting the dag field in the database.
	FieldDag = "dag"
	// EdgeExecutionContext holds the string denoting the execution_context edge name in mutations.
	EdgeExecutionContext = "execution_context"
	// Table holds the table name of the execution in the database.
	Table = "executions"
	// ExecutionContextTable is the table that holds the execution_context relation/edge.
	ExecutionContextTable = "executions"
	// ExecutionContextInverseTable is the table name for the ExecutionContext entity.
	// It exists in this package in order to avoid circular dependency with the "executioncontext" package.
	ExecutionContextInverseTable = "execution_contexts"
	// ExecutionContextColumn is the table column denoting the execution_context relation/edge.
	ExecutionContextColumn = "execution_execution_context"
)

// Columns holds all SQL columns for execution fields.
var Columns = []string{
	FieldID,
	FieldDag,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "executions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"execution_execution_context",
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

// OrderOption defines the ordering options for the Execution queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByExecutionContextField orders the results by execution_context field.
func ByExecutionContextField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExecutionContextStep(), sql.OrderByField(field, opts...))
	}
}
func newExecutionContextStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExecutionContextInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ExecutionContextTable, ExecutionContextColumn),
	)
}
