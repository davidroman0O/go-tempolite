// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/davidroman0O/go-tempolite/ent/activity"
	"github.com/davidroman0O/go-tempolite/ent/run"
	"github.com/davidroman0O/go-tempolite/ent/workflow"
)

// Run is the model entity for the Run schema.
type Run struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// RunID holds the value of the "run_id" field.
	RunID string `json:"run_id,omitempty"`
	// Type holds the value of the "type" field.
	Type run.Type `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RunQuery when eager-loading is set.
	Edges        RunEdges `json:"edges"`
	run_workflow *string
	run_activity *string
	selectValues sql.SelectValues
}

// RunEdges holds the relations/edges for other nodes in the graph.
type RunEdges struct {
	// A run can be connected to a workflow or an activity, not both.
	Workflow *Workflow `json:"workflow,omitempty"`
	// A run can be connected to a workflow or an activity, not both.
	Activity *Activity `json:"activity,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RunEdges) WorkflowOrErr() (*Workflow, error) {
	if e.Workflow != nil {
		return e.Workflow, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: workflow.Label}
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// ActivityOrErr returns the Activity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RunEdges) ActivityOrErr() (*Activity, error) {
	if e.Activity != nil {
		return e.Activity, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: activity.Label}
	}
	return nil, &NotLoadedError{edge: "activity"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Run) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case run.FieldID, run.FieldRunID, run.FieldType:
			values[i] = new(sql.NullString)
		case run.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case run.ForeignKeys[0]: // run_workflow
			values[i] = new(sql.NullString)
		case run.ForeignKeys[1]: // run_activity
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Run fields.
func (r *Run) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case run.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case run.FieldRunID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field run_id", values[i])
			} else if value.Valid {
				r.RunID = value.String
			}
		case run.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = run.Type(value.String)
			}
		case run.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case run.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field run_workflow", values[i])
			} else if value.Valid {
				r.run_workflow = new(string)
				*r.run_workflow = value.String
			}
		case run.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field run_activity", values[i])
			} else if value.Valid {
				r.run_activity = new(string)
				*r.run_activity = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Run.
// This includes values selected through modifiers, order, etc.
func (r *Run) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryWorkflow queries the "workflow" edge of the Run entity.
func (r *Run) QueryWorkflow() *WorkflowQuery {
	return NewRunClient(r.config).QueryWorkflow(r)
}

// QueryActivity queries the "activity" edge of the Run entity.
func (r *Run) QueryActivity() *ActivityQuery {
	return NewRunClient(r.config).QueryActivity(r)
}

// Update returns a builder for updating this Run.
// Note that you need to call Run.Unwrap() before calling this method if this Run
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Run) Update() *RunUpdateOne {
	return NewRunClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Run entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Run) Unwrap() *Run {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Run is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Run) String() string {
	var builder strings.Builder
	builder.WriteString("Run(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("run_id=")
	builder.WriteString(r.RunID)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", r.Type))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Runs is a parsable slice of Run.
type Runs []*Run