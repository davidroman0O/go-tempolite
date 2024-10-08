// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
	"github.com/davidroman0O/go-tempolite/ent/taskcontext"
)

// TaskContextDelete is the builder for deleting a TaskContext entity.
type TaskContextDelete struct {
	config
	hooks    []Hook
	mutation *TaskContextMutation
}

// Where appends a list predicates to the TaskContextDelete builder.
func (tcd *TaskContextDelete) Where(ps ...predicate.TaskContext) *TaskContextDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TaskContextDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tcd.sqlExec, tcd.mutation, tcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TaskContextDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TaskContextDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(taskcontext.Table, sqlgraph.NewFieldSpec(taskcontext.FieldID, field.TypeString))
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcd.mutation.done = true
	return affected, err
}

// TaskContextDeleteOne is the builder for deleting a single TaskContext entity.
type TaskContextDeleteOne struct {
	tcd *TaskContextDelete
}

// Where appends a list predicates to the TaskContextDelete builder.
func (tcdo *TaskContextDeleteOne) Where(ps ...predicate.TaskContext) *TaskContextDeleteOne {
	tcdo.tcd.mutation.Where(ps...)
	return tcdo
}

// Exec executes the deletion query.
func (tcdo *TaskContextDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskcontext.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TaskContextDeleteOne) ExecX(ctx context.Context) {
	if err := tcdo.Exec(ctx); err != nil {
		panic(err)
	}
}