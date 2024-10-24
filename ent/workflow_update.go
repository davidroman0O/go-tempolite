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
	"github.com/davidroman0O/tempolite/ent/predicate"
	"github.com/davidroman0O/tempolite/ent/schema"
	"github.com/davidroman0O/tempolite/ent/workflow"
	"github.com/davidroman0O/tempolite/ent/workflowexecution"
)

// WorkflowUpdate is the builder for updating Workflow entities.
type WorkflowUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowMutation
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wu *WorkflowUpdate) Where(ps ...predicate.Workflow) *WorkflowUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetStepID sets the "step_id" field.
func (wu *WorkflowUpdate) SetStepID(s string) *WorkflowUpdate {
	wu.mutation.SetStepID(s)
	return wu
}

// SetNillableStepID sets the "step_id" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableStepID(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetStepID(*s)
	}
	return wu
}

// SetStatus sets the "status" field.
func (wu *WorkflowUpdate) SetStatus(w workflow.Status) *WorkflowUpdate {
	wu.mutation.SetStatus(w)
	return wu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableStatus(w *workflow.Status) *WorkflowUpdate {
	if w != nil {
		wu.SetStatus(*w)
	}
	return wu
}

// SetIdentity sets the "identity" field.
func (wu *WorkflowUpdate) SetIdentity(s string) *WorkflowUpdate {
	wu.mutation.SetIdentity(s)
	return wu
}

// SetNillableIdentity sets the "identity" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableIdentity(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetIdentity(*s)
	}
	return wu
}

// SetHandlerName sets the "handler_name" field.
func (wu *WorkflowUpdate) SetHandlerName(s string) *WorkflowUpdate {
	wu.mutation.SetHandlerName(s)
	return wu
}

// SetNillableHandlerName sets the "handler_name" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableHandlerName(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetHandlerName(*s)
	}
	return wu
}

// SetInput sets the "input" field.
func (wu *WorkflowUpdate) SetInput(i []interface{}) *WorkflowUpdate {
	wu.mutation.SetInput(i)
	return wu
}

// AppendInput appends i to the "input" field.
func (wu *WorkflowUpdate) AppendInput(i []interface{}) *WorkflowUpdate {
	wu.mutation.AppendInput(i)
	return wu
}

// SetRetryPolicy sets the "retry_policy" field.
func (wu *WorkflowUpdate) SetRetryPolicy(sp schema.RetryPolicy) *WorkflowUpdate {
	wu.mutation.SetRetryPolicy(sp)
	return wu
}

// SetNillableRetryPolicy sets the "retry_policy" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableRetryPolicy(sp *schema.RetryPolicy) *WorkflowUpdate {
	if sp != nil {
		wu.SetRetryPolicy(*sp)
	}
	return wu
}

// ClearRetryPolicy clears the value of the "retry_policy" field.
func (wu *WorkflowUpdate) ClearRetryPolicy() *WorkflowUpdate {
	wu.mutation.ClearRetryPolicy()
	return wu
}

// SetIsPaused sets the "is_paused" field.
func (wu *WorkflowUpdate) SetIsPaused(b bool) *WorkflowUpdate {
	wu.mutation.SetIsPaused(b)
	return wu
}

// SetNillableIsPaused sets the "is_paused" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableIsPaused(b *bool) *WorkflowUpdate {
	if b != nil {
		wu.SetIsPaused(*b)
	}
	return wu
}

// SetIsReady sets the "is_ready" field.
func (wu *WorkflowUpdate) SetIsReady(b bool) *WorkflowUpdate {
	wu.mutation.SetIsReady(b)
	return wu
}

// SetNillableIsReady sets the "is_ready" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableIsReady(b *bool) *WorkflowUpdate {
	if b != nil {
		wu.SetIsReady(*b)
	}
	return wu
}

// SetTimeout sets the "timeout" field.
func (wu *WorkflowUpdate) SetTimeout(t time.Time) *WorkflowUpdate {
	wu.mutation.SetTimeout(t)
	return wu
}

