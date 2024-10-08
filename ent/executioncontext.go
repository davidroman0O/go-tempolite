// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/davidroman0O/go-tempolite/ent/executioncontext"
)

// ExecutionContext is the model entity for the ExecutionContext schema.
type ExecutionContext struct {
	config
	// ID of the ent.
	ID           string `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExecutionContext) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case executioncontext.FieldID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExecutionContext fields.
func (ec *ExecutionContext) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case executioncontext.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ec.ID = value.String
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExecutionContext.
// This includes values selected through modifiers, order, etc.
func (ec *ExecutionContext) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// Update returns a builder for updating this ExecutionContext.
// Note that you need to call ExecutionContext.Unwrap() before calling this method if this ExecutionContext
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *ExecutionContext) Update() *ExecutionContextUpdateOne {
	return NewExecutionContextClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the ExecutionContext entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *ExecutionContext) Unwrap() *ExecutionContext {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExecutionContext is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *ExecutionContext) String() string {
	var builder strings.Builder
	builder.WriteString("ExecutionContext(")
	builder.WriteString(fmt.Sprintf("id=%v", ec.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ExecutionContexts is a parsable slice of ExecutionContext.
type ExecutionContexts []*ExecutionContext
