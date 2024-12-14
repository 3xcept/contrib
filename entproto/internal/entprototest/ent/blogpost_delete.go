// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/blogpost"
	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogPostDelete is the builder for deleting a BlogPost entity.
type BlogPostDelete struct {
	config
	hooks    []Hook
	mutation *BlogPostMutation
}

// Where appends a list predicates to the BlogPostDelete builder.
func (bpd *BlogPostDelete) Where(ps ...predicate.BlogPost) *BlogPostDelete {
	bpd.mutation.Where(ps...)
	return bpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bpd *BlogPostDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bpd.sqlExec, bpd.mutation, bpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bpd *BlogPostDelete) ExecX(ctx context.Context) int {
	n, err := bpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bpd *BlogPostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(blogpost.Table, sqlgraph.NewFieldSpec(blogpost.FieldID, field.TypeInt))
	if ps := bpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bpd.mutation.done = true
	return affected, err
}

// BlogPostDeleteOne is the builder for deleting a single BlogPost entity.
type BlogPostDeleteOne struct {
	bpd *BlogPostDelete
}

// Where appends a list predicates to the BlogPostDelete builder.
func (bpdo *BlogPostDeleteOne) Where(ps ...predicate.BlogPost) *BlogPostDeleteOne {
	bpdo.bpd.mutation.Where(ps...)
	return bpdo
}

// Exec executes the deletion query.
func (bpdo *BlogPostDeleteOne) Exec(ctx context.Context) error {
	n, err := bpdo.bpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{blogpost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bpdo *BlogPostDeleteOne) ExecX(ctx context.Context) {
	if err := bpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
