// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
	"github.com/davidroman0O/go-tempolite/ent/taskcontext"
)

// TaskContextUpdate is the builder for updating TaskContext entities.
type TaskContextUpdate struct {
	config
	hooks    []Hook
	mutation *TaskContextMutation
}

// Where appends a list predicates to the TaskContextUpdate builder.
func (tcu *TaskContextUpdate) Where(ps ...predicate.TaskContext) *TaskContextUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetRetryCount sets the "RetryCount" field.
func (tcu *TaskContextUpdate) SetRetryCount(i int) *TaskContextUpdate {
	tcu.mutation.ResetRetryCount()
	tcu.mutation.SetRetryCount(i)
	return tcu
}

// SetNillableRetryCount sets the "RetryCount" field if the given value is not nil.
func (tcu *TaskContextUpdate) SetNillableRetryCount(i *int) *TaskContextUpdate {
	if i != nil {
		tcu.SetRetryCount(*i)
	}
	return tcu
}

// AddRetryCount adds i to the "RetryCount" field.
func (tcu *TaskContextUpdate) AddRetryCount(i int) *TaskContextUpdate {
	tcu.mutation.AddRetryCount(i)
	return tcu
}

// SetMaxRetry sets the "MaxRetry" field.
func (tcu *TaskContextUpdate) SetMaxRetry(i int) *TaskContextUpdate {
	tcu.mutation.ResetMaxRetry()
	tcu.mutation.SetMaxRetry(i)
	return tcu
}

// SetNillableMaxRetry sets the "MaxRetry" field if the given value is not nil.
func (tcu *TaskContextUpdate) SetNillableMaxRetry(i *int) *TaskContextUpdate {
	if i != nil {
		tcu.SetMaxRetry(*i)
	}
	return tcu
}

// AddMaxRetry adds i to the "MaxRetry" field.
func (tcu *TaskContextUpdate) AddMaxRetry(i int) *TaskContextUpdate {
	tcu.mutation.AddMaxRetry(i)
	return tcu
}

// Mutation returns the TaskContextMutation object of the builder.
func (tcu *TaskContextUpdate) Mutation() *TaskContextMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TaskContextUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tcu.sqlSave, tcu.mutation, tcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TaskContextUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TaskContextUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TaskContextUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tcu *TaskContextUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(taskcontext.Table, taskcontext.Columns, sqlgraph.NewFieldSpec(taskcontext.FieldID, field.TypeString))
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.RetryCount(); ok {
		_spec.SetField(taskcontext.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := tcu.mutation.AddedRetryCount(); ok {
		_spec.AddField(taskcontext.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := tcu.mutation.MaxRetry(); ok {
		_spec.SetField(taskcontext.FieldMaxRetry, field.TypeInt, value)
	}
	if value, ok := tcu.mutation.AddedMaxRetry(); ok {
		_spec.AddField(taskcontext.FieldMaxRetry, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskcontext.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcu.mutation.done = true
	return n, nil
}

// TaskContextUpdateOne is the builder for updating a single TaskContext entity.
type TaskContextUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskContextMutation
}

// SetRetryCount sets the "RetryCount" field.
func (tcuo *TaskContextUpdateOne) SetRetryCount(i int) *TaskContextUpdateOne {
	tcuo.mutation.ResetRetryCount()
	tcuo.mutation.SetRetryCount(i)
	return tcuo
}

// SetNillableRetryCount sets the "RetryCount" field if the given value is not nil.
func (tcuo *TaskContextUpdateOne) SetNillableRetryCount(i *int) *TaskContextUpdateOne {
	if i != nil {
		tcuo.SetRetryCount(*i)
	}
	return tcuo
}

// AddRetryCount adds i to the "RetryCount" field.
func (tcuo *TaskContextUpdateOne) AddRetryCount(i int) *TaskContextUpdateOne {
	tcuo.mutation.AddRetryCount(i)
	return tcuo
}

// SetMaxRetry sets the "MaxRetry" field.
func (tcuo *TaskContextUpdateOne) SetMaxRetry(i int) *TaskContextUpdateOne {
	tcuo.mutation.ResetMaxRetry()
	tcuo.mutation.SetMaxRetry(i)
	return tcuo
}

// SetNillableMaxRetry sets the "MaxRetry" field if the given value is not nil.
func (tcuo *TaskContextUpdateOne) SetNillableMaxRetry(i *int) *TaskContextUpdateOne {
	if i != nil {
		tcuo.SetMaxRetry(*i)
	}
	return tcuo
}

// AddMaxRetry adds i to the "MaxRetry" field.
func (tcuo *TaskContextUpdateOne) AddMaxRetry(i int) *TaskContextUpdateOne {
	tcuo.mutation.AddMaxRetry(i)
	return tcuo
}

// Mutation returns the TaskContextMutation object of the builder.
func (tcuo *TaskContextUpdateOne) Mutation() *TaskContextMutation {
	return tcuo.mutation
}

// Where appends a list predicates to the TaskContextUpdate builder.
func (tcuo *TaskContextUpdateOne) Where(ps ...predicate.TaskContext) *TaskContextUpdateOne {
	tcuo.mutation.Where(ps...)
	return tcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TaskContextUpdateOne) Select(field string, fields ...string) *TaskContextUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TaskContext entity.
func (tcuo *TaskContextUpdateOne) Save(ctx context.Context) (*TaskContext, error) {
	return withHooks(ctx, tcuo.sqlSave, tcuo.mutation, tcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TaskContextUpdateOne) SaveX(ctx context.Context) *TaskContext {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TaskContextUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TaskContextUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tcuo *TaskContextUpdateOne) sqlSave(ctx context.Context) (_node *TaskContext, err error) {
	_spec := sqlgraph.NewUpdateSpec(taskcontext.Table, taskcontext.Columns, sqlgraph.NewFieldSpec(taskcontext.FieldID, field.TypeString))
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskContext.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskcontext.FieldID)
		for _, f := range fields {
			if !taskcontext.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskcontext.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.RetryCount(); ok {
		_spec.SetField(taskcontext.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := tcuo.mutation.AddedRetryCount(); ok {
		_spec.AddField(taskcontext.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := tcuo.mutation.MaxRetry(); ok {
		_spec.SetField(taskcontext.FieldMaxRetry, field.TypeInt, value)
	}
	if value, ok := tcuo.mutation.AddedMaxRetry(); ok {
		_spec.AddField(taskcontext.FieldMaxRetry, field.TypeInt, value)
	}
	_node = &TaskContext{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskcontext.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcuo.mutation.done = true
	return _node, nil
}
