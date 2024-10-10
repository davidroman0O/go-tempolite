// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/executioncontext"
	"github.com/davidroman0O/go-tempolite/ent/handlerexecution"
	"github.com/davidroman0O/go-tempolite/ent/handlertask"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
	"github.com/davidroman0O/go-tempolite/ent/sagastepexecution"
)

// HandlerExecutionUpdate is the builder for updating HandlerExecution entities.
type HandlerExecutionUpdate struct {
	config
	hooks    []Hook
	mutation *HandlerExecutionMutation
}

// Where appends a list predicates to the HandlerExecutionUpdate builder.
func (heu *HandlerExecutionUpdate) Where(ps ...predicate.HandlerExecution) *HandlerExecutionUpdate {
	heu.mutation.Where(ps...)
	return heu
}

// SetRunID sets the "run_id" field.
func (heu *HandlerExecutionUpdate) SetRunID(s string) *HandlerExecutionUpdate {
	heu.mutation.SetRunID(s)
	return heu
}

// SetNillableRunID sets the "run_id" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableRunID(s *string) *HandlerExecutionUpdate {
	if s != nil {
		heu.SetRunID(*s)
	}
	return heu
}

// SetHandlerName sets the "handler_name" field.
func (heu *HandlerExecutionUpdate) SetHandlerName(s string) *HandlerExecutionUpdate {
	heu.mutation.SetHandlerName(s)
	return heu
}

// SetNillableHandlerName sets the "handler_name" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableHandlerName(s *string) *HandlerExecutionUpdate {
	if s != nil {
		heu.SetHandlerName(*s)
	}
	return heu
}

// SetStatus sets the "status" field.
func (heu *HandlerExecutionUpdate) SetStatus(h handlerexecution.Status) *HandlerExecutionUpdate {
	heu.mutation.SetStatus(h)
	return heu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableStatus(h *handlerexecution.Status) *HandlerExecutionUpdate {
	if h != nil {
		heu.SetStatus(*h)
	}
	return heu
}

// SetStartTime sets the "start_time" field.
func (heu *HandlerExecutionUpdate) SetStartTime(t time.Time) *HandlerExecutionUpdate {
	heu.mutation.SetStartTime(t)
	return heu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableStartTime(t *time.Time) *HandlerExecutionUpdate {
	if t != nil {
		heu.SetStartTime(*t)
	}
	return heu
}

// SetEndTime sets the "end_time" field.
func (heu *HandlerExecutionUpdate) SetEndTime(t time.Time) *HandlerExecutionUpdate {
	heu.mutation.SetEndTime(t)
	return heu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableEndTime(t *time.Time) *HandlerExecutionUpdate {
	if t != nil {
		heu.SetEndTime(*t)
	}
	return heu
}

// ClearEndTime clears the value of the "end_time" field.
func (heu *HandlerExecutionUpdate) ClearEndTime() *HandlerExecutionUpdate {
	heu.mutation.ClearEndTime()
	return heu
}

// SetRetryCount sets the "retry_count" field.
func (heu *HandlerExecutionUpdate) SetRetryCount(i int) *HandlerExecutionUpdate {
	heu.mutation.ResetRetryCount()
	heu.mutation.SetRetryCount(i)
	return heu
}

// SetNillableRetryCount sets the "retry_count" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableRetryCount(i *int) *HandlerExecutionUpdate {
	if i != nil {
		heu.SetRetryCount(*i)
	}
	return heu
}

// AddRetryCount adds i to the "retry_count" field.
func (heu *HandlerExecutionUpdate) AddRetryCount(i int) *HandlerExecutionUpdate {
	heu.mutation.AddRetryCount(i)
	return heu
}

// SetMaxRetries sets the "max_retries" field.
func (heu *HandlerExecutionUpdate) SetMaxRetries(i int) *HandlerExecutionUpdate {
	heu.mutation.ResetMaxRetries()
	heu.mutation.SetMaxRetries(i)
	return heu
}

// SetNillableMaxRetries sets the "max_retries" field if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableMaxRetries(i *int) *HandlerExecutionUpdate {
	if i != nil {
		heu.SetMaxRetries(*i)
	}
	return heu
}

// AddMaxRetries adds i to the "max_retries" field.
func (heu *HandlerExecutionUpdate) AddMaxRetries(i int) *HandlerExecutionUpdate {
	heu.mutation.AddMaxRetries(i)
	return heu
}

// SetExecutionContextID sets the "execution_context" edge to the ExecutionContext entity by ID.
func (heu *HandlerExecutionUpdate) SetExecutionContextID(id string) *HandlerExecutionUpdate {
	heu.mutation.SetExecutionContextID(id)
	return heu
}

// SetExecutionContext sets the "execution_context" edge to the ExecutionContext entity.
func (heu *HandlerExecutionUpdate) SetExecutionContext(e *ExecutionContext) *HandlerExecutionUpdate {
	return heu.SetExecutionContextID(e.ID)
}

// SetParentID sets the "parent" edge to the HandlerExecution entity by ID.
func (heu *HandlerExecutionUpdate) SetParentID(id string) *HandlerExecutionUpdate {
	heu.mutation.SetParentID(id)
	return heu
}

// SetNillableParentID sets the "parent" edge to the HandlerExecution entity by ID if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableParentID(id *string) *HandlerExecutionUpdate {
	if id != nil {
		heu = heu.SetParentID(*id)
	}
	return heu
}

