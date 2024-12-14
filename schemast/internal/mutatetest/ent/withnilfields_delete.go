// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/3xcept/contrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/3xcept/contrib/schemast/internal/mutatetest/ent/withnilfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithNilFieldsDelete is the builder for deleting a WithNilFields entity.
type WithNilFieldsDelete struct {
	config
	hooks    []Hook
	mutation *WithNilFieldsMutation
}

// Where appends a list predicates to the WithNilFieldsDelete builder.
func (wnfd *WithNilFieldsDelete) Where(ps ...predicate.WithNilFields) *WithNilFieldsDelete {
	wnfd.mutation.Where(ps...)
	return wnfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wnfd *WithNilFieldsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wnfd.sqlExec, wnfd.mutation, wnfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wnfd *WithNilFieldsDelete) ExecX(ctx context.Context) int {
	n, err := wnfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wnfd *WithNilFieldsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(withnilfields.Table, sqlgraph.NewFieldSpec(withnilfields.FieldID, field.TypeInt))
	if ps := wnfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wnfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wnfd.mutation.done = true
	return affected, err
}

// WithNilFieldsDeleteOne is the builder for deleting a single WithNilFields entity.
type WithNilFieldsDeleteOne struct {
	wnfd *WithNilFieldsDelete
}

// Where appends a list predicates to the WithNilFieldsDelete builder.
func (wnfdo *WithNilFieldsDeleteOne) Where(ps ...predicate.WithNilFields) *WithNilFieldsDeleteOne {
	wnfdo.wnfd.mutation.Where(ps...)
	return wnfdo
}

// Exec executes the deletion query.
func (wnfdo *WithNilFieldsDeleteOne) Exec(ctx context.Context) error {
	n, err := wnfdo.wnfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withnilfields.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wnfdo *WithNilFieldsDeleteOne) ExecX(ctx context.Context) {
	if err := wnfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
