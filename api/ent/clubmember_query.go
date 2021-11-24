// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/Wanted-Linx/linx-backend/api/ent/clubmember"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
	"github.com/Wanted-Linx/linx-backend/api/ent/student"
)

// ClubMemberQuery is the builder for querying ClubMember entities.
type ClubMemberQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ClubMember
	// eager-loading edges.
	withStudent *StudentQuery
	withClub    *ClubQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClubMemberQuery builder.
func (cmq *ClubMemberQuery) Where(ps ...predicate.ClubMember) *ClubMemberQuery {
	cmq.predicates = append(cmq.predicates, ps...)
	return cmq
}

// Limit adds a limit step to the query.
func (cmq *ClubMemberQuery) Limit(limit int) *ClubMemberQuery {
	cmq.limit = &limit
	return cmq
}

// Offset adds an offset step to the query.
func (cmq *ClubMemberQuery) Offset(offset int) *ClubMemberQuery {
	cmq.offset = &offset
	return cmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cmq *ClubMemberQuery) Unique(unique bool) *ClubMemberQuery {
	cmq.unique = &unique
	return cmq
}

// Order adds an order step to the query.
func (cmq *ClubMemberQuery) Order(o ...OrderFunc) *ClubMemberQuery {
	cmq.order = append(cmq.order, o...)
	return cmq
}