// SetParent sets the "parent" edge to the HandlerExecution entity.
func (heu *HandlerExecutionUpdate) SetParent(h *HandlerExecution) *HandlerExecutionUpdate {
	return heu.SetParentID(h.ID)
}

// AddChildIDs adds the "children" edge to the HandlerExecution entity by IDs.
func (heu *HandlerExecutionUpdate) AddChildIDs(ids ...string) *HandlerExecutionUpdate {
	heu.mutation.AddChildIDs(ids...)
	return heu
}

// AddChildren adds the "children" edges to the HandlerExecution entity.
func (heu *HandlerExecutionUpdate) AddChildren(h ...*HandlerExecution) *HandlerExecutionUpdate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heu.AddChildIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the HandlerTask entity by IDs.
func (heu *HandlerExecutionUpdate) AddTaskIDs(ids ...string) *HandlerExecutionUpdate {
	heu.mutation.AddTaskIDs(ids...)
	return heu
}

// AddTasks adds the "tasks" edges to the HandlerTask entity.
func (heu *HandlerExecutionUpdate) AddTasks(h ...*HandlerTask) *HandlerExecutionUpdate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heu.AddTaskIDs(ids...)
}

// SetSagaStepExecutionID sets the "saga_step_execution" edge to the SagaStepExecution entity by ID.
func (heu *HandlerExecutionUpdate) SetSagaStepExecutionID(id string) *HandlerExecutionUpdate {
	heu.mutation.SetSagaStepExecutionID(id)
	return heu
}

// SetNillableSagaStepExecutionID sets the "saga_step_execution" edge to the SagaStepExecution entity by ID if the given value is not nil.
func (heu *HandlerExecutionUpdate) SetNillableSagaStepExecutionID(id *string) *HandlerExecutionUpdate {
	if id != nil {
		heu = heu.SetSagaStepExecutionID(*id)
	}
	return heu
}

// SetSagaStepExecution sets the "saga_step_execution" edge to the SagaStepExecution entity.
func (heu *HandlerExecutionUpdate) SetSagaStepExecution(s *SagaStepExecution) *HandlerExecutionUpdate {
	return heu.SetSagaStepExecutionID(s.ID)
}

// Mutation returns the HandlerExecutionMutation object of the builder.
func (heu *HandlerExecutionUpdate) Mutation() *HandlerExecutionMutation {
	return heu.mutation
}

// ClearExecutionContext clears the "execution_context" edge to the ExecutionContext entity.
func (heu *HandlerExecutionUpdate) ClearExecutionContext() *HandlerExecutionUpdate {
	heu.mutation.ClearExecutionContext()
	return heu
}

