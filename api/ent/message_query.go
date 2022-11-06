// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"simple_sns_api/ent/message"
	"simple_sns_api/ent/post"
	"simple_sns_api/ent/predicate"
	"simple_sns_api/ent/room"
	"simple_sns_api/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MessageQuery is the builder for querying Message entities.
type MessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Message
	withRoom   *RoomQuery
	withUser   *UserQuery
	withPost   *PostQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageQuery builder.
func (mq *MessageQuery) Where(ps ...predicate.Message) *MessageQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MessageQuery) Limit(limit int) *MessageQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MessageQuery) Offset(offset int) *MessageQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MessageQuery) Unique(unique bool) *MessageQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MessageQuery) Order(o ...OrderFunc) *MessageQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryRoom chains the current query on the "room" edge.
func (mq *MessageQuery) QueryRoom() *RoomQuery {
	query := &RoomQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.RoomTable, message.RoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (mq *MessageQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.UserTable, message.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPost chains the current query on the "post" edge.
func (mq *MessageQuery) QueryPost() *PostQuery {
	query := &PostQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.PostTable, message.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Message entity from the query.
// Returns a *NotFoundError when no Message was found.
func (mq *MessageQuery) First(ctx context.Context) (*Message, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{message.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MessageQuery) FirstX(ctx context.Context) *Message {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Message ID from the query.
// Returns a *NotFoundError when no Message ID was found.
func (mq *MessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{message.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MessageQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Message entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Message entity is found.
// Returns a *NotFoundError when no Message entities are found.
func (mq *MessageQuery) Only(ctx context.Context) (*Message, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{message.Label}
	default:
		return nil, &NotSingularError{message.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MessageQuery) OnlyX(ctx context.Context) *Message {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Message ID in the query.
// Returns a *NotSingularError when more than one Message ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{message.Label}
	default:
		err = &NotSingularError{message.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Messages.
func (mq *MessageQuery) All(ctx context.Context) ([]*Message, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MessageQuery) AllX(ctx context.Context) []*Message {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Message IDs.
func (mq *MessageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(message.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MessageQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MessageQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MessageQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MessageQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MessageQuery) Clone() *MessageQuery {
	if mq == nil {
		return nil
	}
	return &MessageQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]OrderFunc{}, mq.order...),
		predicates: append([]predicate.Message{}, mq.predicates...),
		withRoom:   mq.withRoom.Clone(),
		withUser:   mq.withUser.Clone(),
		withPost:   mq.withPost.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithRoom tells the query-builder to eager-load the nodes that are connected to
// the "room" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithRoom(opts ...func(*RoomQuery)) *MessageQuery {
	query := &RoomQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withRoom = query
	return mq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithUser(opts ...func(*UserQuery)) *MessageQuery {
	query := &UserQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withUser = query
	return mq
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithPost(opts ...func(*PostQuery)) *MessageQuery {
	query := &PostQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withPost = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Message.Query().
//		GroupBy(message.FieldContent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MessageQuery) GroupBy(field string, fields ...string) *MessageGroupBy {
	grbuild := &MessageGroupBy{config: mq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	grbuild.label = message.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//	}
//
//	client.Message.Query().
//		Select(message.FieldContent).
//		Scan(ctx, &v)
func (mq *MessageQuery) Select(fields ...string) *MessageSelect {
	mq.fields = append(mq.fields, fields...)
	selbuild := &MessageSelect{MessageQuery: mq}
	selbuild.label = message.Label
	selbuild.flds, selbuild.scan = &mq.fields, selbuild.Scan
	return selbuild
}

func (mq *MessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !message.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Message, error) {
	var (
		nodes       = []*Message{}
		_spec       = mq.querySpec()
		loadedTypes = [3]bool{
			mq.withRoom != nil,
			mq.withUser != nil,
			mq.withPost != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Message).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Message{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withRoom; query != nil {
		if err := mq.loadRoom(ctx, query, nodes, nil,
			func(n *Message, e *Room) { n.Edges.Room = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withUser; query != nil {
		if err := mq.loadUser(ctx, query, nodes, nil,
			func(n *Message, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withPost; query != nil {
		if err := mq.loadPost(ctx, query, nodes, nil,
			func(n *Message, e *Post) { n.Edges.Post = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MessageQuery) loadRoom(ctx context.Context, query *RoomQuery, nodes []*Message, init func(*Message), assign func(*Message, *Room)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Message)
	for i := range nodes {
		fk := nodes[i].RoomID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(room.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "room_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MessageQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Message, init func(*Message), assign func(*Message, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Message)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MessageQuery) loadPost(ctx context.Context, query *PostQuery, nodes []*Message, init func(*Message), assign func(*Message, *Post)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Message)
	for i := range nodes {
		fk := nodes[i].PostID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mq *MessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MessageQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mq *MessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for i := range fields {
			if fields[i] != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(message.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = message.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageGroupBy is the group-by builder for Message entities.
type MessageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MessageGroupBy) Aggregate(fns ...AggregateFunc) *MessageGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MessageGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

func (mgb *MessageGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mgb.fields {
		if !message.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MessageGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MessageSelect is the builder for selecting fields of Message entities.
type MessageSelect struct {
	*MessageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MessageSelect) Scan(ctx context.Context, v any) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MessageQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

func (ms *MessageSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
