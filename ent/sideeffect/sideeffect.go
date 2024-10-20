// Code generated by ent, DO NOT EDIT.

package sideeffect

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sideeffect type in the database.
	Label = "side_effect"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIdentity holds the string denoting the identity field in the database.
	FieldIdentity = "identity"
	// FieldStepID holds the string denoting the step_id field in the database.
	FieldStepID = "step_id"
	// FieldHandlerName holds the string denoting the handler_name field in the database.
	FieldHandlerName = "handler_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRetryPolicy holds the string denoting the retry_policy field in the database.
	FieldRetryPolicy = "retry_policy"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeExecutions holds the string denoting the executions edge name in mutations.
	EdgeExecutions = "executions"
	// Table holds the table name of the sideeffect in the database.
	Table = "side_effects"
	// ExecutionsTable is the table that holds the executions relation/edge.
	ExecutionsTable = "side_effect_executions"
	// ExecutionsInverseTable is the table name for the SideEffectExecution entity.
	// It exists in this package in order to avoid circular dependency with the "sideeffectexecution" package.
	ExecutionsInverseTable = "side_effect_executions"
	// ExecutionsColumn is the table column denoting the executions relation/edge.
	ExecutionsColumn = "side_effect_executions"
)

// Columns holds all SQL columns for sideeffect fields.
var Columns = []string{
	FieldID,
	FieldIdentity,
	FieldStepID,
	FieldHandlerName,
	FieldStatus,
	FieldRetryPolicy,
	FieldTimeout,
	FieldCreatedAt,
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
	// IdentityValidator is a validator for the "identity" field. It is called by the builders before save.
	IdentityValidator func(string) error
	// StepIDValidator is a validator for the "step_id" field. It is called by the builders before save.
	StepIDValidator func(string) error
	// HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	HandlerNameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending   Status = "Pending"
	StatusRunning   Status = "Running"
	StatusCompleted Status = "Completed"
	StatusFailed    Status = "Failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusRunning, StatusCompleted, StatusFailed:
		return nil
	default:
		return fmt.Errorf("sideeffect: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the SideEffect queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIdentity orders the results by the identity field.
func ByIdentity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentity, opts...).ToFunc()
}

// ByStepID orders the results by the step_id field.
func ByStepID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStepID, opts...).ToFunc()
}

// ByHandlerName orders the results by the handler_name field.
func ByHandlerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHandlerName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTimeout orders the results by the timeout field.
func ByTimeout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeout, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByExecutionsCount orders the results by executions count.
func ByExecutionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExecutionsStep(), opts...)
	}
}

// ByExecutions orders the results by executions terms.
func ByExecutions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExecutionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExecutionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExecutionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExecutionsTable, ExecutionsColumn),
	)
}