// ClearParent clears the "parent" edge to the HandlerExecution entity.
func (heu *HandlerExecutionUpdate) ClearParent() *HandlerExecutionUpdate {
	heu.mutation.ClearParent()
	return heu
}

// ClearChildren clears all "children" edges to the HandlerExecution entity.
func (heu *HandlerExecutionUpdate) ClearChildren() *HandlerExecutionUpdate {
	heu.mutation.ClearChildren()
	return heu
}

// RemoveChildIDs removes the "children" edge to HandlerExecution entities by IDs.
func (heu *HandlerExecutionUpdate) RemoveChildIDs(ids ...string) *HandlerExecutionUpdate {
	heu.mutation.RemoveChildIDs(ids...)
	return heu
}

// RemoveChildren removes "children" edges to HandlerExecution entities.
func (heu *HandlerExecutionUpdate) RemoveChildren(h ...*HandlerExecution) *HandlerExecutionUpdate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heu.RemoveChildIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the HandlerTask entity.
func (heu *HandlerExecutionUpdate) ClearTasks() *HandlerExecutionUpdate {
	heu.mutation.ClearTasks()
	return heu
}

// RemoveTaskIDs removes the "tasks" edge to HandlerTask entities by IDs.
func (heu *HandlerExecutionUpdate) RemoveTaskIDs(ids ...string) *HandlerExecutionUpdate {
	heu.mutation.RemoveTaskIDs(ids...)
	return heu
}

// RemoveTasks removes "tasks" edges to HandlerTask entities.
func (heu *HandlerExecutionUpdate) RemoveTasks(h ...*HandlerTask) *HandlerExecutionUpdate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heu.RemoveTaskIDs(ids...)
}

