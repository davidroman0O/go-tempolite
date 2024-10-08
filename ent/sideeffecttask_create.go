// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/node"
	"github.com/davidroman0O/go-tempolite/ent/sideeffecttask"
)

// SideEffectTaskCreate is the builder for creating a SideEffectTask entity.
type SideEffectTaskCreate struct {
	config
	mutation *SideEffectTaskMutation
	hooks    []Hook
}

// SetNodeID sets the "node" edge to the Node entity by ID.
func (setc *SideEffectTaskCreate) SetNodeID(id string) *SideEffectTaskCreate {
	setc.mutation.SetNodeID(id)
	return setc
}

// SetNillableNodeID sets the "node" edge to the Node entity by ID if the given value is not nil.
func (setc *SideEffectTaskCreate) SetNillableNodeID(id *string) *SideEffectTaskCreate {
	if id != nil {
		setc = setc.SetNodeID(*id)
	}
	return setc
}

// SetNode sets the "node" edge to the Node entity.
func (setc *SideEffectTaskCreate) SetNode(n *Node) *SideEffectTaskCreate {
	return setc.SetNodeID(n.ID)
}

// Mutation returns the SideEffectTaskMutation object of the builder.
func (setc *SideEffectTaskCreate) Mutation() *SideEffectTaskMutation {
	return setc.mutation
}

// Save creates the SideEffectTask in the database.
func (setc *SideEffectTaskCreate) Save(ctx context.Context) (*SideEffectTask, error) {
	return withHooks(ctx, setc.sqlSave, setc.mutation, setc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (setc *SideEffectTaskCreate) SaveX(ctx context.Context) *SideEffectTask {
	v, err := setc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (setc *SideEffectTaskCreate) Exec(ctx context.Context) error {
	_, err := setc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (setc *SideEffectTaskCreate) ExecX(ctx context.Context) {
	if err := setc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (setc *SideEffectTaskCreate) check() error {
	return nil
}

func (setc *SideEffectTaskCreate) sqlSave(ctx context.Context) (*SideEffectTask, error) {
	if err := setc.check(); err != nil {
		return nil, err
	}
	_node, _spec := setc.createSpec()
	if err := sqlgraph.CreateNode(ctx, setc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	setc.mutation.id = &_node.ID
	setc.mutation.done = true
	return _node, nil
}

func (setc *SideEffectTaskCreate) createSpec() (*SideEffectTask, *sqlgraph.CreateSpec) {
	var (
		_node = &SideEffectTask{config: setc.config}
		_spec = sqlgraph.NewCreateSpec(sideeffecttask.Table, sqlgraph.NewFieldSpec(sideeffecttask.FieldID, field.TypeInt))
	)
	if nodes := setc.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   sideeffecttask.NodeTable,
			Columns: []string{sideeffecttask.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.node_side_effect_task = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SideEffectTaskCreateBulk is the builder for creating many SideEffectTask entities in bulk.
type SideEffectTaskCreateBulk struct {
	config
	err      error
	builders []*SideEffectTaskCreate
}

// Save creates the SideEffectTask entities in the database.
func (setcb *SideEffectTaskCreateBulk) Save(ctx context.Context) ([]*SideEffectTask, error) {
	if setcb.err != nil {
		return nil, setcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(setcb.builders))
	nodes := make([]*SideEffectTask, len(setcb.builders))
	mutators := make([]Mutator, len(setcb.builders))
	for i := range setcb.builders {
		func(i int, root context.Context) {
			builder := setcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SideEffectTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, setcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, setcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, setcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (setcb *SideEffectTaskCreateBulk) SaveX(ctx context.Context) []*SideEffectTask {
	v, err := setcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (setcb *SideEffectTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := setcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (setcb *SideEffectTaskCreateBulk) ExecX(ctx context.Context) {
	if err := setcb.Exec(ctx); err != nil {
		panic(err)
	}
}