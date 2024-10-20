// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/activity"
	"github.com/davidroman0O/go-tempolite/ent/activityexecution"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
)

// ActivityExecutionUpdate is the builder for updating ActivityExecution entities.
type ActivityExecutionUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityExecutionMutation
}

// Where appends a list predicates to the ActivityExecutionUpdate builder.
func (aeu *ActivityExecutionUpdate) Where(ps ...predicate.ActivityExecution) *ActivityExecutionUpdate {
	aeu.mutation.Where(ps...)
	return aeu
}

// SetRunID sets the "run_id" field.
func (aeu *ActivityExecutionUpdate) SetRunID(s string) *ActivityExecutionUpdate {
	aeu.mutation.SetRunID(s)
	return aeu
}

// SetNillableRunID sets the "run_id" field if the given value is not nil.
func (aeu *ActivityExecutionUpdate) SetNillableRunID(s *string) *ActivityExecutionUpdate {
	if s != nil {
		aeu.SetRunID(*s)
	}
	return aeu
}

// SetStatus sets the "status" field.
func (aeu *ActivityExecutionUpdate) SetStatus(a activityexecution.Status) *ActivityExecutionUpdate {
	aeu.mutation.SetStatus(a)
	return aeu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aeu *ActivityExecutionUpdate) SetNillableStatus(a *activityexecution.Status) *ActivityExecutionUpdate {
	if a != nil {
		aeu.SetStatus(*a)
	}
	return aeu
}

// SetAttempt sets the "attempt" field.
func (aeu *ActivityExecutionUpdate) SetAttempt(i int) *ActivityExecutionUpdate {
	aeu.mutation.ResetAttempt()
	aeu.mutation.SetAttempt(i)
	return aeu
}

// SetNillableAttempt sets the "attempt" field if the given value is not nil.
func (aeu *ActivityExecutionUpdate) SetNillableAttempt(i *int) *ActivityExecutionUpdate {
	if i != nil {
		aeu.SetAttempt(*i)
	}
	return aeu
}

// AddAttempt adds i to the "attempt" field.
func (aeu *ActivityExecutionUpdate) AddAttempt(i int) *ActivityExecutionUpdate {
	aeu.mutation.AddAttempt(i)
	return aeu
}

// SetOutput sets the "output" field.
func (aeu *ActivityExecutionUpdate) SetOutput(i []interface{}) *ActivityExecutionUpdate {
	aeu.mutation.SetOutput(i)
	return aeu
}

// AppendOutput appends i to the "output" field.
func (aeu *ActivityExecutionUpdate) AppendOutput(i []interface{}) *ActivityExecutionUpdate {
	aeu.mutation.AppendOutput(i)
	return aeu
}

// ClearOutput clears the value of the "output" field.
func (aeu *ActivityExecutionUpdate) ClearOutput() *ActivityExecutionUpdate {
	aeu.mutation.ClearOutput()
	return aeu
}

