// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/sagatask"
)

// SagaTaskCreate is the builder for creating a SagaTask entity.
type SagaTaskCreate struct {
	config
	mutation *SagaTaskMutation
	hooks    []Hook
}

// Mutation returns the SagaTaskMutation object of the builder.
func (stc *SagaTaskCreate) Mutation() *SagaTaskMutation {
	return stc.mutation
}

// Save creates the SagaTask in the database.
func (stc *SagaTaskCreate) Save(ctx context.Context) (*SagaTask, error) {
	return withHooks(ctx, stc.sqlSave, stc.mutation, stc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (stc *SagaTaskCreate) SaveX(ctx context.Context) *SagaTask {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *SagaTaskCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *SagaTaskCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *SagaTaskCreate) check() error {
	return nil
}

func (stc *SagaTaskCreate) sqlSave(ctx context.Context) (*SagaTask, error) {
	if err := stc.check(); err != nil {
		return nil, err
	}
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	stc.mutation.id = &_node.ID
	stc.mutation.done = true
	return _node, nil
}

func (stc *SagaTaskCreate) createSpec() (*SagaTask, *sqlgraph.CreateSpec) {
	var (
		_node = &SagaTask{config: stc.config}
		_spec = sqlgraph.NewCreateSpec(sagatask.Table, sqlgraph.NewFieldSpec(sagatask.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// SagaTaskCreateBulk is the builder for creating many SagaTask entities in bulk.
type SagaTaskCreateBulk struct {
	config
	err      error
	builders []*SagaTaskCreate
}

// Save creates the SagaTask entities in the database.
func (stcb *SagaTaskCreateBulk) Save(ctx context.Context) ([]*SagaTask, error) {
	if stcb.err != nil {
		return nil, stcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*SagaTask, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SagaTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *SagaTaskCreateBulk) SaveX(ctx context.Context) []*SagaTask {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *SagaTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *SagaTaskCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}