// ClearSagaStepExecution clears the "saga_step_execution" edge to the SagaStepExecution entity.
func (heu *HandlerExecutionUpdate) ClearSagaStepExecution() *HandlerExecutionUpdate {
	heu.mutation.ClearSagaStepExecution()
	return heu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (heu *HandlerExecutionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, heu.sqlSave, heu.mutation, heu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (heu *HandlerExecutionUpdate) SaveX(ctx context.Context) int {
	affected, err := heu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (heu *HandlerExecutionUpdate) Exec(ctx context.Context) error {
	_, err := heu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (heu *HandlerExecutionUpdate) ExecX(ctx context.Context) {
	if err := heu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (heu *HandlerExecutionUpdate) check() error {
	if v, ok := heu.mutation.Status(); ok {
		if err := handlerexecution.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HandlerExecution.status": %w`, err)}
		}
	}
	if heu.mutation.ExecutionContextCleared() && len(heu.mutation.ExecutionContextIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "HandlerExecution.execution_context"`)
	}
	return nil
}

func (heu *HandlerExecutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := heu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(handlerexecution.Table, handlerexecution.Columns, sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString))
	if ps := heu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := heu.mutation.RunID(); ok {
		_spec.SetField(handlerexecution.FieldRunID, field.TypeString, value)
	}
	if value, ok := heu.mutation.HandlerName(); ok {
		_spec.SetField(handlerexecution.FieldHandlerName, field.TypeString, value)
	}
	if value, ok := heu.mutation.Status(); ok {
		_spec.SetField(handlerexecution.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := heu.mutation.StartTime(); ok {
		_spec.SetField(handlerexecution.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := heu.mutation.EndTime(); ok {
		_spec.SetField(handlerexecution.FieldEndTime, field.TypeTime, value)
	}
	if heu.mutation.EndTimeCleared() {
		_spec.ClearField(handlerexecution.FieldEndTime, field.TypeTime)
	}
	if value, ok := heu.mutation.RetryCount(); ok {
		_spec.SetField(handlerexecution.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := heu.mutation.AddedRetryCount(); ok {
		_spec.AddField(handlerexecution.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := heu.mutation.MaxRetries(); ok {
		_spec.SetField(handlerexecution.FieldMaxRetries, field.TypeInt, value)
	}
	if value, ok := heu.mutation.AddedMaxRetries(); ok {
		_spec.AddField(handlerexecution.FieldMaxRetries, field.TypeInt, value)
	}
	if heu.mutation.ExecutionContextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ExecutionContextTable,
			Columns: []string{handlerexecution.ExecutionContextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(executioncontext.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.ExecutionContextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ExecutionContextTable,
			Columns: []string{handlerexecution.ExecutionContextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(executioncontext.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ParentTable,
			Columns: []string{handlerexecution.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ParentTable,
			Columns: []string{handlerexecution.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.ChildrenTable,
			Columns: []string{handlerexecution.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !heu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.ChildrenTable,
			Columns: []string{handlerexecution.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.ChildrenTable,
			Columns: []string{handlerexecution.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.TasksTable,
			Columns: []string{handlerexecution.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !heu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.TasksTable,
			Columns: []string{handlerexecution.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.TasksTable,
			Columns: []string{handlerexecution.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.SagaStepExecutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   handlerexecution.SagaStepExecutionTable,
			Columns: []string{handlerexecution.SagaStepExecutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sagastepexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.SagaStepExecutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   handlerexecution.SagaStepExecutionTable,
			Columns: []string{handlerexecution.SagaStepExecutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sagastepexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, heu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{handlerexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	heu.mutation.done = true
	return n, nil
}

// HandlerExecutionUpdateOne is the builder for updating a single HandlerExecution entity.
type HandlerExecutionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HandlerExecutionMutation
}

// SetRunID sets the "run_id" field.
func (heuo *HandlerExecutionUpdateOne) SetRunID(s string) *HandlerExecutionUpdateOne {
	heuo.mutation.SetRunID(s)
	return heuo
}

// SetNillableRunID sets the "run_id" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableRunID(s *string) *HandlerExecutionUpdateOne {
	if s != nil {
		heuo.SetRunID(*s)
	}
	return heuo
}

// SetHandlerName sets the "handler_name" field.
func (heuo *HandlerExecutionUpdateOne) SetHandlerName(s string) *HandlerExecutionUpdateOne {
	heuo.mutation.SetHandlerName(s)
	return heuo
}

// SetNillableHandlerName sets the "handler_name" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableHandlerName(s *string) *HandlerExecutionUpdateOne {
	if s != nil {
		heuo.SetHandlerName(*s)
	}
	return heuo
}

// SetStatus sets the "status" field.
func (heuo *HandlerExecutionUpdateOne) SetStatus(h handlerexecution.Status) *HandlerExecutionUpdateOne {
	heuo.mutation.SetStatus(h)
	return heuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableStatus(h *handlerexecution.Status) *HandlerExecutionUpdateOne {
	if h != nil {
		heuo.SetStatus(*h)
	}
	return heuo
}

// SetStartTime sets the "start_time" field.
func (heuo *HandlerExecutionUpdateOne) SetStartTime(t time.Time) *HandlerExecutionUpdateOne {
	heuo.mutation.SetStartTime(t)
	return heuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableStartTime(t *time.Time) *HandlerExecutionUpdateOne {
	if t != nil {
		heuo.SetStartTime(*t)
	}
	return heuo
}

// SetEndTime sets the "end_time" field.
func (heuo *HandlerExecutionUpdateOne) SetEndTime(t time.Time) *HandlerExecutionUpdateOne {
	heuo.mutation.SetEndTime(t)
	return heuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableEndTime(t *time.Time) *HandlerExecutionUpdateOne {
	if t != nil {
		heuo.SetEndTime(*t)
	}
	return heuo
}

// ClearEndTime clears the value of the "end_time" field.
func (heuo *HandlerExecutionUpdateOne) ClearEndTime() *HandlerExecutionUpdateOne {
	heuo.mutation.ClearEndTime()
	return heuo
}

// SetRetryCount sets the "retry_count" field.
func (heuo *HandlerExecutionUpdateOne) SetRetryCount(i int) *HandlerExecutionUpdateOne {
	heuo.mutation.ResetRetryCount()
	heuo.mutation.SetRetryCount(i)
	return heuo
}

// SetNillableRetryCount sets the "retry_count" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableRetryCount(i *int) *HandlerExecutionUpdateOne {
	if i != nil {
		heuo.SetRetryCount(*i)
	}
	return heuo
}

// AddRetryCount adds i to the "retry_count" field.
func (heuo *HandlerExecutionUpdateOne) AddRetryCount(i int) *HandlerExecutionUpdateOne {
	heuo.mutation.AddRetryCount(i)
	return heuo
}

// SetMaxRetries sets the "max_retries" field.
func (heuo *HandlerExecutionUpdateOne) SetMaxRetries(i int) *HandlerExecutionUpdateOne {
	heuo.mutation.ResetMaxRetries()
	heuo.mutation.SetMaxRetries(i)
	return heuo
}

// SetNillableMaxRetries sets the "max_retries" field if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableMaxRetries(i *int) *HandlerExecutionUpdateOne {
	if i != nil {
		heuo.SetMaxRetries(*i)
	}
	return heuo
}

// AddMaxRetries adds i to the "max_retries" field.
func (heuo *HandlerExecutionUpdateOne) AddMaxRetries(i int) *HandlerExecutionUpdateOne {
	heuo.mutation.AddMaxRetries(i)
	return heuo
}

// SetExecutionContextID sets the "execution_context" edge to the ExecutionContext entity by ID.
func (heuo *HandlerExecutionUpdateOne) SetExecutionContextID(id string) *HandlerExecutionUpdateOne {
	heuo.mutation.SetExecutionContextID(id)
	return heuo
}

// SetExecutionContext sets the "execution_context" edge to the ExecutionContext entity.
func (heuo *HandlerExecutionUpdateOne) SetExecutionContext(e *ExecutionContext) *HandlerExecutionUpdateOne {
	return heuo.SetExecutionContextID(e.ID)
}

// SetParentID sets the "parent" edge to the HandlerExecution entity by ID.
func (heuo *HandlerExecutionUpdateOne) SetParentID(id string) *HandlerExecutionUpdateOne {
	heuo.mutation.SetParentID(id)
	return heuo
}

// SetNillableParentID sets the "parent" edge to the HandlerExecution entity by ID if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableParentID(id *string) *HandlerExecutionUpdateOne {
	if id != nil {
		heuo = heuo.SetParentID(*id)
	}
	return heuo
}

// SetParent sets the "parent" edge to the HandlerExecution entity.
func (heuo *HandlerExecutionUpdateOne) SetParent(h *HandlerExecution) *HandlerExecutionUpdateOne {
	return heuo.SetParentID(h.ID)
}

// AddChildIDs adds the "children" edge to the HandlerExecution entity by IDs.
func (heuo *HandlerExecutionUpdateOne) AddChildIDs(ids ...string) *HandlerExecutionUpdateOne {
	heuo.mutation.AddChildIDs(ids...)
	return heuo
}

// AddChildren adds the "children" edges to the HandlerExecution entity.
func (heuo *HandlerExecutionUpdateOne) AddChildren(h ...*HandlerExecution) *HandlerExecutionUpdateOne {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heuo.AddChildIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the HandlerTask entity by IDs.
func (heuo *HandlerExecutionUpdateOne) AddTaskIDs(ids ...string) *HandlerExecutionUpdateOne {
	heuo.mutation.AddTaskIDs(ids...)
	return heuo
}

// AddTasks adds the "tasks" edges to the HandlerTask entity.
func (heuo *HandlerExecutionUpdateOne) AddTasks(h ...*HandlerTask) *HandlerExecutionUpdateOne {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heuo.AddTaskIDs(ids...)
}

// SetSagaStepExecutionID sets the "saga_step_execution" edge to the SagaStepExecution entity by ID.
func (heuo *HandlerExecutionUpdateOne) SetSagaStepExecutionID(id string) *HandlerExecutionUpdateOne {
	heuo.mutation.SetSagaStepExecutionID(id)
	return heuo
}

// SetNillableSagaStepExecutionID sets the "saga_step_execution" edge to the SagaStepExecution entity by ID if the given value is not nil.
func (heuo *HandlerExecutionUpdateOne) SetNillableSagaStepExecutionID(id *string) *HandlerExecutionUpdateOne {
	if id != nil {
		heuo = heuo.SetSagaStepExecutionID(*id)
	}
	return heuo
}

// SetSagaStepExecution sets the "saga_step_execution" edge to the SagaStepExecution entity.
func (heuo *HandlerExecutionUpdateOne) SetSagaStepExecution(s *SagaStepExecution) *HandlerExecutionUpdateOne {
	return heuo.SetSagaStepExecutionID(s.ID)
}

// Mutation returns the HandlerExecutionMutation object of the builder.
func (heuo *HandlerExecutionUpdateOne) Mutation() *HandlerExecutionMutation {
	return heuo.mutation
}

// ClearExecutionContext clears the "execution_context" edge to the ExecutionContext entity.
func (heuo *HandlerExecutionUpdateOne) ClearExecutionContext() *HandlerExecutionUpdateOne {
	heuo.mutation.ClearExecutionContext()
	return heuo
}

// ClearParent clears the "parent" edge to the HandlerExecution entity.
func (heuo *HandlerExecutionUpdateOne) ClearParent() *HandlerExecutionUpdateOne {
	heuo.mutation.ClearParent()
	return heuo
}

// ClearChildren clears all "children" edges to the HandlerExecution entity.
func (heuo *HandlerExecutionUpdateOne) ClearChildren() *HandlerExecutionUpdateOne {
	heuo.mutation.ClearChildren()
	return heuo
}

// RemoveChildIDs removes the "children" edge to HandlerExecution entities by IDs.
func (heuo *HandlerExecutionUpdateOne) RemoveChildIDs(ids ...string) *HandlerExecutionUpdateOne {
	heuo.mutation.RemoveChildIDs(ids...)
	return heuo
}

// RemoveChildren removes "children" edges to HandlerExecution entities.
func (heuo *HandlerExecutionUpdateOne) RemoveChildren(h ...*HandlerExecution) *HandlerExecutionUpdateOne {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heuo.RemoveChildIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the HandlerTask entity.
func (heuo *HandlerExecutionUpdateOne) ClearTasks() *HandlerExecutionUpdateOne {
	heuo.mutation.ClearTasks()
	return heuo
}

// RemoveTaskIDs removes the "tasks" edge to HandlerTask entities by IDs.
func (heuo *HandlerExecutionUpdateOne) RemoveTaskIDs(ids ...string) *HandlerExecutionUpdateOne {
	heuo.mutation.RemoveTaskIDs(ids...)
	return heuo
}

// RemoveTasks removes "tasks" edges to HandlerTask entities.
func (heuo *HandlerExecutionUpdateOne) RemoveTasks(h ...*HandlerTask) *HandlerExecutionUpdateOne {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return heuo.RemoveTaskIDs(ids...)
}

// ClearSagaStepExecution clears the "saga_step_execution" edge to the SagaStepExecution entity.
func (heuo *HandlerExecutionUpdateOne) ClearSagaStepExecution() *HandlerExecutionUpdateOne {
	heuo.mutation.ClearSagaStepExecution()
	return heuo
}

// Where appends a list predicates to the HandlerExecutionUpdate builder.
func (heuo *HandlerExecutionUpdateOne) Where(ps ...predicate.HandlerExecution) *HandlerExecutionUpdateOne {
	heuo.mutation.Where(ps...)
	return heuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (heuo *HandlerExecutionUpdateOne) Select(field string, fields ...string) *HandlerExecutionUpdateOne {
	heuo.fields = append([]string{field}, fields...)
	return heuo
}

// Save executes the query and returns the updated HandlerExecution entity.
func (heuo *HandlerExecutionUpdateOne) Save(ctx context.Context) (*HandlerExecution, error) {
	return withHooks(ctx, heuo.sqlSave, heuo.mutation, heuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (heuo *HandlerExecutionUpdateOne) SaveX(ctx context.Context) *HandlerExecution {
	node, err := heuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (heuo *HandlerExecutionUpdateOne) Exec(ctx context.Context) error {
	_, err := heuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (heuo *HandlerExecutionUpdateOne) ExecX(ctx context.Context) {
	if err := heuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (heuo *HandlerExecutionUpdateOne) check() error {
	if v, ok := heuo.mutation.Status(); ok {
		if err := handlerexecution.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HandlerExecution.status": %w`, err)}
		}
	}
	if heuo.mutation.ExecutionContextCleared() && len(heuo.mutation.ExecutionContextIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "HandlerExecution.execution_context"`)
	}
	return nil
}

func (heuo *HandlerExecutionUpdateOne) sqlSave(ctx context.Context) (_node *HandlerExecution, err error) {
	if err := heuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(handlerexecution.Table, handlerexecution.Columns, sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString))
	id, ok := heuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HandlerExecution.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := heuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, handlerexecution.FieldID)
		for _, f := range fields {
			if !handlerexecution.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != handlerexecution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := heuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := heuo.mutation.RunID(); ok {
		_spec.SetField(handlerexecution.FieldRunID, field.TypeString, value)
	}
	if value, ok := heuo.mutation.HandlerName(); ok {
		_spec.SetField(handlerexecution.FieldHandlerName, field.TypeString, value)
	}
	if value, ok := heuo.mutation.Status(); ok {
		_spec.SetField(handlerexecution.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := heuo.mutation.StartTime(); ok {
		_spec.SetField(handlerexecution.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := heuo.mutation.EndTime(); ok {
		_spec.SetField(handlerexecution.FieldEndTime, field.TypeTime, value)
	}
	if heuo.mutation.EndTimeCleared() {
		_spec.ClearField(handlerexecution.FieldEndTime, field.TypeTime)
	}
	if value, ok := heuo.mutation.RetryCount(); ok {
		_spec.SetField(handlerexecution.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := heuo.mutation.AddedRetryCount(); ok {
		_spec.AddField(handlerexecution.FieldRetryCount, field.TypeInt, value)
	}
	if value, ok := heuo.mutation.MaxRetries(); ok {
		_spec.SetField(handlerexecution.FieldMaxRetries, field.TypeInt, value)
	}
	if value, ok := heuo.mutation.AddedMaxRetries(); ok {
		_spec.AddField(handlerexecution.FieldMaxRetries, field.TypeInt, value)
	}
	if heuo.mutation.ExecutionContextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ExecutionContextTable,
			Columns: []string{handlerexecution.ExecutionContextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(executioncontext.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.ExecutionContextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ExecutionContextTable,
			Columns: []string{handlerexecution.ExecutionContextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(executioncontext.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ParentTable,
			Columns: []string{handlerexecution.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlerexecution.ParentTable,
			Columns: []string{handlerexecution.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.ChildrenTable,
			Columns: []string{handlerexecution.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !heuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.ChildrenTable,
			Columns: []string{handlerexecution.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.ChildrenTable,
			Columns: []string{handlerexecution.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.TasksTable,
			Columns: []string{handlerexecution.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !heuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.TasksTable,
			Columns: []string{handlerexecution.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   handlerexecution.TasksTable,
			Columns: []string{handlerexecution.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.SagaStepExecutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   handlerexecution.SagaStepExecutionTable,
			Columns: []string{handlerexecution.SagaStepExecutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sagastepexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.SagaStepExecutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   handlerexecution.SagaStepExecutionTable,
			Columns: []string{handlerexecution.SagaStepExecutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sagastepexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HandlerExecution{config: heuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, heuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{handlerexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	heuo.mutation.done = true
	return _node, nil
}
