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
	"github.com/davidroman0O/tempolite/ent/sideeffect"
	"github.com/davidroman0O/tempolite/ent/sideeffectexecution"
)

// SideEffectExecutionUpdate is the builder for updating SideEffectExecution entities.
type SideEffectExecutionUpdate struct {
	config
	hooks    []Hook
	mutation *SideEffectExecutionMutation
}

// Where appends a list predicates to the SideEffectExecutionUpdate builder.
func (seeu *SideEffectExecutionUpdate) Where(ps ...predicate.SideEffectExecution) *SideEffectExecutionUpdate {
	seeu.mutation.Where(ps...)
	return seeu
}

// SetStatus sets the "status" field.
func (seeu *SideEffectExecutionUpdate) SetStatus(s sideeffectexecution.Status) *SideEffectExecutionUpdate {
	seeu.mutation.SetStatus(s)
	return seeu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (seeu *SideEffectExecutionUpdate) SetNillableStatus(s *sideeffectexecution.Status) *SideEffectExecutionUpdate {
	if s != nil {
		seeu.SetStatus(*s)
	}
	return seeu
}

// SetAttempt sets the "attempt" field.
func (seeu *SideEffectExecutionUpdate) SetAttempt(i int) *SideEffectExecutionUpdate {
	seeu.mutation.ResetAttempt()
	seeu.mutation.SetAttempt(i)
	return seeu
}

// SetNillableAttempt sets the "attempt" field if the given value is not nil.
func (seeu *SideEffectExecutionUpdate) SetNillableAttempt(i *int) *SideEffectExecutionUpdate {
	if i != nil {
		seeu.SetAttempt(*i)
	}
	return seeu
}

// AddAttempt adds i to the "attempt" field.
func (seeu *SideEffectExecutionUpdate) AddAttempt(i int) *SideEffectExecutionUpdate {
	seeu.mutation.AddAttempt(i)
	return seeu
}

// SetOutput sets the "output" field.
func (seeu *SideEffectExecutionUpdate) SetOutput(i []interface{}) *SideEffectExecutionUpdate {
	seeu.mutation.SetOutput(i)
	return seeu
}

// AppendOutput appends i to the "output" field.
func (seeu *SideEffectExecutionUpdate) AppendOutput(i []interface{}) *SideEffectExecutionUpdate {
	seeu.mutation.AppendOutput(i)
	return seeu
}

// ClearOutput clears the value of the "output" field.
func (seeu *SideEffectExecutionUpdate) ClearOutput() *SideEffectExecutionUpdate {
	seeu.mutation.ClearOutput()
	return seeu
}

// SetError sets the "error" field.
func (seeu *SideEffectExecutionUpdate) SetError(s string) *SideEffectExecutionUpdate {
	seeu.mutation.SetError(s)
	return seeu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (seeu *SideEffectExecutionUpdate) SetNillableError(s *string) *SideEffectExecutionUpdate {
	if s != nil {
		seeu.SetError(*s)
	}
	return seeu
}

// ClearError clears the value of the "error" field.
func (seeu *SideEffectExecutionUpdate) ClearError() *SideEffectExecutionUpdate {
	seeu.mutation.ClearError()
	return seeu
}

// SetStartedAt sets the "started_at" field.
func (seeu *SideEffectExecutionUpdate) SetStartedAt(t time.Time) *SideEffectExecutionUpdate {
	seeu.mutation.SetStartedAt(t)
	return seeu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (seeu *SideEffectExecutionUpdate) SetNillableStartedAt(t *time.Time) *SideEffectExecutionUpdate {
	if t != nil {
		seeu.SetStartedAt(*t)
	}
	return seeu
}

// SetUpdatedAt sets the "updated_at" field.
func (seeu *SideEffectExecutionUpdate) SetUpdatedAt(t time.Time) *SideEffectExecutionUpdate {
	seeu.mutation.SetUpdatedAt(t)
	return seeu
}

// SetSideEffectID sets the "side_effect" edge to the SideEffect entity by ID.
func (seeu *SideEffectExecutionUpdate) SetSideEffectID(id string) *SideEffectExecutionUpdate {
	seeu.mutation.SetSideEffectID(id)
	return seeu
}

// SetSideEffect sets the "side_effect" edge to the SideEffect entity.
func (seeu *SideEffectExecutionUpdate) SetSideEffect(s *SideEffect) *SideEffectExecutionUpdate {
	return seeu.SetSideEffectID(s.ID)
}

// Mutation returns the SideEffectExecutionMutation object of the builder.
func (seeu *SideEffectExecutionUpdate) Mutation() *SideEffectExecutionMutation {
	return seeu.mutation
}

// ClearSideEffect clears the "side_effect" edge to the SideEffect entity.
func (seeu *SideEffectExecutionUpdate) ClearSideEffect() *SideEffectExecutionUpdate {
	seeu.mutation.ClearSideEffect()
	return seeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seeu *SideEffectExecutionUpdate) Save(ctx context.Context) (int, error) {
	seeu.defaults()
	return withHooks(ctx, seeu.sqlSave, seeu.mutation, seeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seeu *SideEffectExecutionUpdate) SaveX(ctx context.Context) int {
	affected, err := seeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seeu *SideEffectExecutionUpdate) Exec(ctx context.Context) error {
	_, err := seeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seeu *SideEffectExecutionUpdate) ExecX(ctx context.Context) {
	if err := seeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (seeu *SideEffectExecutionUpdate) defaults() {
	if _, ok := seeu.mutation.UpdatedAt(); !ok {
		v := sideeffectexecution.UpdateDefaultUpdatedAt()
		seeu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (seeu *SideEffectExecutionUpdate) check() error {
	if v, ok := seeu.mutation.Status(); ok {
		if err := sideeffectexecution.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SideEffectExecution.status": %w`, err)}
		}
	}
	if seeu.mutation.SideEffectCleared() && len(seeu.mutation.SideEffectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "SideEffectExecution.side_effect"`)
	}
	return nil
}

func (seeu *SideEffectExecutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := seeu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sideeffectexecution.Table, sideeffectexecution.Columns, sqlgraph.NewFieldSpec(sideeffectexecution.FieldID, field.TypeString))
	if ps := seeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seeu.mutation.Status(); ok {
		_spec.SetField(sideeffectexecution.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := seeu.mutation.Attempt(); ok {
		_spec.SetField(sideeffectexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := seeu.mutation.AddedAttempt(); ok {
		_spec.AddField(sideeffectexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := seeu.mutation.Output(); ok {
		_spec.SetField(sideeffectexecution.FieldOutput, field.TypeJSON, value)
	}
	if value, ok := seeu.mutation.AppendedOutput(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, sideeffectexecution.FieldOutput, value)
		})
	}
	if seeu.mutation.OutputCleared() {
		_spec.ClearField(sideeffectexecution.FieldOutput, field.TypeJSON)
	}
	if value, ok := seeu.mutation.Error(); ok {
		_spec.SetField(sideeffectexecution.FieldError, field.TypeString, value)
	}
	if seeu.mutation.ErrorCleared() {
		_spec.ClearField(sideeffectexecution.FieldError, field.TypeString)
	}
	if value, ok := seeu.mutation.StartedAt(); ok {
		_spec.SetField(sideeffectexecution.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := seeu.mutation.UpdatedAt(); ok {
		_spec.SetField(sideeffectexecution.FieldUpdatedAt, field.TypeTime, value)
	}
	if seeu.mutation.SideEffectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sideeffectexecution.SideEffectTable,
			Columns: []string{sideeffectexecution.SideEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sideeffect.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seeu.mutation.SideEffectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sideeffectexecution.SideEffectTable,
			Columns: []string{sideeffectexecution.SideEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sideeffect.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sideeffectexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	seeu.mutation.done = true
	return n, nil
}

// SideEffectExecutionUpdateOne is the builder for updating a single SideEffectExecution entity.
type SideEffectExecutionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SideEffectExecutionMutation
}

// SetStatus sets the "status" field.
func (seeuo *SideEffectExecutionUpdateOne) SetStatus(s sideeffectexecution.Status) *SideEffectExecutionUpdateOne {
	seeuo.mutation.SetStatus(s)
	return seeuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (seeuo *SideEffectExecutionUpdateOne) SetNillableStatus(s *sideeffectexecution.Status) *SideEffectExecutionUpdateOne {
	if s != nil {
		seeuo.SetStatus(*s)
	}
	return seeuo
}

// SetAttempt sets the "attempt" field.
func (seeuo *SideEffectExecutionUpdateOne) SetAttempt(i int) *SideEffectExecutionUpdateOne {
	seeuo.mutation.ResetAttempt()
	seeuo.mutation.SetAttempt(i)
	return seeuo
}

// SetNillableAttempt sets the "attempt" field if the given value is not nil.
func (seeuo *SideEffectExecutionUpdateOne) SetNillableAttempt(i *int) *SideEffectExecutionUpdateOne {
	if i != nil {
		seeuo.SetAttempt(*i)
	}
	return seeuo
}

// AddAttempt adds i to the "attempt" field.
func (seeuo *SideEffectExecutionUpdateOne) AddAttempt(i int) *SideEffectExecutionUpdateOne {
	seeuo.mutation.AddAttempt(i)
	return seeuo
}

// SetOutput sets the "output" field.
func (seeuo *SideEffectExecutionUpdateOne) SetOutput(i []interface{}) *SideEffectExecutionUpdateOne {
	seeuo.mutation.SetOutput(i)
	return seeuo
}

// AppendOutput appends i to the "output" field.
func (seeuo *SideEffectExecutionUpdateOne) AppendOutput(i []interface{}) *SideEffectExecutionUpdateOne {
	seeuo.mutation.AppendOutput(i)
	return seeuo
}

// ClearOutput clears the value of the "output" field.
func (seeuo *SideEffectExecutionUpdateOne) ClearOutput() *SideEffectExecutionUpdateOne {
	seeuo.mutation.ClearOutput()
	return seeuo
}

// SetError sets the "error" field.
func (seeuo *SideEffectExecutionUpdateOne) SetError(s string) *SideEffectExecutionUpdateOne {
	seeuo.mutation.SetError(s)
	return seeuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (seeuo *SideEffectExecutionUpdateOne) SetNillableError(s *string) *SideEffectExecutionUpdateOne {
	if s != nil {
		seeuo.SetError(*s)
	}
	return seeuo
}

// ClearError clears the value of the "error" field.
func (seeuo *SideEffectExecutionUpdateOne) ClearError() *SideEffectExecutionUpdateOne {
	seeuo.mutation.ClearError()
	return seeuo
}

// SetStartedAt sets the "started_at" field.
func (seeuo *SideEffectExecutionUpdateOne) SetStartedAt(t time.Time) *SideEffectExecutionUpdateOne {
	seeuo.mutation.SetStartedAt(t)
	return seeuo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (seeuo *SideEffectExecutionUpdateOne) SetNillableStartedAt(t *time.Time) *SideEffectExecutionUpdateOne {
	if t != nil {
		seeuo.SetStartedAt(*t)
	}
	return seeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (seeuo *SideEffectExecutionUpdateOne) SetUpdatedAt(t time.Time) *SideEffectExecutionUpdateOne {
	seeuo.mutation.SetUpdatedAt(t)
	return seeuo
}

// SetSideEffectID sets the "side_effect" edge to the SideEffect entity by ID.
func (seeuo *SideEffectExecutionUpdateOne) SetSideEffectID(id string) *SideEffectExecutionUpdateOne {
	seeuo.mutation.SetSideEffectID(id)
	return seeuo
}

// SetSideEffect sets the "side_effect" edge to the SideEffect entity.
func (seeuo *SideEffectExecutionUpdateOne) SetSideEffect(s *SideEffect) *SideEffectExecutionUpdateOne {
	return seeuo.SetSideEffectID(s.ID)
}

// Mutation returns the SideEffectExecutionMutation object of the builder.
func (seeuo *SideEffectExecutionUpdateOne) Mutation() *SideEffectExecutionMutation {
	return seeuo.mutation
}

// ClearSideEffect clears the "side_effect" edge to the SideEffect entity.
func (seeuo *SideEffectExecutionUpdateOne) ClearSideEffect() *SideEffectExecutionUpdateOne {
	seeuo.mutation.ClearSideEffect()
	return seeuo
}

// Where appends a list predicates to the SideEffectExecutionUpdate builder.
func (seeuo *SideEffectExecutionUpdateOne) Where(ps ...predicate.SideEffectExecution) *SideEffectExecutionUpdateOne {
	seeuo.mutation.Where(ps...)
	return seeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seeuo *SideEffectExecutionUpdateOne) Select(field string, fields ...string) *SideEffectExecutionUpdateOne {
	seeuo.fields = append([]string{field}, fields...)
	return seeuo
}

// Save executes the query and returns the updated SideEffectExecution entity.
func (seeuo *SideEffectExecutionUpdateOne) Save(ctx context.Context) (*SideEffectExecution, error) {
	seeuo.defaults()
	return withHooks(ctx, seeuo.sqlSave, seeuo.mutation, seeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seeuo *SideEffectExecutionUpdateOne) SaveX(ctx context.Context) *SideEffectExecution {
	node, err := seeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seeuo *SideEffectExecutionUpdateOne) Exec(ctx context.Context) error {
	_, err := seeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seeuo *SideEffectExecutionUpdateOne) ExecX(ctx context.Context) {
	if err := seeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (seeuo *SideEffectExecutionUpdateOne) defaults() {
	if _, ok := seeuo.mutation.UpdatedAt(); !ok {
		v := sideeffectexecution.UpdateDefaultUpdatedAt()
		seeuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (seeuo *SideEffectExecutionUpdateOne) check() error {
	if v, ok := seeuo.mutation.Status(); ok {
		if err := sideeffectexecution.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SideEffectExecution.status": %w`, err)}
		}
	}
	if seeuo.mutation.SideEffectCleared() && len(seeuo.mutation.SideEffectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "SideEffectExecution.side_effect"`)
	}
	return nil
}

func (seeuo *SideEffectExecutionUpdateOne) sqlSave(ctx context.Context) (_node *SideEffectExecution, err error) {
	if err := seeuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sideeffectexecution.Table, sideeffectexecution.Columns, sqlgraph.NewFieldSpec(sideeffectexecution.FieldID, field.TypeString))
	id, ok := seeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SideEffectExecution.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sideeffectexecution.FieldID)
		for _, f := range fields {
			if !sideeffectexecution.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sideeffectexecution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seeuo.mutation.Status(); ok {
		_spec.SetField(sideeffectexecution.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := seeuo.mutation.Attempt(); ok {
		_spec.SetField(sideeffectexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := seeuo.mutation.AddedAttempt(); ok {
		_spec.AddField(sideeffectexecution.FieldAttempt, field.TypeInt, value)
	}
	if value, ok := seeuo.mutation.Output(); ok {
		_spec.SetField(sideeffectexecution.FieldOutput, field.TypeJSON, value)
	}
	if value, ok := seeuo.mutation.AppendedOutput(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, sideeffectexecution.FieldOutput, value)
		})
	}
	if seeuo.mutation.OutputCleared() {
		_spec.ClearField(sideeffectexecution.FieldOutput, field.TypeJSON)
	}
	if value, ok := seeuo.mutation.Error(); ok {
		_spec.SetField(sideeffectexecution.FieldError, field.TypeString, value)
	}
	if seeuo.mutation.ErrorCleared() {
		_spec.ClearField(sideeffectexecution.FieldError, field.TypeString)
	}
	if value, ok := seeuo.mutation.StartedAt(); ok {
		_spec.SetField(sideeffectexecution.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := seeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sideeffectexecution.FieldUpdatedAt, field.TypeTime, value)
	}
	if seeuo.mutation.SideEffectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sideeffectexecution.SideEffectTable,
			Columns: []string{sideeffectexecution.SideEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sideeffect.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seeuo.mutation.SideEffectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sideeffectexecution.SideEffectTable,
			Columns: []string{sideeffectexecution.SideEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sideeffect.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SideEffectExecution{config: seeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sideeffectexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	seeuo.mutation.done = true
	return _node, nil
}
