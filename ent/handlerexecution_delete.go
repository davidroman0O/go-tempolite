// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/handlerexecution"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
)

// HandlerExecutionDelete is the builder for deleting a HandlerExecution entity.
type HandlerExecutionDelete struct {
	config
	hooks    []Hook
	mutation *HandlerExecutionMutation
}

// Where appends a list predicates to the HandlerExecutionDelete builder.
func (hed *HandlerExecutionDelete) Where(ps ...predicate.HandlerExecution) *HandlerExecutionDelete {
	hed.mutation.Where(ps...)
	return hed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hed *HandlerExecutionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hed.sqlExec, hed.mutation, hed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hed *HandlerExecutionDelete) ExecX(ctx context.Context) int {
	n, err := hed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hed *HandlerExecutionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(handlerexecution.Table, sqlgraph.NewFieldSpec(handlerexecution.FieldID, field.TypeString))
	if ps := hed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hed.mutation.done = true
	return affected, err
}

// HandlerExecutionDeleteOne is the builder for deleting a single HandlerExecution entity.
type HandlerExecutionDeleteOne struct {
	hed *HandlerExecutionDelete
}

// Where appends a list predicates to the HandlerExecutionDelete builder.
func (hedo *HandlerExecutionDeleteOne) Where(ps ...predicate.HandlerExecution) *HandlerExecutionDeleteOne {
	hedo.hed.mutation.Where(ps...)
	return hedo
}

// Exec executes the deletion query.
func (hedo *HandlerExecutionDeleteOne) Exec(ctx context.Context) error {
	n, err := hedo.hed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{handlerexecution.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hedo *HandlerExecutionDeleteOne) ExecX(ctx context.Context) {
	if err := hedo.Exec(ctx); err != nil {
		panic(err)
	}
}
