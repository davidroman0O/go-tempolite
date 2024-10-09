// Code generated by ent, DO NOT EDIT.

package sagatransaction

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sagatransaction type in the database.
	Label = "saga_transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldNextTransactionName holds the string denoting the next_transaction_name field in the database.
	FieldNextTransactionName = "next_transaction_name"
	// FieldFailureCompensationName holds the string denoting the failure_compensation_name field in the database.
	FieldFailureCompensationName = "failure_compensation_name"
	// EdgeExecutionUnit holds the string denoting the execution_unit edge name in mutations.
	EdgeExecutionUnit = "execution_unit"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// EdgeCompensation holds the string denoting the compensation edge name in mutations.
	EdgeCompensation = "compensation"
	// Table holds the table name of the sagatransaction in the database.
	Table = "saga_transactions"
	// ExecutionUnitTable is the table that holds the execution_unit relation/edge.
	ExecutionUnitTable = "saga_transactions"
	// ExecutionUnitInverseTable is the table name for the ExecutionUnit entity.
	// It exists in this package in order to avoid circular dependency with the "executionunit" package.
	ExecutionUnitInverseTable = "execution_units"
	// ExecutionUnitColumn is the table column denoting the execution_unit relation/edge.
	ExecutionUnitColumn = "execution_unit_saga_transactions"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "saga_transactions"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "saga_transaction_task"
	// CompensationTable is the table that holds the compensation relation/edge.
	CompensationTable = "saga_compensations"
	// CompensationInverseTable is the table name for the SagaCompensation entity.
	// It exists in this package in order to avoid circular dependency with the "sagacompensation" package.
	CompensationInverseTable = "saga_compensations"
	// CompensationColumn is the table column denoting the compensation relation/edge.
	CompensationColumn = "saga_transaction_compensation"
)

// Columns holds all SQL columns for sagatransaction fields.
var Columns = []string{
	FieldID,
	FieldOrder,
	FieldNextTransactionName,
	FieldFailureCompensationName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "saga_transactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"execution_unit_saga_transactions",
	"saga_transaction_task",
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

// OrderOption defines the ordering options for the SagaTransaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByNextTransactionName orders the results by the next_transaction_name field.
func ByNextTransactionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextTransactionName, opts...).ToFunc()
}

// ByFailureCompensationName orders the results by the failure_compensation_name field.
func ByFailureCompensationName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFailureCompensationName, opts...).ToFunc()
}

// ByExecutionUnitField orders the results by execution_unit field.
func ByExecutionUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExecutionUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByTaskField orders the results by task field.
func ByTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompensationField orders the results by compensation field.
func ByCompensationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompensationStep(), sql.OrderByField(field, opts...))
	}
}
func newExecutionUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExecutionUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExecutionUnitTable, ExecutionUnitColumn),
	)
}
func newTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TaskTable, TaskColumn),
	)
}
func newCompensationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompensationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CompensationTable, CompensationColumn),
	)
}
