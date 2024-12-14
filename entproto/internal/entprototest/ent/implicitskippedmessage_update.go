// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImplicitSkippedMessageUpdate is the builder for updating ImplicitSkippedMessage entities.
type ImplicitSkippedMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ImplicitSkippedMessageMutation
}

// Where appends a list predicates to the ImplicitSkippedMessageUpdate builder.
func (ismu *ImplicitSkippedMessageUpdate) Where(ps ...predicate.ImplicitSkippedMessage) *ImplicitSkippedMessageUpdate {
	ismu.mutation.Where(ps...)
	return ismu
}

// Mutation returns the ImplicitSkippedMessageMutation object of the builder.
func (ismu *ImplicitSkippedMessageUpdate) Mutation() *ImplicitSkippedMessageMutation {
	return ismu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ismu *ImplicitSkippedMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ismu.sqlSave, ismu.mutation, ismu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ismu *ImplicitSkippedMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := ismu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ismu *ImplicitSkippedMessageUpdate) Exec(ctx context.Context) error {
	_, err := ismu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ismu *ImplicitSkippedMessageUpdate) ExecX(ctx context.Context) {
	if err := ismu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ismu *ImplicitSkippedMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(implicitskippedmessage.Table, implicitskippedmessage.Columns, sqlgraph.NewFieldSpec(implicitskippedmessage.FieldID, field.TypeInt))
	if ps := ismu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ismu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{implicitskippedmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ismu.mutation.done = true
	return n, nil
}

// ImplicitSkippedMessageUpdateOne is the builder for updating a single ImplicitSkippedMessage entity.
type ImplicitSkippedMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImplicitSkippedMessageMutation
}

// Mutation returns the ImplicitSkippedMessageMutation object of the builder.
func (ismuo *ImplicitSkippedMessageUpdateOne) Mutation() *ImplicitSkippedMessageMutation {
	return ismuo.mutation
}

// Where appends a list predicates to the ImplicitSkippedMessageUpdate builder.
func (ismuo *ImplicitSkippedMessageUpdateOne) Where(ps ...predicate.ImplicitSkippedMessage) *ImplicitSkippedMessageUpdateOne {
	ismuo.mutation.Where(ps...)
	return ismuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ismuo *ImplicitSkippedMessageUpdateOne) Select(field string, fields ...string) *ImplicitSkippedMessageUpdateOne {
	ismuo.fields = append([]string{field}, fields...)
	return ismuo
}

// Save executes the query and returns the updated ImplicitSkippedMessage entity.
func (ismuo *ImplicitSkippedMessageUpdateOne) Save(ctx context.Context) (*ImplicitSkippedMessage, error) {
	return withHooks(ctx, ismuo.sqlSave, ismuo.mutation, ismuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ismuo *ImplicitSkippedMessageUpdateOne) SaveX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ismuo *ImplicitSkippedMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := ismuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ismuo *ImplicitSkippedMessageUpdateOne) ExecX(ctx context.Context) {
	if err := ismuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ismuo *ImplicitSkippedMessageUpdateOne) sqlSave(ctx context.Context) (_node *ImplicitSkippedMessage, err error) {
	_spec := sqlgraph.NewUpdateSpec(implicitskippedmessage.Table, implicitskippedmessage.Columns, sqlgraph.NewFieldSpec(implicitskippedmessage.FieldID, field.TypeInt))
	id, ok := ismuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ImplicitSkippedMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ismuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.FieldID)
		for _, f := range fields {
			if !implicitskippedmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != implicitskippedmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ismuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &ImplicitSkippedMessage{config: ismuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ismuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{implicitskippedmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ismuo.mutation.done = true
	return _node, nil
}
