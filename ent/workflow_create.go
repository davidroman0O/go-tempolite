// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/activity"
	"github.com/davidroman0O/go-tempolite/ent/schema"
	"github.com/davidroman0O/go-tempolite/ent/workflow"
	"github.com/davidroman0O/go-tempolite/ent/workflowexecution"
)

// WorkflowCreate is the builder for creating a Workflow entity.
type WorkflowCreate struct {
	config
	mutation *WorkflowMutation
	hooks    []Hook
}

// SetIdentity sets the "identity" field.
func (wc *WorkflowCreate) SetIdentity(s string) *WorkflowCreate {
	wc.mutation.SetIdentity(s)
	return wc
}

// SetHandlerName sets the "handler_name" field.
func (wc *WorkflowCreate) SetHandlerName(s string) *WorkflowCreate {
	wc.mutation.SetHandlerName(s)
	return wc
}

// SetInput sets the "input" field.
func (wc *WorkflowCreate) SetInput(i []interface{}) *WorkflowCreate {
	wc.mutation.SetInput(i)
	return wc
}

// SetRetryPolicy sets the "retry_policy" field.
func (wc *WorkflowCreate) SetRetryPolicy(sp schema.RetryPolicy) *WorkflowCreate {
	wc.mutation.SetRetryPolicy(sp)
	return wc
}

// SetNillableRetryPolicy sets the "retry_policy" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableRetryPolicy(sp *schema.RetryPolicy) *WorkflowCreate {
	if sp != nil {
		wc.SetRetryPolicy(*sp)
	}
	return wc
}

// SetTimeout sets the "timeout" field.
func (wc *WorkflowCreate) SetTimeout(t time.Time) *WorkflowCreate {
	wc.mutation.SetTimeout(t)
	return wc
}

// SetNillableTimeout sets the "timeout" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableTimeout(t *time.Time) *WorkflowCreate {
	if t != nil {
		wc.SetTimeout(*t)
	}
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkflowCreate) SetCreatedAt(t time.Time) *WorkflowCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableCreatedAt(t *time.Time) *WorkflowCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkflowCreate) SetID(s string) *WorkflowCreate {
	wc.mutation.SetID(s)
	return wc
}

// AddExecutionIDs adds the "executions" edge to the WorkflowExecution entity by IDs.
func (wc *WorkflowCreate) AddExecutionIDs(ids ...string) *WorkflowCreate {
	wc.mutation.AddExecutionIDs(ids...)
	return wc
}

// AddExecutions adds the "executions" edges to the WorkflowExecution entity.
func (wc *WorkflowCreate) AddExecutions(w ...*WorkflowExecution) *WorkflowCreate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddExecutionIDs(ids...)
}

// AddActivityIDs adds the "activities" edge to the Activity entity by IDs.
func (wc *WorkflowCreate) AddActivityIDs(ids ...string) *WorkflowCreate {
	wc.mutation.AddActivityIDs(ids...)
	return wc
}

// AddActivities adds the "activities" edges to the Activity entity.
func (wc *WorkflowCreate) AddActivities(a ...*Activity) *WorkflowCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wc.AddActivityIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wc *WorkflowCreate) Mutation() *WorkflowMutation {
	return wc.mutation
}

// Save creates the Workflow in the database.
func (wc *WorkflowCreate) Save(ctx context.Context) (*Workflow, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkflowCreate) SaveX(ctx context.Context) *Workflow {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkflowCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkflowCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkflowCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workflow.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkflowCreate) check() error {
	if _, ok := wc.mutation.Identity(); !ok {
		return &ValidationError{Name: "identity", err: errors.New(`ent: missing required field "Workflow.identity"`)}
	}
	if v, ok := wc.mutation.Identity(); ok {
		if err := workflow.IdentityValidator(v); err != nil {
			return &ValidationError{Name: "identity", err: fmt.Errorf(`ent: validator failed for field "Workflow.identity": %w`, err)}
		}
	}
	if _, ok := wc.mutation.HandlerName(); !ok {
		return &ValidationError{Name: "handler_name", err: errors.New(`ent: missing required field "Workflow.handler_name"`)}
	}
	if v, ok := wc.mutation.HandlerName(); ok {
		if err := workflow.HandlerNameValidator(v); err != nil {
			return &ValidationError{Name: "handler_name", err: fmt.Errorf(`ent: validator failed for field "Workflow.handler_name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Input(); !ok {
		return &ValidationError{Name: "input", err: errors.New(`ent: missing required field "Workflow.input"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Workflow.created_at"`)}
	}
	return nil
}

func (wc *WorkflowCreate) sqlSave(ctx context.Context) (*Workflow, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Workflow.ID type: %T", _spec.ID.Value)
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkflowCreate) createSpec() (*Workflow, *sqlgraph.CreateSpec) {
	var (
		_node = &Workflow{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workflow.Table, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.Identity(); ok {
		_spec.SetField(workflow.FieldIdentity, field.TypeString, value)
		_node.Identity = value
	}
	if value, ok := wc.mutation.HandlerName(); ok {
		_spec.SetField(workflow.FieldHandlerName, field.TypeString, value)
		_node.HandlerName = value
	}
	if value, ok := wc.mutation.Input(); ok {
		_spec.SetField(workflow.FieldInput, field.TypeJSON, value)
		_node.Input = value
	}
	if value, ok := wc.mutation.RetryPolicy(); ok {
		_spec.SetField(workflow.FieldRetryPolicy, field.TypeJSON, value)
		_node.RetryPolicy = value
	}
	if value, ok := wc.mutation.Timeout(); ok {
		_spec.SetField(workflow.FieldTimeout, field.TypeTime, value)
		_node.Timeout = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workflow.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := wc.mutation.ExecutionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ActivitiesTable,
			Columns: []string{workflow.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkflowCreateBulk is the builder for creating many Workflow entities in bulk.
type WorkflowCreateBulk struct {
	config
	err      error
	builders []*WorkflowCreate
}

// Save creates the Workflow entities in the database.
func (wcb *WorkflowCreateBulk) Save(ctx context.Context) ([]*Workflow, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workflow, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) SaveX(ctx context.Context) []*Workflow {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkflowCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
