// Code generated by ent, DO NOT EDIT.

package activityexecution

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the activityexecution type in the database.
	Label = "activity_execution"
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
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeActivity holds the string denoting the activity edge name in mutations.
	EdgeActivity = "activity"
	// EdgeWorkflowExecution holds the string denoting the workflow_execution edge name in mutations.
	EdgeWorkflowExecution = "workflow_execution"
	// EdgeSideEffectExecutions holds the string denoting the side_effect_executions edge name in mutations.
	EdgeSideEffectExecutions = "side_effect_executions"
	// Table holds the table name of the activityexecution in the database.
	Table = "activity_executions"
	// ActivityTable is the table that holds the activity relation/edge.
	ActivityTable = "activity_executions"
	// ActivityInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivityInverseTable = "activities"
	// ActivityColumn is the table column denoting the activity relation/edge.
	ActivityColumn = "activity_executions"
	// WorkflowExecutionTable is the table that holds the workflow_execution relation/edge.
	WorkflowExecutionTable = "activity_executions"
	// WorkflowExecutionInverseTable is the table name for the WorkflowExecution entity.
	// It exists in this package in order to avoid circular dependency with the "workflowexecution" package.
	WorkflowExecutionInverseTable = "workflow_executions"
	// WorkflowExecutionColumn is the table column denoting the workflow_execution relation/edge.
	WorkflowExecutionColumn = "workflow_execution_activity_executions"
	// SideEffectExecutionsTable is the table that holds the side_effect_executions relation/edge.
	SideEffectExecutionsTable = "side_effect_executions"
	// SideEffectExecutionsInverseTable is the table name for the SideEffectExecution entity.
	// It exists in this package in order to avoid circular dependency with the "sideeffectexecution" package.
	SideEffectExecutionsInverseTable = "side_effect_executions"
	// SideEffectExecutionsColumn is the table column denoting the side_effect_executions relation/edge.
	SideEffectExecutionsColumn = "activity_execution_side_effect_executions"
)

// Columns holds all SQL columns for activityexecution fields.
var Columns = []string{
	FieldID,
	FieldRunID,
	FieldStatus,
	FieldAttempt,
	FieldOutput,
	FieldError,
	FieldStartedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "activity_executions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"activity_executions",
	"workflow_execution_activity_executions",
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
	StatusRetried   Status = "Retried"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusRunning, StatusCompleted, StatusFailed, StatusRetried:
		return nil
	default:
		return fmt.Errorf("activityexecution: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the ActivityExecution queries.
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

// ByError orders the results by the error field.
func ByError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldError, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByActivityField orders the results by activity field.
func ByActivityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActivityStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkflowExecutionField orders the results by workflow_execution field.
func ByWorkflowExecutionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowExecutionStep(), sql.OrderByField(field, opts...))
	}
}

// BySideEffectExecutionsCount orders the results by side_effect_executions count.
func BySideEffectExecutionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSideEffectExecutionsStep(), opts...)
	}
}

// BySideEffectExecutions orders the results by side_effect_executions terms.
func BySideEffectExecutions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSideEffectExecutionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newActivityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActivityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
	)
}
func newWorkflowExecutionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowExecutionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WorkflowExecutionTable, WorkflowExecutionColumn),
	)
}
func newSideEffectExecutionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SideEffectExecutionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SideEffectExecutionsTable, SideEffectExecutionsColumn),
	)
}
