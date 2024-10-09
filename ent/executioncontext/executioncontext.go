// Code generated by ent, DO NOT EDIT.

package executioncontext

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the executioncontext type in the database.
	Label = "execution_context"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCurrentRunID holds the string denoting the current_run_id field in the database.
	FieldCurrentRunID = "current_run_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// EdgeHandlerExecutions holds the string denoting the handler_executions edge name in mutations.
	EdgeHandlerExecutions = "handler_executions"
	// Table holds the table name of the executioncontext in the database.
	Table = "execution_contexts"
	// HandlerExecutionsTable is the table that holds the handler_executions relation/edge.
	HandlerExecutionsTable = "handler_executions"
	// HandlerExecutionsInverseTable is the table name for the HandlerExecution entity.
	// It exists in this package in order to avoid circular dependency with the "handlerexecution" package.
	HandlerExecutionsInverseTable = "handler_executions"
	// HandlerExecutionsColumn is the table column denoting the handler_executions relation/edge.
	HandlerExecutionsColumn = "execution_context_handler_executions"
)

// Columns holds all SQL columns for executioncontext fields.
var Columns = []string{
	FieldID,
	FieldCurrentRunID,
	FieldStatus,
	FieldStartTime,
	FieldEndTime,
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

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusRunning, StatusCompleted, StatusFailed:
		return nil
	default:
		return fmt.Errorf("executioncontext: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the ExecutionContext queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCurrentRunID orders the results by the current_run_id field.
func ByCurrentRunID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentRunID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByHandlerExecutionsCount orders the results by handler_executions count.
func ByHandlerExecutionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHandlerExecutionsStep(), opts...)
	}
}

// ByHandlerExecutions orders the results by handler_executions terms.
func ByHandlerExecutions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHandlerExecutionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHandlerExecutionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HandlerExecutionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HandlerExecutionsTable, HandlerExecutionsColumn),
	)
}
