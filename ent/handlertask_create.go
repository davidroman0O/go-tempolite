// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/handlerexecution"
	"github.com/davidroman0O/go-tempolite/ent/handlertask"
)

// HandlerTaskCreate is the builder for creating a HandlerTask entity.
type HandlerTaskCreate struct {
	config
	mutation *HandlerTaskMutation
	hooks    []Hook
}

// SetHandlerName sets the "handler_name" field.
func (htc *HandlerTaskCreate) SetHandlerName(s string) *HandlerTaskCreate {
	htc.mutation.SetHandlerName(s)
	return htc
}

// SetPayload sets the "payload" field.
func (htc *HandlerTaskCreate) SetPayload(b []byte) *HandlerTaskCreate {
	htc.mutation.SetPayload(b)
	return htc
}

// SetResult sets the "result" field.
func (htc *HandlerTaskCreate) SetResult(b []byte) *HandlerTaskCreate {
	htc.mutation.SetResult(b)
	return htc
}

// SetError sets the "error" field.
func (htc *HandlerTaskCreate) SetError(b []byte) *HandlerTaskCreate {
	htc.mutation.SetError(b)
	return htc
}

// SetStatus sets the "status" field.
func (htc *HandlerTaskCreate) SetStatus(h handlertask.Status) *HandlerTaskCreate {
	htc.mutation.SetStatus(h)
	return htc
}

// SetCreatedAt sets the "created_at" field.
func (htc *HandlerTaskCreate) SetCreatedAt(t time.Time) *HandlerTaskCreate {
	htc.mutation.SetCreatedAt(t)
	return htc
}

// SetCompletedAt sets the "completed_at" field.
func (htc *HandlerTaskCreate) SetCompletedAt(t time.Time) *HandlerTaskCreate {
	htc.mutation.SetCompletedAt(t)
	return htc
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (htc *HandlerTaskCreate) SetNillableCompletedAt(t *time.Time) *HandlerTaskCreate {
	if t != nil {
		htc.SetCompletedAt(*t)
	}
	return htc
}

// SetID sets the "id" field.
func (htc *HandlerTaskCreate) SetID(s string) *HandlerTaskCreate {
	htc.mutation.SetID(s)
	return htc
}

// SetHandlerExecutionID sets the "handler_execution" edge to the HandlerExecution entity by ID.
func (htc *HandlerTaskCreate) SetHandlerExecutionID(id string) *HandlerTaskCreate {
	htc.mutation.SetHandlerExecutionID(id)
	return htc
}

// SetNillableHandlerExecutionID sets the "handler_execution" edge to the HandlerExecution entity by ID if the given value is not nil.
func (htc *HandlerTaskCreate) SetNillableHandlerExecutionID(id *string) *HandlerTaskCreate {
	if id != nil {
		htc = htc.SetHandlerExecutionID(*id)
	}
	return htc
}

// SetHandlerExecution sets the "handler_execution" edge to the HandlerExecution entity.
func (htc *HandlerTaskCreate) SetHandlerExecution(h *HandlerExecution) *HandlerTaskCreate {
	return htc.SetHandlerExecutionID(h.ID)
}

// Mutation returns the HandlerTaskMutation object of the builder.
func (htc *HandlerTaskCreate) Mutation() *HandlerTaskMutation {
	return htc.mutation
}

// Save creates the HandlerTask in the database.
func (htc *HandlerTaskCreate) Save(ctx context.Context) (*HandlerTask, error) {
	return withHooks(ctx, htc.sqlSave, htc.mutation, htc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (htc *HandlerTaskCreate) SaveX(ctx context.Context) *HandlerTask {
	v, err := htc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (htc *HandlerTaskCreate) Exec(ctx context.Context) error {
	_, err := htc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (htc *HandlerTaskCreate) ExecX(ctx context.Context) {
	if err := htc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (htc *HandlerTaskCreate) check() error {
	if _, ok := htc.mutation.HandlerName(); !ok {
		return &ValidationError{Name: "handler_name", err: errors.New(`ent: missing required field "HandlerTask.handler_name"`)}
	}
	if _, ok := htc.mutation.Payload(); !ok {
		return &ValidationError{Name: "payload", err: errors.New(`ent: missing required field "HandlerTask.payload"`)}
	}
	if _, ok := htc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "HandlerTask.status"`)}
	}
	if v, ok := htc.mutation.Status(); ok {
		if err := handlertask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HandlerTask.status": %w`, err)}
		}
	}
	if _, ok := htc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HandlerTask.created_at"`)}
	}
	return nil
}

func (htc *HandlerTaskCreate) sqlSave(ctx context.Context) (*HandlerTask, error) {
	if err := htc.check(); err != nil {
		return nil, err
	}
	_node, _spec := htc.createSpec()
	if err := sqlgraph.CreateNode(ctx, htc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected HandlerTask.ID type: %T", _spec.ID.Value)
		}
	}
	htc.mutation.id = &_node.ID
	htc.mutation.done = true
	return _node, nil
}

func (htc *HandlerTaskCreate) createSpec() (*HandlerTask, *sqlgraph.CreateSpec) {
	var (
		_node = &HandlerTask{config: htc.config}
		_spec = sqlgraph.NewCreateSpec(handlertask.Table, sqlgraph.NewFieldSpec(handlertask.FieldID, field.TypeString))
	)
	if id, ok := htc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := htc.mutation.HandlerName(); ok {
		_spec.SetField(handlertask.FieldHandlerName, field.TypeString, value)
		_node.HandlerName = value
	}
	if value, ok := htc.mutation.Payload(); ok {
		_spec.SetField(handlertask.FieldPayload, field.TypeBytes, value)
		_node.Payload = value
	}
	if value, ok := htc.mutation.Result(); ok {
		_spec.SetField(handlertask.FieldResult, field.TypeBytes, value)
		_node.Result = value
	}
	if value, ok := htc.mutation.Error(); ok {
		_spec.SetField(handlertask.FieldError, field.TypeBytes, value)
		_node.Error = value
	}
	if value, ok := htc.mutation.Status(); ok {
		_spec.SetField(handlertask.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := htc.mutation.CreatedAt(); ok {
		_spec.SetField(handlertask.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := htc.mutation.CompletedAt(); ok {
		_spec.SetField(handlertask.FieldCompletedAt, field.TypeTime, value)
		_node.CompletedAt = value
	}
	if nodes := htc.mutation.HandlerExecutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   handlertask.HandlerExecutionTable,
			Columns: []string{handlertask.HandlerExecutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.handler_execution_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HandlerTaskCreateBulk is the builder for creating many HandlerTask entities in bulk.
type HandlerTaskCreateBulk struct {
	config
	err      error
	builders []*HandlerTaskCreate
}

// Save creates the HandlerTask entities in the database.
func (htcb *HandlerTaskCreateBulk) Save(ctx context.Context) ([]*HandlerTask, error) {
	if htcb.err != nil {
		return nil, htcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(htcb.builders))
	nodes := make([]*HandlerTask, len(htcb.builders))
	mutators := make([]Mutator, len(htcb.builders))
	for i := range htcb.builders {
		func(i int, root context.Context) {
			builder := htcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HandlerTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, htcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, htcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, htcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (htcb *HandlerTaskCreateBulk) SaveX(ctx context.Context) []*HandlerTask {
	v, err := htcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (htcb *HandlerTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := htcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (htcb *HandlerTaskCreateBulk) ExecX(ctx context.Context) {
	if err := htcb.Exec(ctx); err != nil {
		panic(err)
	}
}
