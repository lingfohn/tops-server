// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/lingfohn/lime/ent/k8scluster"
	"github.com/lingfohn/lime/ent/namespace"
	"github.com/lingfohn/lime/ent/predicate"
)

// K8sClusterQuery is the builder for querying K8sCluster entities.
type K8sClusterQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.K8sCluster
	// eager-loading edges.
	withNamespaces *NamespaceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (kcq *K8sClusterQuery) Where(ps ...predicate.K8sCluster) *K8sClusterQuery {
	kcq.predicates = append(kcq.predicates, ps...)
	return kcq
}

// Limit adds a limit step to the query.
func (kcq *K8sClusterQuery) Limit(limit int) *K8sClusterQuery {
	kcq.limit = &limit
	return kcq
}

// Offset adds an offset step to the query.
func (kcq *K8sClusterQuery) Offset(offset int) *K8sClusterQuery {
	kcq.offset = &offset
	return kcq
}

// Order adds an order step to the query.
func (kcq *K8sClusterQuery) Order(o ...Order) *K8sClusterQuery {
	kcq.order = append(kcq.order, o...)
	return kcq
}

// QueryNamespaces chains the current query on the namespaces edge.
func (kcq *K8sClusterQuery) QueryNamespaces() *NamespaceQuery {
	query := &NamespaceQuery{config: kcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(k8scluster.Table, k8scluster.FieldID, kcq.sqlQuery()),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, k8scluster.NamespacesTable, k8scluster.NamespacesColumn),
		)
		fromU = sqlgraph.SetNeighbors(kcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first K8sCluster entity in the query. Returns *NotFoundError when no k8scluster was found.
func (kcq *K8sClusterQuery) First(ctx context.Context) (*K8sCluster, error) {
	kcs, err := kcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(kcs) == 0 {
		return nil, &NotFoundError{k8scluster.Label}
	}
	return kcs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kcq *K8sClusterQuery) FirstX(ctx context.Context) *K8sCluster {
	kc, err := kcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return kc
}

// FirstID returns the first K8sCluster id in the query. Returns *NotFoundError when no id was found.
func (kcq *K8sClusterQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{k8scluster.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (kcq *K8sClusterQuery) FirstXID(ctx context.Context) int {
	id, err := kcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only K8sCluster entity in the query, returns an error if not exactly one entity was returned.
func (kcq *K8sClusterQuery) Only(ctx context.Context) (*K8sCluster, error) {
	kcs, err := kcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(kcs) {
	case 1:
		return kcs[0], nil
	case 0:
		return nil, &NotFoundError{k8scluster.Label}
	default:
		return nil, &NotSingularError{k8scluster.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kcq *K8sClusterQuery) OnlyX(ctx context.Context) *K8sCluster {
	kc, err := kcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return kc
}

// OnlyID returns the only K8sCluster id in the query, returns an error if not exactly one id was returned.
func (kcq *K8sClusterQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{k8scluster.Label}
	default:
		err = &NotSingularError{k8scluster.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (kcq *K8sClusterQuery) OnlyXID(ctx context.Context) int {
	id, err := kcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of K8sClusters.
func (kcq *K8sClusterQuery) All(ctx context.Context) ([]*K8sCluster, error) {
	if err := kcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kcq *K8sClusterQuery) AllX(ctx context.Context) []*K8sCluster {
	kcs, err := kcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return kcs
}

// IDs executes the query and returns a list of K8sCluster ids.
func (kcq *K8sClusterQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := kcq.Select(k8scluster.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kcq *K8sClusterQuery) IDsX(ctx context.Context) []int {
	ids, err := kcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kcq *K8sClusterQuery) Count(ctx context.Context) (int, error) {
	if err := kcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kcq *K8sClusterQuery) CountX(ctx context.Context) int {
	count, err := kcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kcq *K8sClusterQuery) Exist(ctx context.Context) (bool, error) {
	if err := kcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kcq *K8sClusterQuery) ExistX(ctx context.Context) bool {
	exist, err := kcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kcq *K8sClusterQuery) Clone() *K8sClusterQuery {
	return &K8sClusterQuery{
		config:     kcq.config,
		limit:      kcq.limit,
		offset:     kcq.offset,
		order:      append([]Order{}, kcq.order...),
		unique:     append([]string{}, kcq.unique...),
		predicates: append([]predicate.K8sCluster{}, kcq.predicates...),
		// clone intermediate query.
		sql:  kcq.sql.Clone(),
		path: kcq.path,
	}
}

//  WithNamespaces tells the query-builder to eager-loads the nodes that are connected to
// the "namespaces" edge. The optional arguments used to configure the query builder of the edge.
func (kcq *K8sClusterQuery) WithNamespaces(opts ...func(*NamespaceQuery)) *K8sClusterQuery {
	query := &NamespaceQuery{config: kcq.config}
	for _, opt := range opts {
		opt(query)
	}
	kcq.withNamespaces = query
	return kcq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Cluster string `json:"cluster,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.K8sCluster.Query().
//		GroupBy(k8scluster.FieldCluster).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (kcq *K8sClusterQuery) GroupBy(field string, fields ...string) *K8sClusterGroupBy {
	group := &K8sClusterGroupBy{config: kcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kcq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Cluster string `json:"cluster,omitempty"`
//	}
//
//	client.K8sCluster.Query().
//		Select(k8scluster.FieldCluster).
//		Scan(ctx, &v)
//
func (kcq *K8sClusterQuery) Select(field string, fields ...string) *K8sClusterSelect {
	selector := &K8sClusterSelect{config: kcq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kcq.sqlQuery(), nil
	}
	return selector
}

func (kcq *K8sClusterQuery) prepareQuery(ctx context.Context) error {
	if kcq.path != nil {
		prev, err := kcq.path(ctx)
		if err != nil {
			return err
		}
		kcq.sql = prev
	}
	return nil
}

func (kcq *K8sClusterQuery) sqlAll(ctx context.Context) ([]*K8sCluster, error) {
	var (
		nodes       = []*K8sCluster{}
		_spec       = kcq.querySpec()
		loadedTypes = [1]bool{
			kcq.withNamespaces != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &K8sCluster{config: kcq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, kcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := kcq.withNamespaces; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*K8sCluster)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Namespace(func(s *sql.Selector) {
			s.Where(sql.InValues(k8scluster.NamespacesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.k8s_cluster_namespaces
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "k8s_cluster_namespaces" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "k8s_cluster_namespaces" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Namespaces = append(node.Edges.Namespaces, n)
		}
	}

	return nodes, nil
}

func (kcq *K8sClusterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kcq.querySpec()
	return sqlgraph.CountNodes(ctx, kcq.driver, _spec)
}

func (kcq *K8sClusterQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (kcq *K8sClusterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   k8scluster.Table,
			Columns: k8scluster.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: k8scluster.FieldID,
			},
		},
		From:   kcq.sql,
		Unique: true,
	}
	if ps := kcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kcq *K8sClusterQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(kcq.driver.Dialect())
	t1 := builder.Table(k8scluster.Table)
	selector := builder.Select(t1.Columns(k8scluster.Columns...)...).From(t1)
	if kcq.sql != nil {
		selector = kcq.sql
		selector.Select(selector.Columns(k8scluster.Columns...)...)
	}
	for _, p := range kcq.predicates {
		p(selector)
	}
	for _, p := range kcq.order {
		p(selector)
	}
	if offset := kcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// K8sClusterGroupBy is the builder for group-by K8sCluster entities.
type K8sClusterGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kcgb *K8sClusterGroupBy) Aggregate(fns ...Aggregate) *K8sClusterGroupBy {
	kcgb.fns = append(kcgb.fns, fns...)
	return kcgb
}

// Scan applies the group-by query and scan the result into the given value.
func (kcgb *K8sClusterGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kcgb.path(ctx)
	if err != nil {
		return err
	}
	kcgb.sql = query
	return kcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kcgb *K8sClusterGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := kcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (kcgb *K8sClusterGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(kcgb.fields) > 1 {
		return nil, errors.New("ent: K8sClusterGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := kcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kcgb *K8sClusterGroupBy) StringsX(ctx context.Context) []string {
	v, err := kcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (kcgb *K8sClusterGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(kcgb.fields) > 1 {
		return nil, errors.New("ent: K8sClusterGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := kcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kcgb *K8sClusterGroupBy) IntsX(ctx context.Context) []int {
	v, err := kcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (kcgb *K8sClusterGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(kcgb.fields) > 1 {
		return nil, errors.New("ent: K8sClusterGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := kcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kcgb *K8sClusterGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := kcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (kcgb *K8sClusterGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(kcgb.fields) > 1 {
		return nil, errors.New("ent: K8sClusterGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := kcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kcgb *K8sClusterGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := kcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kcgb *K8sClusterGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := kcgb.sqlQuery().Query()
	if err := kcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kcgb *K8sClusterGroupBy) sqlQuery() *sql.Selector {
	selector := kcgb.sql
	columns := make([]string, 0, len(kcgb.fields)+len(kcgb.fns))
	columns = append(columns, kcgb.fields...)
	for _, fn := range kcgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(kcgb.fields...)
}

// K8sClusterSelect is the builder for select fields of K8sCluster entities.
type K8sClusterSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (kcs *K8sClusterSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := kcs.path(ctx)
	if err != nil {
		return err
	}
	kcs.sql = query
	return kcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kcs *K8sClusterSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (kcs *K8sClusterSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kcs.fields) > 1 {
		return nil, errors.New("ent: K8sClusterSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kcs *K8sClusterSelect) StringsX(ctx context.Context) []string {
	v, err := kcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (kcs *K8sClusterSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kcs.fields) > 1 {
		return nil, errors.New("ent: K8sClusterSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kcs *K8sClusterSelect) IntsX(ctx context.Context) []int {
	v, err := kcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (kcs *K8sClusterSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kcs.fields) > 1 {
		return nil, errors.New("ent: K8sClusterSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kcs *K8sClusterSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (kcs *K8sClusterSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kcs.fields) > 1 {
		return nil, errors.New("ent: K8sClusterSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kcs *K8sClusterSelect) BoolsX(ctx context.Context) []bool {
	v, err := kcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kcs *K8sClusterSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := kcs.sqlQuery().Query()
	if err := kcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kcs *K8sClusterSelect) sqlQuery() sql.Querier {
	selector := kcs.sql
	selector.Select(selector.Columns(kcs.fields...)...)
	return selector
}