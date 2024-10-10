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
	"github.com/davidroman0O/go-tempolite/ent/sideeffectresult"
)

// SideEffectResultUpdate is the builder for updating SideEffectResult entities.
type SideEffectResultUpdate struct {
	config
	hooks    []Hook
	mutation *SideEffectResultMutation
}

// Where appends a list predicates to the SideEffectResultUpdate builder.
func (seru *SideEffectResultUpdate) Where(ps ...predicate.SideEffectResult) *SideEffectResultUpdate {
	seru.mutation.Where(ps...)
	return seru
}

// SetExecutionContextID sets the "execution_context_id" field.
func (seru *SideEffectResultUpdate) SetExecutionContextID(s string) *SideEffectResultUpdate {
	seru.mutation.SetExecutionContextID(s)
	return seru
}

// SetNillableExecutionContextID sets the "execution_context_id" field if the given value is not nil.
func (seru *SideEffectResultUpdate) SetNillableExecutionContextID(s *string) *SideEffectResultUpdate {
	if s != nil {
		seru.SetExecutionContextID(*s)
	}
	return seru
}

// SetName sets the "name" field.
func (seru *SideEffectResultUpdate) SetName(s string) *SideEffectResultUpdate {
	seru.mutation.SetName(s)
	return seru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (seru *SideEffectResultUpdate) SetNillableName(s *string) *SideEffectResultUpdate {
	if s != nil {
		seru.SetName(*s)
	}
	return seru
}

// SetResult sets the "result" field.
func (seru *SideEffectResultUpdate) SetResult(b []byte) *SideEffectResultUpdate {
	seru.mutation.SetResult(b)
	return seru
}

// Mutation returns the SideEffectResultMutation object of the builder.
func (seru *SideEffectResultUpdate) Mutation() *SideEffectResultMutation {
	return seru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seru *SideEffectResultUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, seru.sqlSave, seru.mutation, seru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seru *SideEffectResultUpdate) SaveX(ctx context.Context) int {
	affected, err := seru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seru *SideEffectResultUpdate) Exec(ctx context.Context) error {
	_, err := seru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seru *SideEffectResultUpdate) ExecX(ctx context.Context) {
	if err := seru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seru *SideEffectResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sideeffectresult.Table, sideeffectresult.Columns, sqlgraph.NewFieldSpec(sideeffectresult.FieldID, field.TypeString))
	if ps := seru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seru.mutation.ExecutionContextID(); ok {
		_spec.SetField(sideeffectresult.FieldExecutionContextID, field.TypeString, value)
	}
	if value, ok := seru.mutation.Name(); ok {
		_spec.SetField(sideeffectresult.FieldName, field.TypeString, value)
	}
	if value, ok := seru.mutation.Result(); ok {
		_spec.SetField(sideeffectresult.FieldResult, field.TypeBytes, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sideeffectresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	seru.mutation.done = true
	return n, nil
}

// SideEffectResultUpdateOne is the builder for updating a single SideEffectResult entity.
type SideEffectResultUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SideEffectResultMutation
}

// SetExecutionContextID sets the "execution_context_id" field.
func (seruo *SideEffectResultUpdateOne) SetExecutionContextID(s string) *SideEffectResultUpdateOne {
	seruo.mutation.SetExecutionContextID(s)
	return seruo
}

// SetNillableExecutionContextID sets the "execution_context_id" field if the given value is not nil.
func (seruo *SideEffectResultUpdateOne) SetNillableExecutionContextID(s *string) *SideEffectResultUpdateOne {
	if s != nil {
		seruo.SetExecutionContextID(*s)
	}
	return seruo
}

// SetName sets the "name" field.
func (seruo *SideEffectResultUpdateOne) SetName(s string) *SideEffectResultUpdateOne {
	seruo.mutation.SetName(s)
	return seruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (seruo *SideEffectResultUpdateOne) SetNillableName(s *string) *SideEffectResultUpdateOne {
	if s != nil {
		seruo.SetName(*s)
	}
	return seruo
}

// SetResult sets the "result" field.
func (seruo *SideEffectResultUpdateOne) SetResult(b []byte) *SideEffectResultUpdateOne {
	seruo.mutation.SetResult(b)
	return seruo
}

// Mutation returns the SideEffectResultMutation object of the builder.
func (seruo *SideEffectResultUpdateOne) Mutation() *SideEffectResultMutation {
	return seruo.mutation
}

// Where appends a list predicates to the SideEffectResultUpdate builder.
func (seruo *SideEffectResultUpdateOne) Where(ps ...predicate.SideEffectResult) *SideEffectResultUpdateOne {
	seruo.mutation.Where(ps...)
	return seruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seruo *SideEffectResultUpdateOne) Select(field string, fields ...string) *SideEffectResultUpdateOne {
	seruo.fields = append([]string{field}, fields...)
	return seruo
}

// Save executes the query and returns the updated SideEffectResult entity.
func (seruo *SideEffectResultUpdateOne) Save(ctx context.Context) (*SideEffectResult, error) {
	return withHooks(ctx, seruo.sqlSave, seruo.mutation, seruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seruo *SideEffectResultUpdateOne) SaveX(ctx context.Context) *SideEffectResult {
	node, err := seruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seruo *SideEffectResultUpdateOne) Exec(ctx context.Context) error {
	_, err := seruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seruo *SideEffectResultUpdateOne) ExecX(ctx context.Context) {
	if err := seruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seruo *SideEffectResultUpdateOne) sqlSave(ctx context.Context) (_node *SideEffectResult, err error) {
	_spec := sqlgraph.NewUpdateSpec(sideeffectresult.Table, sideeffectresult.Columns, sqlgraph.NewFieldSpec(sideeffectresult.FieldID, field.TypeString))
	id, ok := seruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SideEffectResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sideeffectresult.FieldID)
		for _, f := range fields {
			if !sideeffectresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sideeffectresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seruo.mutation.ExecutionContextID(); ok {
		_spec.SetField(sideeffectresult.FieldExecutionContextID, field.TypeString, value)
	}
	if value, ok := seruo.mutation.Name(); ok {
		_spec.SetField(sideeffectresult.FieldName, field.TypeString, value)
	}
	if value, ok := seruo.mutation.Result(); ok {
		_spec.SetField(sideeffectresult.FieldResult, field.TypeBytes, value)
	}
	_node = &SideEffectResult{config: seruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sideeffectresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	seruo.mutation.done = true
	return _node, nil
}
