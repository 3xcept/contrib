// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/messagewithstrings"
	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithStringsQuery is the builder for querying MessageWithStrings entities.
type MessageWithStringsQuery struct {
	config
	ctx        *QueryContext
	order      []messagewithstrings.OrderOption
	inters     []Interceptor
	predicates []predicate.MessageWithStrings
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageWithStringsQuery builder.
func (mwsq *MessageWithStringsQuery) Where(ps ...predicate.MessageWithStrings) *MessageWithStringsQuery {
	mwsq.predicates = append(mwsq.predicates, ps...)
	return mwsq
}

// Limit the number of records to be returned by this query.
func (mwsq *MessageWithStringsQuery) Limit(limit int) *MessageWithStringsQuery {
	mwsq.ctx.Limit = &limit
	return mwsq
}

// Offset to start from.
func (mwsq *MessageWithStringsQuery) Offset(offset int) *MessageWithStringsQuery {
	mwsq.ctx.Offset = &offset
	return mwsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mwsq *MessageWithStringsQuery) Unique(unique bool) *MessageWithStringsQuery {
	mwsq.ctx.Unique = &unique
	return mwsq
}

// Order specifies how the records should be ordered.
func (mwsq *MessageWithStringsQuery) Order(o ...messagewithstrings.OrderOption) *MessageWithStringsQuery {
	mwsq.order = append(mwsq.order, o...)
	return mwsq
}

// First returns the first MessageWithStrings entity from the query.
// Returns a *NotFoundError when no MessageWithStrings was found.
func (mwsq *MessageWithStringsQuery) First(ctx context.Context) (*MessageWithStrings, error) {
	nodes, err := mwsq.Limit(1).All(setContextOp(ctx, mwsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messagewithstrings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) FirstX(ctx context.Context) *MessageWithStrings {
	node, err := mwsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageWithStrings ID from the query.
// Returns a *NotFoundError when no MessageWithStrings ID was found.
func (mwsq *MessageWithStringsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwsq.Limit(1).IDs(setContextOp(ctx, mwsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messagewithstrings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) FirstIDX(ctx context.Context) int {
	id, err := mwsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageWithStrings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageWithStrings entity is found.
// Returns a *NotFoundError when no MessageWithStrings entities are found.
func (mwsq *MessageWithStringsQuery) Only(ctx context.Context) (*MessageWithStrings, error) {
	nodes, err := mwsq.Limit(2).All(setContextOp(ctx, mwsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messagewithstrings.Label}
	default:
		return nil, &NotSingularError{messagewithstrings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) OnlyX(ctx context.Context) *MessageWithStrings {
	node, err := mwsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageWithStrings ID in the query.
// Returns a *NotSingularError when more than one MessageWithStrings ID is found.
// Returns a *NotFoundError when no entities are found.
func (mwsq *MessageWithStringsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwsq.Limit(2).IDs(setContextOp(ctx, mwsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messagewithstrings.Label}
	default:
		err = &NotSingularError{messagewithstrings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) OnlyIDX(ctx context.Context) int {
	id, err := mwsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageWithStringsSlice.
func (mwsq *MessageWithStringsQuery) All(ctx context.Context) ([]*MessageWithStrings, error) {
	ctx = setContextOp(ctx, mwsq.ctx, ent.OpQueryAll)
	if err := mwsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MessageWithStrings, *MessageWithStringsQuery]()
	return withInterceptors[[]*MessageWithStrings](ctx, mwsq, qr, mwsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) AllX(ctx context.Context) []*MessageWithStrings {
	nodes, err := mwsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageWithStrings IDs.
func (mwsq *MessageWithStringsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mwsq.ctx.Unique == nil && mwsq.path != nil {
		mwsq.Unique(true)
	}
	ctx = setContextOp(ctx, mwsq.ctx, ent.OpQueryIDs)
	if err = mwsq.Select(messagewithstrings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) IDsX(ctx context.Context) []int {
	ids, err := mwsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mwsq *MessageWithStringsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mwsq.ctx, ent.OpQueryCount)
	if err := mwsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mwsq, querierCount[*MessageWithStringsQuery](), mwsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) CountX(ctx context.Context) int {
	count, err := mwsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mwsq *MessageWithStringsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mwsq.ctx, ent.OpQueryExist)
	switch _, err := mwsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mwsq *MessageWithStringsQuery) ExistX(ctx context.Context) bool {
	exist, err := mwsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageWithStringsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mwsq *MessageWithStringsQuery) Clone() *MessageWithStringsQuery {
	if mwsq == nil {
		return nil
	}
	return &MessageWithStringsQuery{
		config:     mwsq.config,
		ctx:        mwsq.ctx.Clone(),
		order:      append([]messagewithstrings.OrderOption{}, mwsq.order...),
		inters:     append([]Interceptor{}, mwsq.inters...),
		predicates: append([]predicate.MessageWithStrings{}, mwsq.predicates...),
		// clone intermediate query.
		sql:  mwsq.sql.Clone(),
		path: mwsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Strings []string `json:"strings,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageWithStrings.Query().
//		GroupBy(messagewithstrings.FieldStrings).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mwsq *MessageWithStringsQuery) GroupBy(field string, fields ...string) *MessageWithStringsGroupBy {
	mwsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MessageWithStringsGroupBy{build: mwsq}
	grbuild.flds = &mwsq.ctx.Fields
	grbuild.label = messagewithstrings.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Strings []string `json:"strings,omitempty"`
//	}
//
//	client.MessageWithStrings.Query().
//		Select(messagewithstrings.FieldStrings).
//		Scan(ctx, &v)
func (mwsq *MessageWithStringsQuery) Select(fields ...string) *MessageWithStringsSelect {
	mwsq.ctx.Fields = append(mwsq.ctx.Fields, fields...)
	sbuild := &MessageWithStringsSelect{MessageWithStringsQuery: mwsq}
	sbuild.label = messagewithstrings.Label
	sbuild.flds, sbuild.scan = &mwsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MessageWithStringsSelect configured with the given aggregations.
func (mwsq *MessageWithStringsQuery) Aggregate(fns ...AggregateFunc) *MessageWithStringsSelect {
	return mwsq.Select().Aggregate(fns...)
}

func (mwsq *MessageWithStringsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mwsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mwsq); err != nil {
				return err
			}
		}
	}
	for _, f := range mwsq.ctx.Fields {
		if !messagewithstrings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mwsq.path != nil {
		prev, err := mwsq.path(ctx)
		if err != nil {
			return err
		}
		mwsq.sql = prev
	}
	return nil
}

func (mwsq *MessageWithStringsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MessageWithStrings, error) {
	var (
		nodes = []*MessageWithStrings{}
		_spec = mwsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MessageWithStrings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MessageWithStrings{config: mwsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mwsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mwsq *MessageWithStringsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mwsq.querySpec()
	_spec.Node.Columns = mwsq.ctx.Fields
	if len(mwsq.ctx.Fields) > 0 {
		_spec.Unique = mwsq.ctx.Unique != nil && *mwsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mwsq.driver, _spec)
}

func (mwsq *MessageWithStringsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(messagewithstrings.Table, messagewithstrings.Columns, sqlgraph.NewFieldSpec(messagewithstrings.FieldID, field.TypeInt))
	_spec.From = mwsq.sql
	if unique := mwsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mwsq.path != nil {
		_spec.Unique = true
	}
	if fields := mwsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithstrings.FieldID)
		for i := range fields {
			if fields[i] != messagewithstrings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mwsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mwsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mwsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mwsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mwsq *MessageWithStringsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mwsq.driver.Dialect())
	t1 := builder.Table(messagewithstrings.Table)
	columns := mwsq.ctx.Fields
	if len(columns) == 0 {
		columns = messagewithstrings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mwsq.sql != nil {
		selector = mwsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mwsq.ctx.Unique != nil && *mwsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mwsq.predicates {
		p(selector)
	}
	for _, p := range mwsq.order {
		p(selector)
	}
	if offset := mwsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mwsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageWithStringsGroupBy is the group-by builder for MessageWithStrings entities.
type MessageWithStringsGroupBy struct {
	selector
	build *MessageWithStringsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mwsgb *MessageWithStringsGroupBy) Aggregate(fns ...AggregateFunc) *MessageWithStringsGroupBy {
	mwsgb.fns = append(mwsgb.fns, fns...)
	return mwsgb
}

// Scan applies the selector query and scans the result into the given value.
func (mwsgb *MessageWithStringsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mwsgb.build.ctx, ent.OpQueryGroupBy)
	if err := mwsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageWithStringsQuery, *MessageWithStringsGroupBy](ctx, mwsgb.build, mwsgb, mwsgb.build.inters, v)
}

func (mwsgb *MessageWithStringsGroupBy) sqlScan(ctx context.Context, root *MessageWithStringsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mwsgb.fns))
	for _, fn := range mwsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mwsgb.flds)+len(mwsgb.fns))
		for _, f := range *mwsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mwsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MessageWithStringsSelect is the builder for selecting fields of MessageWithStrings entities.
type MessageWithStringsSelect struct {
	*MessageWithStringsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mwss *MessageWithStringsSelect) Aggregate(fns ...AggregateFunc) *MessageWithStringsSelect {
	mwss.fns = append(mwss.fns, fns...)
	return mwss
}

// Scan applies the selector query and scans the result into the given value.
func (mwss *MessageWithStringsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mwss.ctx, ent.OpQuerySelect)
	if err := mwss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageWithStringsQuery, *MessageWithStringsSelect](ctx, mwss.MessageWithStringsQuery, mwss, mwss.inters, v)
}

func (mwss *MessageWithStringsSelect) sqlScan(ctx context.Context, root *MessageWithStringsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mwss.fns))
	for _, fn := range mwss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mwss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