// SetNillableTimeout sets the "timeout" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableTimeout(t *time.Time) *WorkflowUpdate {
	if t != nil {
		wu.SetTimeout(*t)
	}
	return wu
}

// ClearTimeout clears the value of the "timeout" field.
func (wu *WorkflowUpdate) ClearTimeout() *WorkflowUpdate {
	wu.mutation.ClearTimeout()
	return wu
}

// SetCreatedAt sets the "created_at" field.
func (wu *WorkflowUpdate) SetCreatedAt(t time.Time) *WorkflowUpdate {
	wu.mutation.SetCreatedAt(t)
	return wu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableCreatedAt(t *time.Time) *WorkflowUpdate {
	if t != nil {
		wu.SetCreatedAt(*t)
	}
	return wu
}

// SetContinuedFromID sets the "continued_from_id" field.
func (wu *WorkflowUpdate) SetContinuedFromID(s string) *WorkflowUpdate {
	wu.mutation.SetContinuedFromID(s)
	return wu
}

// SetNillableContinuedFromID sets the "continued_from_id" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableContinuedFromID(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetContinuedFromID(*s)
	}
	return wu
}

// ClearContinuedFromID clears the value of the "continued_from_id" field.
func (wu *WorkflowUpdate) ClearContinuedFromID() *WorkflowUpdate {
	wu.mutation.ClearContinuedFromID()
	return wu
}

// SetRetriedFromID sets the "retried_from_id" field.
func (wu *WorkflowUpdate) SetRetriedFromID(s string) *WorkflowUpdate {
	wu.mutation.SetRetriedFromID(s)
	return wu
}

// SetNillableRetriedFromID sets the "retried_from_id" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableRetriedFromID(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetRetriedFromID(*s)
	}
	return wu
}

// ClearRetriedFromID clears the value of the "retried_from_id" field.
func (wu *WorkflowUpdate) ClearRetriedFromID() *WorkflowUpdate {
	wu.mutation.ClearRetriedFromID()
	return wu
}

// AddExecutionIDs adds the "executions" edge to the WorkflowExecution entity by IDs.
func (wu *WorkflowUpdate) AddExecutionIDs(ids ...string) *WorkflowUpdate {
	wu.mutation.AddExecutionIDs(ids...)
	return wu
}

// AddExecutions adds the "executions" edges to the WorkflowExecution entity.
func (wu *WorkflowUpdate) AddExecutions(w ...*WorkflowExecution) *WorkflowUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddExecutionIDs(ids...)
}

// SetContinuedFrom sets the "continued_from" edge to the Workflow entity.
func (wu *WorkflowUpdate) SetContinuedFrom(w *Workflow) *WorkflowUpdate {
	return wu.SetContinuedFromID(w.ID)
}

// SetContinuedToID sets the "continued_to" edge to the Workflow entity by ID.
func (wu *WorkflowUpdate) SetContinuedToID(id string) *WorkflowUpdate {
	wu.mutation.SetContinuedToID(id)
	return wu
}

// SetNillableContinuedToID sets the "continued_to" edge to the Workflow entity by ID if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableContinuedToID(id *string) *WorkflowUpdate {
	if id != nil {
		wu = wu.SetContinuedToID(*id)
	}
	return wu
}

// SetContinuedTo sets the "continued_to" edge to the Workflow entity.
func (wu *WorkflowUpdate) SetContinuedTo(w *Workflow) *WorkflowUpdate {
	return wu.SetContinuedToID(w.ID)
}

// SetRetriedFrom sets the "retried_from" edge to the Workflow entity.
func (wu *WorkflowUpdate) SetRetriedFrom(w *Workflow) *WorkflowUpdate {
	return wu.SetRetriedFromID(w.ID)
}

// AddRetriedToIDs adds the "retried_to" edge to the Workflow entity by IDs.
func (wu *WorkflowUpdate) AddRetriedToIDs(ids ...string) *WorkflowUpdate {
	wu.mutation.AddRetriedToIDs(ids...)
	return wu
}