// SetError sets the "error" field.
func (aeu *ActivityExecutionUpdate) SetError(s string) *ActivityExecutionUpdate {
	aeu.mutation.SetError(s)
	return aeu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (aeu *ActivityExecutionUpdate) SetNillableError(s *string) *ActivityExecutionUpdate {
	if s != nil {
		aeu.SetError(*s)
	}
	return aeu
}

// ClearError clears the value of the "error" field.
func (aeu *ActivityExecutionUpdate) ClearError() *ActivityExecutionUpdate {
	aeu.mutation.ClearError()
	return aeu
}

// SetStartedAt sets the "started_at" field.
func (aeu *ActivityExecutionUpdate) SetStartedAt(t time.Time) *ActivityExecutionUpdate {
	aeu.mutation.SetStartedAt(t)
	return aeu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (aeu *ActivityExecutionUpdate) SetNillableStartedAt(t *time.Time) *ActivityExecutionUpdate {
	if t != nil {
		aeu.SetStartedAt(*t)
	}
	return aeu
}

// SetUpdatedAt sets the "updated_at" field.
func (aeu *ActivityExecutionUpdate) SetUpdatedAt(t time.Time) *ActivityExecutionUpdate {
	aeu.mutation.SetUpdatedAt(t)
	return aeu
}

// SetActivityID sets the "activity" edge to the Activity entity by ID.
func (aeu *ActivityExecutionUpdate) SetActivityID(id string) *ActivityExecutionUpdate {
	aeu.mutation.SetActivityID(id)
	return aeu
}

// SetActivity sets the "activity" edge to the Activity entity.
func (aeu *ActivityExecutionUpdate) SetActivity(a *Activity) *ActivityExecutionUpdate {
	return aeu.SetActivityID(a.ID)
}

// Mutation returns the ActivityExecutionMutation object of the builder.
func (aeu *ActivityExecutionUpdate) Mutation() *ActivityExecutionMutation {
	return aeu.mutation
}

// ClearActivity clears the "activity" edge to the Activity entity.
func (aeu *ActivityExecutionUpdate) ClearActivity() *ActivityExecutionUpdate {
	aeu.mutation.ClearActivity()
	return aeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aeu *ActivityExecutionUpdate) Save(ctx context.Context) (int, error) {
	aeu.defaults()
	return withHooks(ctx, aeu.sqlSave, aeu.mutation, aeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeu *ActivityExecutionUpdate) SaveX(ctx context.Context) int {
	affected, err := aeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aeu *ActivityExecutionUpdate) Exec(ctx context.Context) error {
	_, err := aeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeu *ActivityExecutionUpdate) ExecX(ctx context.Context) {
	if err := aeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aeu *ActivityExecutionUpdate) defaults() {
	if _, ok := aeu.mutation.UpdatedAt(); !ok {
		v := activityexecution.UpdateDefaultUpdatedAt()
		aeu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aeu *ActivityExecutionUpdate) check() error {
	if v, ok := aeu.mutation.Status(); ok {
		if err := activityexecution.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ActivityExecution.status": %w`, err)}
		}
	}
	if aeu.mutation.ActivityCleared() && len(aeu.mutation.ActivityIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ActivityExecution.activity"`)
	}
	return nil
}

func (aeu *ActivityExecutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aeu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(activityexecution.Table, activityexecution.Columns, sqlgraph.NewFieldSpec(activityexecution.FieldID, field.TypeString))
	if ps := aeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeu.mutation.RunID(); ok {
		_spec.SetField(activityexecution.FieldRunID, field.TypeString, value)
	}
	if value, ok := aeu.mutation.Status(); ok {
		_spec.SetField(activityexecution.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := aeu.mutation.Attempt(); ok {
		_spec.SetField(activityexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := aeu.mutation.AddedAttempt(); ok {
		_spec.AddField(activityexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := aeu.mutation.Output(); ok {
		_spec.SetField(activityexecution.FieldOutput, field.TypeJSON, value)
	}
	if value, ok := aeu.mutation.AppendedOutput(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, activityexecution.FieldOutput, value)
		})
	}
	if aeu.mutation.OutputCleared() {
		_spec.ClearField(activityexecution.FieldOutput, field.TypeJSON)
	}
	if value, ok := aeu.mutation.Error(); ok {
		_spec.SetField(activityexecution.FieldError, field.TypeString, value)
	}
	if aeu.mutation.ErrorCleared() {
		_spec.ClearField(activityexecution.FieldError, field.TypeString)
	}
	if value, ok := aeu.mutation.StartedAt(); ok {
		_spec.SetField(activityexecution.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := aeu.mutation.UpdatedAt(); ok {
		_spec.SetField(activityexecution.FieldUpdatedAt, field.TypeTime, value)
	}
	if aeu.mutation.ActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activityexecution.ActivityTable,
			Columns: []string{activityexecution.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aeu.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activityexecution.ActivityTable,
			Columns: []string{activityexecution.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activityexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aeu.mutation.done = true
	return n, nil
}

// ActivityExecutionUpdateOne is the builder for updating a single ActivityExecution entity.
type ActivityExecutionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityExecutionMutation
}

// SetRunID sets the "run_id" field.
func (aeuo *ActivityExecutionUpdateOne) SetRunID(s string) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetRunID(s)
	return aeuo
}

// SetNillableRunID sets the "run_id" field if the given value is not nil.
func (aeuo *ActivityExecutionUpdateOne) SetNillableRunID(s *string) *ActivityExecutionUpdateOne {
	if s != nil {
		aeuo.SetRunID(*s)
	}
	return aeuo
}

// SetStatus sets the "status" field.
func (aeuo *ActivityExecutionUpdateOne) SetStatus(a activityexecution.Status) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetStatus(a)
	return aeuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aeuo *ActivityExecutionUpdateOne) SetNillableStatus(a *activityexecution.Status) *ActivityExecutionUpdateOne {
	if a != nil {
		aeuo.SetStatus(*a)
	}
	return aeuo
}

// SetAttempt sets the "attempt" field.
func (aeuo *ActivityExecutionUpdateOne) SetAttempt(i int) *ActivityExecutionUpdateOne {
	aeuo.mutation.ResetAttempt()
	aeuo.mutation.SetAttempt(i)
	return aeuo
}

// SetNillableAttempt sets the "attempt" field if the given value is not nil.
func (aeuo *ActivityExecutionUpdateOne) SetNillableAttempt(i *int) *ActivityExecutionUpdateOne {
	if i != nil {
		aeuo.SetAttempt(*i)
	}
	return aeuo
}

// AddAttempt adds i to the "attempt" field.
func (aeuo *ActivityExecutionUpdateOne) AddAttempt(i int) *ActivityExecutionUpdateOne {
	aeuo.mutation.AddAttempt(i)
	return aeuo
}

// SetOutput sets the "output" field.
func (aeuo *ActivityExecutionUpdateOne) SetOutput(i []interface{}) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetOutput(i)
	return aeuo
}

// AppendOutput appends i to the "output" field.
func (aeuo *ActivityExecutionUpdateOne) AppendOutput(i []interface{}) *ActivityExecutionUpdateOne {
	aeuo.mutation.AppendOutput(i)
	return aeuo
}

// ClearOutput clears the value of the "output" field.
func (aeuo *ActivityExecutionUpdateOne) ClearOutput() *ActivityExecutionUpdateOne {
	aeuo.mutation.ClearOutput()
	return aeuo
}

// SetError sets the "error" field.
func (aeuo *ActivityExecutionUpdateOne) SetError(s string) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetError(s)
	return aeuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (aeuo *ActivityExecutionUpdateOne) SetNillableError(s *string) *ActivityExecutionUpdateOne {
	if s != nil {
		aeuo.SetError(*s)
	}
	return aeuo
}

// ClearError clears the value of the "error" field.
func (aeuo *ActivityExecutionUpdateOne) ClearError() *ActivityExecutionUpdateOne {
	aeuo.mutation.ClearError()
	return aeuo
}

// SetStartedAt sets the "started_at" field.
func (aeuo *ActivityExecutionUpdateOne) SetStartedAt(t time.Time) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetStartedAt(t)
	return aeuo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (aeuo *ActivityExecutionUpdateOne) SetNillableStartedAt(t *time.Time) *ActivityExecutionUpdateOne {
	if t != nil {
		aeuo.SetStartedAt(*t)
	}
	return aeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aeuo *ActivityExecutionUpdateOne) SetUpdatedAt(t time.Time) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetUpdatedAt(t)
	return aeuo
}

// SetActivityID sets the "activity" edge to the Activity entity by ID.
func (aeuo *ActivityExecutionUpdateOne) SetActivityID(id string) *ActivityExecutionUpdateOne {
	aeuo.mutation.SetActivityID(id)
	return aeuo
}

// SetActivity sets the "activity" edge to the Activity entity.
func (aeuo *ActivityExecutionUpdateOne) SetActivity(a *Activity) *ActivityExecutionUpdateOne {
	return aeuo.SetActivityID(a.ID)
}

// Mutation returns the ActivityExecutionMutation object of the builder.
func (aeuo *ActivityExecutionUpdateOne) Mutation() *ActivityExecutionMutation {
	return aeuo.mutation
}

// ClearActivity clears the "activity" edge to the Activity entity.
func (aeuo *ActivityExecutionUpdateOne) ClearActivity() *ActivityExecutionUpdateOne {
	aeuo.mutation.ClearActivity()
	return aeuo
}

// Where appends a list predicates to the ActivityExecutionUpdate builder.
func (aeuo *ActivityExecutionUpdateOne) Where(ps ...predicate.ActivityExecution) *ActivityExecutionUpdateOne {
	aeuo.mutation.Where(ps...)
	return aeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aeuo *ActivityExecutionUpdateOne) Select(field string, fields ...string) *ActivityExecutionUpdateOne {
	aeuo.fields = append([]string{field}, fields...)
	return aeuo
}

// Save executes the query and returns the updated ActivityExecution entity.
func (aeuo *ActivityExecutionUpdateOne) Save(ctx context.Context) (*ActivityExecution, error) {
	aeuo.defaults()
	return withHooks(ctx, aeuo.sqlSave, aeuo.mutation, aeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeuo *ActivityExecutionUpdateOne) SaveX(ctx context.Context) *ActivityExecution {
	node, err := aeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aeuo *ActivityExecutionUpdateOne) Exec(ctx context.Context) error {
	_, err := aeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeuo *ActivityExecutionUpdateOne) ExecX(ctx context.Context) {
	if err := aeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aeuo *ActivityExecutionUpdateOne) defaults() {
	if _, ok := aeuo.mutation.UpdatedAt(); !ok {
		v := activityexecution.UpdateDefaultUpdatedAt()
		aeuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aeuo *ActivityExecutionUpdateOne) check() error {
	if v, ok := aeuo.mutation.Status(); ok {
		if err := activityexecution.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ActivityExecution.status": %w`, err)}
		}
	}
	if aeuo.mutation.ActivityCleared() && len(aeuo.mutation.ActivityIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ActivityExecution.activity"`)
	}
	return nil
}

func (aeuo *ActivityExecutionUpdateOne) sqlSave(ctx context.Context) (_node *ActivityExecution, err error) {
	if err := aeuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(activityexecution.Table, activityexecution.Columns, sqlgraph.NewFieldSpec(activityexecution.FieldID, field.TypeString))
	id, ok := aeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ActivityExecution.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activityexecution.FieldID)
		for _, f := range fields {
			if !activityexecution.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activityexecution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeuo.mutation.RunID(); ok {
		_spec.SetField(activityexecution.FieldRunID, field.TypeString, value)
	}
	if value, ok := aeuo.mutation.Status(); ok {
		_spec.SetField(activityexecution.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := aeuo.mutation.Attempt(); ok {
		_spec.SetField(activityexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := aeuo.mutation.AddedAttempt(); ok {
		_spec.AddField(activityexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := aeuo.mutation.Output(); ok {
		_spec.SetField(activityexecution.FieldOutput, field.TypeJSON, value)
	}
	if value, ok := aeuo.mutation.AppendedOutput(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, activityexecution.FieldOutput, value)
		})
	}
	if aeuo.mutation.OutputCleared() {
		_spec.ClearField(activityexecution.FieldOutput, field.TypeJSON)
	}
	if value, ok := aeuo.mutation.Error(); ok {
		_spec.SetField(activityexecution.FieldError, field.TypeString, value)
	}
	if aeuo.mutation.ErrorCleared() {
		_spec.ClearField(activityexecution.FieldError, field.TypeString)
	}
	if value, ok := aeuo.mutation.StartedAt(); ok {
		_spec.SetField(activityexecution.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := aeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(activityexecution.FieldUpdatedAt, field.TypeTime, value)
	}
	if aeuo.mutation.ActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activityexecution.ActivityTable,
			Columns: []string{activityexecution.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aeuo.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activityexecution.ActivityTable,
			Columns: []string{activityexecution.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ActivityExecution{config: aeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activityexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aeuo.mutation.done = true
	return _node, nil
}
