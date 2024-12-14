// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/3xcept/contrib/entgql/internal/todopulid/ent/predicate"
	"github.com/3xcept/contrib/entgql/internal/todopulid/ent/verysecret"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VerySecretDelete is the builder for deleting a VerySecret entity.
type VerySecretDelete struct {
	config
	hooks    []Hook
	mutation *VerySecretMutation
}

// Where appends a list predicates to the VerySecretDelete builder.
func (vsd *VerySecretDelete) Where(ps ...predicate.VerySecret) *VerySecretDelete {
	vsd.mutation.Where(ps...)
	return vsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vsd *VerySecretDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vsd.sqlExec, vsd.mutation, vsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vsd *VerySecretDelete) ExecX(ctx context.Context) int {
	n, err := vsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vsd *VerySecretDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(verysecret.Table, sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeString))
	if ps := vsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vsd.mutation.done = true
	return affected, err
}

// VerySecretDeleteOne is the builder for deleting a single VerySecret entity.
type VerySecretDeleteOne struct {
	vsd *VerySecretDelete
}

// Where appends a list predicates to the VerySecretDelete builder.
func (vsdo *VerySecretDeleteOne) Where(ps ...predicate.VerySecret) *VerySecretDeleteOne {
	vsdo.vsd.mutation.Where(ps...)
	return vsdo
}

// Exec executes the deletion query.
func (vsdo *VerySecretDeleteOne) Exec(ctx context.Context) error {
	n, err := vsdo.vsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{verysecret.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vsdo *VerySecretDeleteOne) ExecX(ctx context.Context) {
	if err := vsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