// AddRetriedTo adds the "retried_to" edges to the Workflow entity.
func (wu *WorkflowUpdate) AddRetriedTo(w ...*Workflow) *WorkflowUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddRetriedToIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wu *WorkflowUpdate) Mutation() *WorkflowMutation {
	return wu.mutation
}

// ClearExecutions clears all "executions" edges to the WorkflowExecution entity.
func (wu *WorkflowUpdate) ClearExecutions() *WorkflowUpdate {
	wu.mutation.ClearExecutions()
	return wu
}

// RemoveExecutionIDs removes the "executions" edge to WorkflowExecution entities by IDs.
func (wu *WorkflowUpdate) RemoveExecutionIDs(ids ...string) *WorkflowUpdate {
	wu.mutation.RemoveExecutionIDs(ids...)
	return wu
}

// RemoveExecutions removes "executions" edges to WorkflowExecution entities.
func (wu *WorkflowUpdate) RemoveExecutions(w ...*WorkflowExecution) *WorkflowUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveExecutionIDs(ids...)
}

// ClearContinuedFrom clears the "continued_from" edge to the Workflow entity.
func (wu *WorkflowUpdate) ClearContinuedFrom() *WorkflowUpdate {
	wu.mutation.ClearContinuedFrom()
	return wu
}

// ClearContinuedTo clears the "continued_to" edge to the Workflow entity.
func (wu *WorkflowUpdate) ClearContinuedTo() *WorkflowUpdate {
	wu.mutation.ClearContinuedTo()
	return wu
}

// ClearRetriedFrom clears the "retried_from" edge to the Workflow entity.
func (wu *WorkflowUpdate) ClearRetriedFrom() *WorkflowUpdate {
	wu.mutation.ClearRetriedFrom()
	return wu
}

// ClearRetriedTo clears all "retried_to" edges to the Workflow entity.
func (wu *WorkflowUpdate) ClearRetriedTo() *WorkflowUpdate {
	wu.mutation.ClearRetriedTo()
	return wu
}

// RemoveRetriedToIDs removes the "retried_to" edge to Workflow entities by IDs.
func (wu *WorkflowUpdate) RemoveRetriedToIDs(ids ...string) *WorkflowUpdate {
	wu.mutation.RemoveRetriedToIDs(ids...)
	return wu
}

