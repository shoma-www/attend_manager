// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/grpc/ent/migrate"

	"github.com/shoma-www/attend_manager/grpc/ent/attendancegroup"
	"github.com/shoma-www/attend_manager/grpc/ent/user"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AttendanceGroup is the client for interacting with the AttendanceGroup builders.
	AttendanceGroup *AttendanceGroupClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AttendanceGroup = NewAttendanceGroupClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		AttendanceGroup: NewAttendanceGroupClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:          cfg,
		AttendanceGroup: NewAttendanceGroupClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AttendanceGroup.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AttendanceGroup.Use(hooks...)
	c.User.Use(hooks...)
}

// AttendanceGroupClient is a client for the AttendanceGroup schema.
type AttendanceGroupClient struct {
	config
}

// NewAttendanceGroupClient returns a client for the AttendanceGroup from the given config.
func NewAttendanceGroupClient(c config) *AttendanceGroupClient {
	return &AttendanceGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attendancegroup.Hooks(f(g(h())))`.
func (c *AttendanceGroupClient) Use(hooks ...Hook) {
	c.hooks.AttendanceGroup = append(c.hooks.AttendanceGroup, hooks...)
}

// Create returns a create builder for AttendanceGroup.
func (c *AttendanceGroupClient) Create() *AttendanceGroupCreate {
	mutation := newAttendanceGroupMutation(c.config, OpCreate)
	return &AttendanceGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of AttendanceGroup entities.
func (c *AttendanceGroupClient) CreateBulk(builders ...*AttendanceGroupCreate) *AttendanceGroupCreateBulk {
	return &AttendanceGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AttendanceGroup.
func (c *AttendanceGroupClient) Update() *AttendanceGroupUpdate {
	mutation := newAttendanceGroupMutation(c.config, OpUpdate)
	return &AttendanceGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttendanceGroupClient) UpdateOne(ag *AttendanceGroup) *AttendanceGroupUpdateOne {
	mutation := newAttendanceGroupMutation(c.config, OpUpdateOne, withAttendanceGroup(ag))
	return &AttendanceGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttendanceGroupClient) UpdateOneID(id xid.ID) *AttendanceGroupUpdateOne {
	mutation := newAttendanceGroupMutation(c.config, OpUpdateOne, withAttendanceGroupID(id))
	return &AttendanceGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AttendanceGroup.
func (c *AttendanceGroupClient) Delete() *AttendanceGroupDelete {
	mutation := newAttendanceGroupMutation(c.config, OpDelete)
	return &AttendanceGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AttendanceGroupClient) DeleteOne(ag *AttendanceGroup) *AttendanceGroupDeleteOne {
	return c.DeleteOneID(ag.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AttendanceGroupClient) DeleteOneID(id xid.ID) *AttendanceGroupDeleteOne {
	builder := c.Delete().Where(attendancegroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttendanceGroupDeleteOne{builder}
}

// Query returns a query builder for AttendanceGroup.
func (c *AttendanceGroupClient) Query() *AttendanceGroupQuery {
	return &AttendanceGroupQuery{config: c.config}
}

// Get returns a AttendanceGroup entity by its id.
func (c *AttendanceGroupClient) Get(ctx context.Context, id xid.ID) (*AttendanceGroup, error) {
	return c.Query().Where(attendancegroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttendanceGroupClient) GetX(ctx context.Context, id xid.ID) *AttendanceGroup {
	ag, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ag
}

// QueryUsers queries the users edge of a AttendanceGroup.
func (c *AttendanceGroupClient) QueryUsers(ag *AttendanceGroup) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ag.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(attendancegroup.Table, attendancegroup.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, attendancegroup.UsersTable, attendancegroup.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(ag.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AttendanceGroupClient) Hooks() []Hook {
	return c.hooks.AttendanceGroup
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id xid.ID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id xid.ID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id xid.ID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id xid.ID) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryGroup queries the group edge of a User.
func (c *UserClient) QueryGroup(u *User) *AttendanceGroupQuery {
	query := &AttendanceGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(attendancegroup.Table, attendancegroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.GroupTable, user.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
