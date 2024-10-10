// Code generated by ent, DO NOT EDIT.

package sideeffectexecution

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sideeffectexecution type in the database.
	Label = "side_effect_execution"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRunID holds the string denoting the run_id field in the database.
	FieldRunID = "run_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAttempt holds the string denoting the attempt field in the database.
	FieldAttempt = "attempt"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeSideEffect holds the string denoting the side_effect edge name in mutations.
	EdgeSideEffect = "side_effect"
	// EdgeActivityExecution holds the string denoting the activity_execution edge name in mutations.
	EdgeActivityExecution = "activity_execution"
	// Table holds the table name of the sideeffectexecution in the database.
	Table = "side_effect_executions"
	// SideEffectTable is the table that holds the side_effect relation/edge.
	SideEffectTable = "side_effect_executions"
	// SideEffectInverseTable is the table name for the SideEffect entity.
	// It exists in this package in order to avoid circular dependency with the "sideeffect" package.
	SideEffectInverseTable = "side_effects"
	// SideEffectColumn is the table column denoting the side_effect relation/edge.
	SideEffectColumn = "side_effect_executions"
	// ActivityExecutionTable is the table that holds the activity_execution relation/edge.
	ActivityExecutionTable = "side_effect_executions"
	// ActivityExecutionInverseTable is the table name for the ActivityExecution entity.
	// It exists in this package in order to avoid circular dependency with the "activityexecution" package.
	ActivityExecutionInverseTable = "activity_executions"
	// ActivityExecutionColumn is the table column denoting the activity_execution relation/edge.
	ActivityExecutionColumn = "activity_execution_side_effect_executions"
)

// Columns holds all SQL columns for sideeffectexecution fields.
var Columns = []string{
	FieldID,
	FieldRunID,
	FieldStatus,
	FieldAttempt,
	FieldOutput,
	FieldStartedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "side_effect_executions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"activity_execution_side_effect_executions",
	"side_effect_executions",
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

var (
	// DefaultAttempt holds the default value on creation for the "attempt" field.
	DefaultAttempt int
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
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
		return fmt.Errorf("sideeffectexecution: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the SideEffectExecution queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRunID orders the results by the run_id field.
func ByRunID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRunID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAttempt orders the results by the attempt field.
func ByAttempt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttempt, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySideEffectField orders the results by side_effect field.
func BySideEffectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSideEffectStep(), sql.OrderByField(field, opts...))
	}
}

// ByActivityExecutionField orders the results by activity_execution field.
func ByActivityExecutionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActivityExecutionStep(), sql.OrderByField(field, opts...))
	}
}
func newSideEffectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SideEffectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SideEffectTable, SideEffectColumn),
	)
}
func newActivityExecutionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActivityExecutionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActivityExecutionTable, ActivityExecutionColumn),
	)
}