// RemoveRetriedTo removes "retried_to" edges to Workflow entities.
func (wu *WorkflowUpdate) RemoveRetriedTo(w ...*Workflow) *WorkflowUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveRetriedToIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkflowUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkflowUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkflowUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkflowUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkflowUpdate) check() error {
	if v, ok := wu.mutation.StepID(); ok {
		if err := workflow.StepIDValidator(v); err != nil {
			return &ValidationError{Name: "step_id", err: fmt.Errorf(`ent: validator failed for field "Workflow.step_id": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Status(); ok {
		if err := workflow.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Workflow.status": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Identity(); ok {
		if err := workflow.IdentityValidator(v); err != nil {
			return &ValidationError{Name: "identity", err: fmt.Errorf(`ent: validator failed for field "Workflow.identity": %w`, err)}
		}
	}
	if v, ok := wu.mutation.HandlerName(); ok {
		if err := workflow.HandlerNameValidator(v); err != nil {
			return &ValidationError{Name: "handler_name", err: fmt.Errorf(`ent: validator failed for field "Workflow.handler_name": %w`, err)}
		}
	}
	return nil
}

func (wu *WorkflowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.StepID(); ok {
		_spec.SetField(workflow.FieldStepID, field.TypeString, value)
	}
	if value, ok := wu.mutation.Status(); ok {
		_spec.SetField(workflow.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := wu.mutation.Identity(); ok {
		_spec.SetField(workflow.FieldIdentity, field.TypeString, value)
	}
	if value, ok := wu.mutation.HandlerName(); ok {
		_spec.SetField(workflow.FieldHandlerName, field.TypeString, value)
	}
	if value, ok := wu.mutation.Input(); ok {
		_spec.SetField(workflow.FieldInput, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedInput(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldInput, value)
		})
	}
	if value, ok := wu.mutation.RetryPolicy(); ok {
		_spec.SetField(workflow.FieldRetryPolicy, field.TypeJSON, value)
	}
	if wu.mutation.RetryPolicyCleared() {
		_spec.ClearField(workflow.FieldRetryPolicy, field.TypeJSON)
	}
	if value, ok := wu.mutation.IsPaused(); ok {
		_spec.SetField(workflow.FieldIsPaused, field.TypeBool, value)
	}
	if value, ok := wu.mutation.IsReady(); ok {
		_spec.SetField(workflow.FieldIsReady, field.TypeBool, value)
	}
	if value, ok := wu.mutation.Timeout(); ok {
		_spec.SetField(workflow.FieldTimeout, field.TypeTime, value)
	}
	if wu.mutation.TimeoutCleared() {
		_spec.ClearField(workflow.FieldTimeout, field.TypeTime)
	}
	if value, ok := wu.mutation.CreatedAt(); ok {
		_spec.SetField(workflow.FieldCreatedAt, field.TypeTime, value)
	}
	if wu.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedExecutionsIDs(); len(nodes) > 0 && !wu.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ExecutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.ContinuedFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workflow.ContinuedFromTable,
			Columns: []string{workflow.ContinuedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ContinuedFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workflow.ContinuedFromTable,
			Columns: []string{workflow.ContinuedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.ContinuedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workflow.ContinuedToTable,
			Columns: []string{workflow.ContinuedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ContinuedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workflow.ContinuedToTable,
			Columns: []string{workflow.ContinuedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.RetriedFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.RetriedFromTable,
			Columns: []string{workflow.RetriedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RetriedFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.RetriedFromTable,
			Columns: []string{workflow.RetriedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.RetriedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RetriedToTable,
			Columns: []string{workflow.RetriedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedRetriedToIDs(); len(nodes) > 0 && !wu.mutation.RetriedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RetriedToTable,
			Columns: []string{workflow.RetriedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RetriedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RetriedToTable,
			Columns: []string{workflow.RetriedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WorkflowUpdateOne is the builder for updating a single Workflow entity.
type WorkflowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowMutation
}

// SetStepID sets the "step_id" field.
func (wuo *WorkflowUpdateOne) SetStepID(s string) *WorkflowUpdateOne {
	wuo.mutation.SetStepID(s)
	return wuo
}

// SetNillableStepID sets the "step_id" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableStepID(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetStepID(*s)
	}
	return wuo
}

// SetStatus sets the "status" field.
func (wuo *WorkflowUpdateOne) SetStatus(w workflow.Status) *WorkflowUpdateOne {
	wuo.mutation.SetStatus(w)
	return wuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableStatus(w *workflow.Status) *WorkflowUpdateOne {
	if w != nil {
		wuo.SetStatus(*w)
	}
	return wuo
}

// SetIdentity sets the "identity" field.
func (wuo *WorkflowUpdateOne) SetIdentity(s string) *WorkflowUpdateOne {
	wuo.mutation.SetIdentity(s)
	return wuo
}

// SetNillableIdentity sets the "identity" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableIdentity(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetIdentity(*s)
	}
	return wuo
}

// SetHandlerName sets the "handler_name" field.
func (wuo *WorkflowUpdateOne) SetHandlerName(s string) *WorkflowUpdateOne {
	wuo.mutation.SetHandlerName(s)
	return wuo
}

// SetNillableHandlerName sets the "handler_name" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableHandlerName(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetHandlerName(*s)
	}
	return wuo
}

// SetInput sets the "input" field.
func (wuo *WorkflowUpdateOne) SetInput(i []interface{}) *WorkflowUpdateOne {
	wuo.mutation.SetInput(i)
	return wuo
}

// AppendInput appends i to the "input" field.
func (wuo *WorkflowUpdateOne) AppendInput(i []interface{}) *WorkflowUpdateOne {
	wuo.mutation.AppendInput(i)
	return wuo
}

// SetRetryPolicy sets the "retry_policy" field.
func (wuo *WorkflowUpdateOne) SetRetryPolicy(sp schema.RetryPolicy) *WorkflowUpdateOne {
	wuo.mutation.SetRetryPolicy(sp)
	return wuo
}

// SetNillableRetryPolicy sets the "retry_policy" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableRetryPolicy(sp *schema.RetryPolicy) *WorkflowUpdateOne {
	if sp != nil {
		wuo.SetRetryPolicy(*sp)
	}
	return wuo
}

// ClearRetryPolicy clears the value of the "retry_policy" field.
func (wuo *WorkflowUpdateOne) ClearRetryPolicy() *WorkflowUpdateOne {
	wuo.mutation.ClearRetryPolicy()
	return wuo
}

// SetIsPaused sets the "is_paused" field.
func (wuo *WorkflowUpdateOne) SetIsPaused(b bool) *WorkflowUpdateOne {
	wuo.mutation.SetIsPaused(b)
	return wuo
}

// SetNillableIsPaused sets the "is_paused" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableIsPaused(b *bool) *WorkflowUpdateOne {
	if b != nil {
		wuo.SetIsPaused(*b)
	}
	return wuo
}

// SetIsReady sets the "is_ready" field.
func (wuo *WorkflowUpdateOne) SetIsReady(b bool) *WorkflowUpdateOne {
	wuo.mutation.SetIsReady(b)
	return wuo
}

// SetNillableIsReady sets the "is_ready" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableIsReady(b *bool) *WorkflowUpdateOne {
	if b != nil {
		wuo.SetIsReady(*b)
	}
	return wuo
}

// SetTimeout sets the "timeout" field.
func (wuo *WorkflowUpdateOne) SetTimeout(t time.Time) *WorkflowUpdateOne {
	wuo.mutation.SetTimeout(t)
	return wuo
}

// SetNillableTimeout sets the "timeout" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableTimeout(t *time.Time) *WorkflowUpdateOne {
	if t != nil {
		wuo.SetTimeout(*t)
	}
	return wuo
}

// ClearTimeout clears the value of the "timeout" field.
func (wuo *WorkflowUpdateOne) ClearTimeout() *WorkflowUpdateOne {
	wuo.mutation.ClearTimeout()
	return wuo
}

// SetCreatedAt sets the "created_at" field.
func (wuo *WorkflowUpdateOne) SetCreatedAt(t time.Time) *WorkflowUpdateOne {
	wuo.mutation.SetCreatedAt(t)
	return wuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableCreatedAt(t *time.Time) *WorkflowUpdateOne {
	if t != nil {
		wuo.SetCreatedAt(*t)
	}
	return wuo
}

// SetContinuedFromID sets the "continued_from_id" field.
func (wuo *WorkflowUpdateOne) SetContinuedFromID(s string) *WorkflowUpdateOne {
	wuo.mutation.SetContinuedFromID(s)
	return wuo
}

// SetNillableContinuedFromID sets the "continued_from_id" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableContinuedFromID(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetContinuedFromID(*s)
	}
	return wuo
}

// ClearContinuedFromID clears the value of the "continued_from_id" field.
func (wuo *WorkflowUpdateOne) ClearContinuedFromID() *WorkflowUpdateOne {
	wuo.mutation.ClearContinuedFromID()
	return wuo
}

// SetRetriedFromID sets the "retried_from_id" field.
func (wuo *WorkflowUpdateOne) SetRetriedFromID(s string) *WorkflowUpdateOne {
	wuo.mutation.SetRetriedFromID(s)
	return wuo
}

// SetNillableRetriedFromID sets the "retried_from_id" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableRetriedFromID(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetRetriedFromID(*s)
	}
	return wuo
}

// ClearRetriedFromID clears the value of the "retried_from_id" field.
func (wuo *WorkflowUpdateOne) ClearRetriedFromID() *WorkflowUpdateOne {
	wuo.mutation.ClearRetriedFromID()
	return wuo
}

// AddExecutionIDs adds the "executions" edge to the WorkflowExecution entity by IDs.
func (wuo *WorkflowUpdateOne) AddExecutionIDs(ids ...string) *WorkflowUpdateOne {
	wuo.mutation.AddExecutionIDs(ids...)
	return wuo
}

// AddExecutions adds the "executions" edges to the WorkflowExecution entity.
func (wuo *WorkflowUpdateOne) AddExecutions(w ...*WorkflowExecution) *WorkflowUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddExecutionIDs(ids...)
}

// SetContinuedFrom sets the "continued_from" edge to the Workflow entity.
func (wuo *WorkflowUpdateOne) SetContinuedFrom(w *Workflow) *WorkflowUpdateOne {
	return wuo.SetContinuedFromID(w.ID)
}

// SetContinuedToID sets the "continued_to" edge to the Workflow entity by ID.
func (wuo *WorkflowUpdateOne) SetContinuedToID(id string) *WorkflowUpdateOne {
	wuo.mutation.SetContinuedToID(id)
	return wuo
}

// SetNillableContinuedToID sets the "continued_to" edge to the Workflow entity by ID if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableContinuedToID(id *string) *WorkflowUpdateOne {
	if id != nil {
		wuo = wuo.SetContinuedToID(*id)
	}
	return wuo
}

// SetContinuedTo sets the "continued_to" edge to the Workflow entity.
func (wuo *WorkflowUpdateOne) SetContinuedTo(w *Workflow) *WorkflowUpdateOne {
	return wuo.SetContinuedToID(w.ID)
}

// SetRetriedFrom sets the "retried_from" edge to the Workflow entity.
func (wuo *WorkflowUpdateOne) SetRetriedFrom(w *Workflow) *WorkflowUpdateOne {
	return wuo.SetRetriedFromID(w.ID)
}

// AddRetriedToIDs adds the "retried_to" edge to the Workflow entity by IDs.
func (wuo *WorkflowUpdateOne) AddRetriedToIDs(ids ...string) *WorkflowUpdateOne {
	wuo.mutation.AddRetriedToIDs(ids...)
	return wuo
}

// AddRetriedTo adds the "retried_to" edges to the Workflow entity.
func (wuo *WorkflowUpdateOne) AddRetriedTo(w ...*Workflow) *WorkflowUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddRetriedToIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wuo *WorkflowUpdateOne) Mutation() *WorkflowMutation {
	return wuo.mutation
}

// ClearExecutions clears all "executions" edges to the WorkflowExecution entity.
func (wuo *WorkflowUpdateOne) ClearExecutions() *WorkflowUpdateOne {
	wuo.mutation.ClearExecutions()
	return wuo
}

// RemoveExecutionIDs removes the "executions" edge to WorkflowExecution entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveExecutionIDs(ids ...string) *WorkflowUpdateOne {
	wuo.mutation.RemoveExecutionIDs(ids...)
	return wuo
}

// RemoveExecutions removes "executions" edges to WorkflowExecution entities.
func (wuo *WorkflowUpdateOne) RemoveExecutions(w ...*WorkflowExecution) *WorkflowUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveExecutionIDs(ids...)
}

// ClearContinuedFrom clears the "continued_from" edge to the Workflow entity.
func (wuo *WorkflowUpdateOne) ClearContinuedFrom() *WorkflowUpdateOne {
	wuo.mutation.ClearContinuedFrom()
	return wuo
}

// ClearContinuedTo clears the "continued_to" edge to the Workflow entity.
func (wuo *WorkflowUpdateOne) ClearContinuedTo() *WorkflowUpdateOne {
	wuo.mutation.ClearContinuedTo()
	return wuo
}

// ClearRetriedFrom clears the "retried_from" edge to the Workflow entity.
func (wuo *WorkflowUpdateOne) ClearRetriedFrom() *WorkflowUpdateOne {
	wuo.mutation.ClearRetriedFrom()
	return wuo
}

// ClearRetriedTo clears all "retried_to" edges to the Workflow entity.
func (wuo *WorkflowUpdateOne) ClearRetriedTo() *WorkflowUpdateOne {
	wuo.mutation.ClearRetriedTo()
	return wuo
}

// RemoveRetriedToIDs removes the "retried_to" edge to Workflow entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveRetriedToIDs(ids ...string) *WorkflowUpdateOne {
	wuo.mutation.RemoveRetriedToIDs(ids...)
	return wuo
}

// RemoveRetriedTo removes "retried_to" edges to Workflow entities.
func (wuo *WorkflowUpdateOne) RemoveRetriedTo(w ...*Workflow) *WorkflowUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveRetriedToIDs(ids...)
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wuo *WorkflowUpdateOne) Where(ps ...predicate.Workflow) *WorkflowUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkflowUpdateOne) Select(field string, fields ...string) *WorkflowUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workflow entity.
func (wuo *WorkflowUpdateOne) Save(ctx context.Context) (*Workflow, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) SaveX(ctx context.Context) *Workflow {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkflowUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkflowUpdateOne) check() error {
	if v, ok := wuo.mutation.StepID(); ok {
		if err := workflow.StepIDValidator(v); err != nil {
			return &ValidationError{Name: "step_id", err: fmt.Errorf(`ent: validator failed for field "Workflow.step_id": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Status(); ok {
		if err := workflow.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Workflow.status": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Identity(); ok {
		if err := workflow.IdentityValidator(v); err != nil {
			return &ValidationError{Name: "identity", err: fmt.Errorf(`ent: validator failed for field "Workflow.identity": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.HandlerName(); ok {
		if err := workflow.HandlerNameValidator(v); err != nil {
			return &ValidationError{Name: "handler_name", err: fmt.Errorf(`ent: validator failed for field "Workflow.handler_name": %w`, err)}
		}
	}
	return nil
}

func (wuo *WorkflowUpdateOne) sqlSave(ctx context.Context) (_node *Workflow, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Workflow.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.FieldID)
		for _, f := range fields {
			if !workflow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.StepID(); ok {
		_spec.SetField(workflow.FieldStepID, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Status(); ok {
		_spec.SetField(workflow.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := wuo.mutation.Identity(); ok {
		_spec.SetField(workflow.FieldIdentity, field.TypeString, value)
	}
	if value, ok := wuo.mutation.HandlerName(); ok {
		_spec.SetField(workflow.FieldHandlerName, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Input(); ok {
		_spec.SetField(workflow.FieldInput, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedInput(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldInput, value)
		})
	}
	if value, ok := wuo.mutation.RetryPolicy(); ok {
		_spec.SetField(workflow.FieldRetryPolicy, field.TypeJSON, value)
	}
	if wuo.mutation.RetryPolicyCleared() {
		_spec.ClearField(workflow.FieldRetryPolicy, field.TypeJSON)
	}
	if value, ok := wuo.mutation.IsPaused(); ok {
		_spec.SetField(workflow.FieldIsPaused, field.TypeBool, value)
	}
	if value, ok := wuo.mutation.IsReady(); ok {
		_spec.SetField(workflow.FieldIsReady, field.TypeBool, value)
	}
	if value, ok := wuo.mutation.Timeout(); ok {
		_spec.SetField(workflow.FieldTimeout, field.TypeTime, value)
	}
	if wuo.mutation.TimeoutCleared() {
		_spec.ClearField(workflow.FieldTimeout, field.TypeTime)
	}
	if value, ok := wuo.mutation.CreatedAt(); ok {
		_spec.SetField(workflow.FieldCreatedAt, field.TypeTime, value)
	}
	if wuo.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedExecutionsIDs(); len(nodes) > 0 && !wuo.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ExecutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.ContinuedFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workflow.ContinuedFromTable,
			Columns: []string{workflow.ContinuedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ContinuedFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workflow.ContinuedFromTable,
			Columns: []string{workflow.ContinuedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.ContinuedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workflow.ContinuedToTable,
			Columns: []string{workflow.ContinuedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ContinuedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workflow.ContinuedToTable,
			Columns: []string{workflow.ContinuedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.RetriedFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.RetriedFromTable,
			Columns: []string{workflow.RetriedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RetriedFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.RetriedFromTable,
			Columns: []string{workflow.RetriedFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.RetriedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RetriedToTable,
			Columns: []string{workflow.RetriedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedRetriedToIDs(); len(nodes) > 0 && !wuo.mutation.RetriedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RetriedToTable,
			Columns: []string{workflow.RetriedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RetriedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RetriedToTable,
			Columns: []string{workflow.RetriedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Workflow{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