// QueryStudent chains the current query on the "student" edge.
func (cmq *ClubMemberQuery) QueryStudent() *StudentQuery {
	query := &StudentQuery{config: cmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clubmember.Table, clubmember.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, clubmember.StudentTable, clubmember.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClub chains the current query on the "club" edge.
func (cmq *ClubMemberQuery) QueryClub() *ClubQuery {
	query := &ClubQuery{config: cmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clubmember.Table, clubmember.FieldID, selector),
			sqlgraph.To(club.Table, club.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, clubmember.ClubTable, clubmember.ClubColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClubMember entity from the query.
// Returns a *NotFoundError when no ClubMember was found.
func (cmq *ClubMemberQuery) First(ctx context.Context) (*ClubMember, error) {
	nodes, err := cmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{clubmember.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cmq *ClubMemberQuery) FirstX(ctx context.Context) *ClubMember {
	node, err := cmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClubMember ID from the query.
// Returns a *NotFoundError when no ClubMember ID was found.
func (cmq *ClubMemberQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{clubmember.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cmq *ClubMemberQuery) FirstIDX(ctx context.Context) int {
	id, err := cmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClubMember entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ClubMember entity is not found.
// Returns a *NotFoundError when no ClubMember entities are found.
func (cmq *ClubMemberQuery) Only(ctx context.Context) (*ClubMember, error) {
	nodes, err := cmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{clubmember.Label}
	default:
		return nil, &NotSingularError{clubmember.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cmq *ClubMemberQuery) OnlyX(ctx context.Context) *ClubMember {
	node, err := cmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClubMember ID in the query.
// Returns a *NotSingularError when exactly one ClubMember ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cmq *ClubMemberQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = &NotSingularError{clubmember.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cmq *ClubMemberQuery) OnlyIDX(ctx context.Context) int {
	id, err := cmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClubMembers.
func (cmq *ClubMemberQuery) All(ctx context.Context) ([]*ClubMember, error) {
	if err := cmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cmq *ClubMemberQuery) AllX(ctx context.Context) []*ClubMember {
	nodes, err := cmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClubMember IDs.
func (cmq *ClubMemberQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cmq.Select(clubmember.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cmq *ClubMemberQuery) IDsX(ctx context.Context) []int {
	ids, err := cmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cmq *ClubMemberQuery) Count(ctx context.Context) (int, error) {
	if err := cmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cmq *ClubMemberQuery) CountX(ctx context.Context) int {
	count, err := cmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cmq *ClubMemberQuery) Exist(ctx context.Context) (bool, error) {
	if err := cmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cmq *ClubMemberQuery) ExistX(ctx context.Context) bool {
	exist, err := cmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClubMemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cmq *ClubMemberQuery) Clone() *ClubMemberQuery {
	if cmq == nil {
		return nil
	}
	return &ClubMemberQuery{
		config:      cmq.config,
		limit:       cmq.limit,
		offset:      cmq.offset,
		order:       append([]OrderFunc{}, cmq.order...),
		predicates:  append([]predicate.ClubMember{}, cmq.predicates...),
		withStudent: cmq.withStudent.Clone(),
		withClub:    cmq.withClub.Clone(),
		// clone intermediate query.
		sql:  cmq.sql.Clone(),
		path: cmq.path,
	}
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (cmq *ClubMemberQuery) WithStudent(opts ...func(*StudentQuery)) *ClubMemberQuery {
	query := &StudentQuery{config: cmq.config}
	for _, opt := range opts {
		opt(query)
	}
	cmq.withStudent = query
	return cmq
}

// WithClub tells the query-builder to eager-load the nodes that are connected to
// the "club" edge. The optional arguments are used to configure the query builder of the edge.
func (cmq *ClubMemberQuery) WithClub(opts ...func(*ClubQuery)) *ClubMemberQuery {
	query := &ClubQuery{config: cmq.config}
	for _, opt := range opts {
		opt(query)
	}
	cmq.withClub = query
	return cmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClubID int `json:"club_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClubMember.Query().
//		GroupBy(clubmember.FieldClubID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cmq *ClubMemberQuery) GroupBy(field string, fields ...string) *ClubMemberGroupBy {
	group := &ClubMemberGroupBy{config: cmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cmq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClubID int `json:"club_id,omitempty"`
//	}
//
//	client.ClubMember.Query().
//		Select(clubmember.FieldClubID).
//		Scan(ctx, &v)
//
func (cmq *ClubMemberQuery) Select(fields ...string) *ClubMemberSelect {
	cmq.fields = append(cmq.fields, fields...)
	return &ClubMemberSelect{ClubMemberQuery: cmq}
}

func (cmq *ClubMemberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cmq.fields {
		if !clubmember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cmq.path != nil {
		prev, err := cmq.path(ctx)
		if err != nil {
			return err
		}
		cmq.sql = prev
	}
	return nil
}

func (cmq *ClubMemberQuery) sqlAll(ctx context.Context) ([]*ClubMember, error) {
	var (
		nodes       = []*ClubMember{}
		withFKs     = cmq.withFKs
		_spec       = cmq.querySpec()
		loadedTypes = [2]bool{
			cmq.withStudent != nil,
			cmq.withClub != nil,
		}
	)
	if cmq.withStudent != nil || cmq.withClub != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, clubmember.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ClubMember{config: cmq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, cmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cmq.withStudent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ClubMember)
		for i := range nodes {
			fk := nodes[i].StudentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(student.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "student_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Student = n
			}
		}
	}

	if query := cmq.withClub; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ClubMember)
		for i := range nodes {
			fk := nodes[i].ClubID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(club.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "club_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Club = n
			}
		}
	}

	return nodes, nil
}

func (cmq *ClubMemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cmq.querySpec()
	return sqlgraph.CountNodes(ctx, cmq.driver, _spec)
}

func (cmq *ClubMemberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cmq *ClubMemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clubmember.Table,
			Columns: clubmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubmember.FieldID,
			},
		},
		From:   cmq.sql,
		Unique: true,
	}
	if unique := cmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clubmember.FieldID)
		for i := range fields {
			if fields[i] != clubmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cmq *ClubMemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cmq.driver.Dialect())
	t1 := builder.Table(clubmember.Table)
	columns := cmq.fields
	if len(columns) == 0 {
		columns = clubmember.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cmq.sql != nil {
		selector = cmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range cmq.predicates {
		p(selector)
	}
	for _, p := range cmq.order {
		p(selector)
	}
	if offset := cmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClubMemberGroupBy is the group-by builder for ClubMember entities.
type ClubMemberGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cmgb *ClubMemberGroupBy) Aggregate(fns ...AggregateFunc) *ClubMemberGroupBy {
	cmgb.fns = append(cmgb.fns, fns...)
	return cmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cmgb *ClubMemberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cmgb.path(ctx)
	if err != nil {
		return err
	}
	cmgb.sql = query
	return cmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("ent: ClubMemberGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) StringsX(ctx context.Context) []string {
	v, err := cmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) StringX(ctx context.Context) string {
	v, err := cmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("ent: ClubMemberGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) IntsX(ctx context.Context) []int {
	v, err := cmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) IntX(ctx context.Context) int {
	v, err := cmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("ent: ClubMemberGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("ent: ClubMemberGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *ClubMemberGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cmgb *ClubMemberGroupBy) BoolX(ctx context.Context) bool {
	v, err := cmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cmgb *ClubMemberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cmgb.fields {
		if !clubmember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cmgb *ClubMemberGroupBy) sqlQuery() *sql.Selector {
	selector := cmgb.sql.Select()
	aggregation := make([]string, 0, len(cmgb.fns))
	for _, fn := range cmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cmgb.fields)+len(cmgb.fns))
		for _, f := range cmgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cmgb.fields...)...)
}

// ClubMemberSelect is the builder for selecting fields of ClubMember entities.
type ClubMemberSelect struct {
	*ClubMemberQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cms *ClubMemberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cms.prepareQuery(ctx); err != nil {
		return err
	}
	cms.sql = cms.ClubMemberQuery.sqlQuery(ctx)
	return cms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cms *ClubMemberSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("ent: ClubMemberSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cms *ClubMemberSelect) StringsX(ctx context.Context) []string {
	v, err := cms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cms *ClubMemberSelect) StringX(ctx context.Context) string {
	v, err := cms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("ent: ClubMemberSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cms *ClubMemberSelect) IntsX(ctx context.Context) []int {
	v, err := cms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cms *ClubMemberSelect) IntX(ctx context.Context) int {
	v, err := cms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("ent: ClubMemberSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cms *ClubMemberSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cms *ClubMemberSelect) Float64X(ctx context.Context) float64 {
	v, err := cms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("ent: ClubMemberSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cms *ClubMemberSelect) BoolsX(ctx context.Context) []bool {
	v, err := cms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cms *ClubMemberSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubmember.Label}
	default:
		err = fmt.Errorf("ent: ClubMemberSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cms *ClubMemberSelect) BoolX(ctx context.Context) bool {
	v, err := cms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cms *ClubMemberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cms.sql.Query()
	if err := cms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
