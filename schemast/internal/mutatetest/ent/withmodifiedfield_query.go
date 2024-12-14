// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/3xcept/contrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/3xcept/contrib/schemast/internal/mutatetest/ent/user"
	"github.com/3xcept/contrib/schemast/internal/mutatetest/ent/withmodifiedfield"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithModifiedFieldQuery is the builder for querying WithModifiedField entities.
type WithModifiedFieldQuery struct {
	config
	ctx        *QueryContext
	order      []withmodifiedfield.OrderOption
	inters     []Interceptor
	predicates []predicate.WithModifiedField
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithModifiedFieldQuery builder.
func (wmfq *WithModifiedFieldQuery) Where(ps ...predicate.WithModifiedField) *WithModifiedFieldQuery {
	wmfq.predicates = append(wmfq.predicates, ps...)
	return wmfq
}

// Limit the number of records to be returned by this query.
func (wmfq *WithModifiedFieldQuery) Limit(limit int) *WithModifiedFieldQuery {
	wmfq.ctx.Limit = &limit
	return wmfq
}

// Offset to start from.
func (wmfq *WithModifiedFieldQuery) Offset(offset int) *WithModifiedFieldQuery {
	wmfq.ctx.Offset = &offset
	return wmfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wmfq *WithModifiedFieldQuery) Unique(unique bool) *WithModifiedFieldQuery {
	wmfq.ctx.Unique = &unique
	return wmfq
}

// Order specifies how the records should be ordered.
func (wmfq *WithModifiedFieldQuery) Order(o ...withmodifiedfield.OrderOption) *WithModifiedFieldQuery {
	wmfq.order = append(wmfq.order, o...)
	return wmfq
}

// QueryOwner chains the current query on the "owner" edge.
func (wmfq *WithModifiedFieldQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: wmfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wmfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wmfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(withmodifiedfield.Table, withmodifiedfield.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, withmodifiedfield.OwnerTable, withmodifiedfield.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(wmfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WithModifiedField entity from the query.
// Returns a *NotFoundError when no WithModifiedField was found.
func (wmfq *WithModifiedFieldQuery) First(ctx context.Context) (*WithModifiedField, error) {
	nodes, err := wmfq.Limit(1).All(setContextOp(ctx, wmfq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withmodifiedfield.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) FirstX(ctx context.Context) *WithModifiedField {
	node, err := wmfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithModifiedField ID from the query.
// Returns a *NotFoundError when no WithModifiedField ID was found.
func (wmfq *WithModifiedFieldQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wmfq.Limit(1).IDs(setContextOp(ctx, wmfq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withmodifiedfield.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) FirstIDX(ctx context.Context) int {
	id, err := wmfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithModifiedField entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WithModifiedField entity is found.
// Returns a *NotFoundError when no WithModifiedField entities are found.
func (wmfq *WithModifiedFieldQuery) Only(ctx context.Context) (*WithModifiedField, error) {
	nodes, err := wmfq.Limit(2).All(setContextOp(ctx, wmfq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withmodifiedfield.Label}
	default:
		return nil, &NotSingularError{withmodifiedfield.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) OnlyX(ctx context.Context) *WithModifiedField {
	node, err := wmfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithModifiedField ID in the query.
// Returns a *NotSingularError when more than one WithModifiedField ID is found.
// Returns a *NotFoundError when no entities are found.
func (wmfq *WithModifiedFieldQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wmfq.Limit(2).IDs(setContextOp(ctx, wmfq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = &NotSingularError{withmodifiedfield.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) OnlyIDX(ctx context.Context) int {
	id, err := wmfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithModifiedFields.
func (wmfq *WithModifiedFieldQuery) All(ctx context.Context) ([]*WithModifiedField, error) {
	ctx = setContextOp(ctx, wmfq.ctx, ent.OpQueryAll)
	if err := wmfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WithModifiedField, *WithModifiedFieldQuery]()
	return withInterceptors[[]*WithModifiedField](ctx, wmfq, qr, wmfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) AllX(ctx context.Context) []*WithModifiedField {
	nodes, err := wmfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithModifiedField IDs.
func (wmfq *WithModifiedFieldQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wmfq.ctx.Unique == nil && wmfq.path != nil {
		wmfq.Unique(true)
	}
	ctx = setContextOp(ctx, wmfq.ctx, ent.OpQueryIDs)
	if err = wmfq.Select(withmodifiedfield.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) IDsX(ctx context.Context) []int {
	ids, err := wmfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wmfq *WithModifiedFieldQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wmfq.ctx, ent.OpQueryCount)
	if err := wmfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wmfq, querierCount[*WithModifiedFieldQuery](), wmfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) CountX(ctx context.Context) int {
	count, err := wmfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wmfq *WithModifiedFieldQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wmfq.ctx, ent.OpQueryExist)
	switch _, err := wmfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) ExistX(ctx context.Context) bool {
	exist, err := wmfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithModifiedFieldQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wmfq *WithModifiedFieldQuery) Clone() *WithModifiedFieldQuery {
	if wmfq == nil {
		return nil
	}
	return &WithModifiedFieldQuery{
		config:     wmfq.config,
		ctx:        wmfq.ctx.Clone(),
		order:      append([]withmodifiedfield.OrderOption{}, wmfq.order...),
		inters:     append([]Interceptor{}, wmfq.inters...),
		predicates: append([]predicate.WithModifiedField{}, wmfq.predicates...),
		withOwner:  wmfq.withOwner.Clone(),
		// clone intermediate query.
		sql:  wmfq.sql.Clone(),
		path: wmfq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (wmfq *WithModifiedFieldQuery) WithOwner(opts ...func(*UserQuery)) *WithModifiedFieldQuery {
	query := (&UserClient{config: wmfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wmfq.withOwner = query
	return wmfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WithModifiedField.Query().
//		GroupBy(withmodifiedfield.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wmfq *WithModifiedFieldQuery) GroupBy(field string, fields ...string) *WithModifiedFieldGroupBy {
	wmfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WithModifiedFieldGroupBy{build: wmfq}
	grbuild.flds = &wmfq.ctx.Fields
	grbuild.label = withmodifiedfield.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WithModifiedField.Query().
//		Select(withmodifiedfield.FieldName).
//		Scan(ctx, &v)
func (wmfq *WithModifiedFieldQuery) Select(fields ...string) *WithModifiedFieldSelect {
	wmfq.ctx.Fields = append(wmfq.ctx.Fields, fields...)
	sbuild := &WithModifiedFieldSelect{WithModifiedFieldQuery: wmfq}
	sbuild.label = withmodifiedfield.Label
	sbuild.flds, sbuild.scan = &wmfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WithModifiedFieldSelect configured with the given aggregations.
func (wmfq *WithModifiedFieldQuery) Aggregate(fns ...AggregateFunc) *WithModifiedFieldSelect {
	return wmfq.Select().Aggregate(fns...)
}

func (wmfq *WithModifiedFieldQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wmfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wmfq); err != nil {
				return err
			}
		}
	}
	for _, f := range wmfq.ctx.Fields {
		if !withmodifiedfield.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wmfq.path != nil {
		prev, err := wmfq.path(ctx)
		if err != nil {
			return err
		}
		wmfq.sql = prev
	}
	return nil
}

func (wmfq *WithModifiedFieldQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WithModifiedField, error) {
	var (
		nodes       = []*WithModifiedField{}
		withFKs     = wmfq.withFKs
		_spec       = wmfq.querySpec()
		loadedTypes = [1]bool{
			wmfq.withOwner != nil,
		}
	)
	if wmfq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, withmodifiedfield.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WithModifiedField).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WithModifiedField{config: wmfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wmfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wmfq.withOwner; query != nil {
		if err := wmfq.loadOwner(ctx, query, nodes, nil,
			func(n *WithModifiedField, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wmfq *WithModifiedFieldQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*WithModifiedField, init func(*WithModifiedField), assign func(*WithModifiedField, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*WithModifiedField)
	for i := range nodes {
		if nodes[i].with_modified_field_owner == nil {
			continue
		}
		fk := *nodes[i].with_modified_field_owner
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "with_modified_field_owner" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wmfq *WithModifiedFieldQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wmfq.querySpec()
	_spec.Node.Columns = wmfq.ctx.Fields
	if len(wmfq.ctx.Fields) > 0 {
		_spec.Unique = wmfq.ctx.Unique != nil && *wmfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wmfq.driver, _spec)
}

func (wmfq *WithModifiedFieldQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(withmodifiedfield.Table, withmodifiedfield.Columns, sqlgraph.NewFieldSpec(withmodifiedfield.FieldID, field.TypeInt))
	_spec.From = wmfq.sql
	if unique := wmfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wmfq.path != nil {
		_spec.Unique = true
	}
	if fields := wmfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withmodifiedfield.FieldID)
		for i := range fields {
			if fields[i] != withmodifiedfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wmfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wmfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wmfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wmfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wmfq *WithModifiedFieldQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wmfq.driver.Dialect())
	t1 := builder.Table(withmodifiedfield.Table)
	columns := wmfq.ctx.Fields
	if len(columns) == 0 {
		columns = withmodifiedfield.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wmfq.sql != nil {
		selector = wmfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wmfq.ctx.Unique != nil && *wmfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wmfq.predicates {
		p(selector)
	}
	for _, p := range wmfq.order {
		p(selector)
	}
	if offset := wmfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wmfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithModifiedFieldGroupBy is the group-by builder for WithModifiedField entities.
type WithModifiedFieldGroupBy struct {
	selector
	build *WithModifiedFieldQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wmfgb *WithModifiedFieldGroupBy) Aggregate(fns ...AggregateFunc) *WithModifiedFieldGroupBy {
	wmfgb.fns = append(wmfgb.fns, fns...)
	return wmfgb
}

// Scan applies the selector query and scans the result into the given value.
func (wmfgb *WithModifiedFieldGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wmfgb.build.ctx, ent.OpQueryGroupBy)
	if err := wmfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithModifiedFieldQuery, *WithModifiedFieldGroupBy](ctx, wmfgb.build, wmfgb, wmfgb.build.inters, v)
}

func (wmfgb *WithModifiedFieldGroupBy) sqlScan(ctx context.Context, root *WithModifiedFieldQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wmfgb.fns))
	for _, fn := range wmfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wmfgb.flds)+len(wmfgb.fns))
		for _, f := range *wmfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wmfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wmfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WithModifiedFieldSelect is the builder for selecting fields of WithModifiedField entities.
type WithModifiedFieldSelect struct {
	*WithModifiedFieldQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wmfs *WithModifiedFieldSelect) Aggregate(fns ...AggregateFunc) *WithModifiedFieldSelect {
	wmfs.fns = append(wmfs.fns, fns...)
	return wmfs
}

// Scan applies the selector query and scans the result into the given value.
func (wmfs *WithModifiedFieldSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wmfs.ctx, ent.OpQuerySelect)
	if err := wmfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithModifiedFieldQuery, *WithModifiedFieldSelect](ctx, wmfs.WithModifiedFieldQuery, wmfs, wmfs.inters, v)
}

func (wmfs *WithModifiedFieldSelect) sqlScan(ctx context.Context, root *WithModifiedFieldQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wmfs.fns))
	for _, fn := range wmfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wmfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wmfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
