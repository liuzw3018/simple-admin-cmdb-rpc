// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdb"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbassets"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbip"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbipuser"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cmdb is the client for interacting with the Cmdb builders.
	Cmdb *CmdbClient
	// CmdbAssets is the client for interacting with the CmdbAssets builders.
	CmdbAssets *CmdbAssetsClient
	// CmdbIp is the client for interacting with the CmdbIp builders.
	CmdbIp *CmdbIpClient
	// CmdbIpUser is the client for interacting with the CmdbIpUser builders.
	CmdbIpUser *CmdbIpUserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Cmdb = NewCmdbClient(c.config)
	c.CmdbAssets = NewCmdbAssetsClient(c.config)
	c.CmdbIp = NewCmdbIpClient(c.config)
	c.CmdbIpUser = NewCmdbIpUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Cmdb:       NewCmdbClient(cfg),
		CmdbAssets: NewCmdbAssetsClient(cfg),
		CmdbIp:     NewCmdbIpClient(cfg),
		CmdbIpUser: NewCmdbIpUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Cmdb:       NewCmdbClient(cfg),
		CmdbAssets: NewCmdbAssetsClient(cfg),
		CmdbIp:     NewCmdbIpClient(cfg),
		CmdbIpUser: NewCmdbIpUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cmdb.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Cmdb.Use(hooks...)
	c.CmdbAssets.Use(hooks...)
	c.CmdbIp.Use(hooks...)
	c.CmdbIpUser.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Cmdb.Intercept(interceptors...)
	c.CmdbAssets.Intercept(interceptors...)
	c.CmdbIp.Intercept(interceptors...)
	c.CmdbIpUser.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CmdbMutation:
		return c.Cmdb.mutate(ctx, m)
	case *CmdbAssetsMutation:
		return c.CmdbAssets.mutate(ctx, m)
	case *CmdbIpMutation:
		return c.CmdbIp.mutate(ctx, m)
	case *CmdbIpUserMutation:
		return c.CmdbIpUser.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CmdbClient is a client for the Cmdb schema.
type CmdbClient struct {
	config
}

// NewCmdbClient returns a client for the Cmdb from the given config.
func NewCmdbClient(c config) *CmdbClient {
	return &CmdbClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cmdb.Hooks(f(g(h())))`.
func (c *CmdbClient) Use(hooks ...Hook) {
	c.hooks.Cmdb = append(c.hooks.Cmdb, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cmdb.Intercept(f(g(h())))`.
func (c *CmdbClient) Intercept(interceptors ...Interceptor) {
	c.inters.Cmdb = append(c.inters.Cmdb, interceptors...)
}

// Create returns a builder for creating a Cmdb entity.
func (c *CmdbClient) Create() *CmdbCreate {
	mutation := newCmdbMutation(c.config, OpCreate)
	return &CmdbCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cmdb entities.
func (c *CmdbClient) CreateBulk(builders ...*CmdbCreate) *CmdbCreateBulk {
	return &CmdbCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CmdbClient) MapCreateBulk(slice any, setFunc func(*CmdbCreate, int)) *CmdbCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CmdbCreateBulk{err: fmt.Errorf("calling to CmdbClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CmdbCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CmdbCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cmdb.
func (c *CmdbClient) Update() *CmdbUpdate {
	mutation := newCmdbMutation(c.config, OpUpdate)
	return &CmdbUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CmdbClient) UpdateOne(cm *Cmdb) *CmdbUpdateOne {
	mutation := newCmdbMutation(c.config, OpUpdateOne, withCmdb(cm))
	return &CmdbUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CmdbClient) UpdateOneID(id int) *CmdbUpdateOne {
	mutation := newCmdbMutation(c.config, OpUpdateOne, withCmdbID(id))
	return &CmdbUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cmdb.
func (c *CmdbClient) Delete() *CmdbDelete {
	mutation := newCmdbMutation(c.config, OpDelete)
	return &CmdbDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CmdbClient) DeleteOne(cm *Cmdb) *CmdbDeleteOne {
	return c.DeleteOneID(cm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CmdbClient) DeleteOneID(id int) *CmdbDeleteOne {
	builder := c.Delete().Where(cmdb.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CmdbDeleteOne{builder}
}

// Query returns a query builder for Cmdb.
func (c *CmdbClient) Query() *CmdbQuery {
	return &CmdbQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCmdb},
		inters: c.Interceptors(),
	}
}

// Get returns a Cmdb entity by its id.
func (c *CmdbClient) Get(ctx context.Context, id int) (*Cmdb, error) {
	return c.Query().Where(cmdb.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CmdbClient) GetX(ctx context.Context, id int) *Cmdb {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CmdbClient) Hooks() []Hook {
	return c.hooks.Cmdb
}

// Interceptors returns the client interceptors.
func (c *CmdbClient) Interceptors() []Interceptor {
	return c.inters.Cmdb
}

func (c *CmdbClient) mutate(ctx context.Context, m *CmdbMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CmdbCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CmdbUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CmdbUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CmdbDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Cmdb mutation op: %q", m.Op())
	}
}

// CmdbAssetsClient is a client for the CmdbAssets schema.
type CmdbAssetsClient struct {
	config
}

// NewCmdbAssetsClient returns a client for the CmdbAssets from the given config.
func NewCmdbAssetsClient(c config) *CmdbAssetsClient {
	return &CmdbAssetsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cmdbassets.Hooks(f(g(h())))`.
func (c *CmdbAssetsClient) Use(hooks ...Hook) {
	c.hooks.CmdbAssets = append(c.hooks.CmdbAssets, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cmdbassets.Intercept(f(g(h())))`.
func (c *CmdbAssetsClient) Intercept(interceptors ...Interceptor) {
	c.inters.CmdbAssets = append(c.inters.CmdbAssets, interceptors...)
}

// Create returns a builder for creating a CmdbAssets entity.
func (c *CmdbAssetsClient) Create() *CmdbAssetsCreate {
	mutation := newCmdbAssetsMutation(c.config, OpCreate)
	return &CmdbAssetsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CmdbAssets entities.
func (c *CmdbAssetsClient) CreateBulk(builders ...*CmdbAssetsCreate) *CmdbAssetsCreateBulk {
	return &CmdbAssetsCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CmdbAssetsClient) MapCreateBulk(slice any, setFunc func(*CmdbAssetsCreate, int)) *CmdbAssetsCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CmdbAssetsCreateBulk{err: fmt.Errorf("calling to CmdbAssetsClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CmdbAssetsCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CmdbAssetsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CmdbAssets.
func (c *CmdbAssetsClient) Update() *CmdbAssetsUpdate {
	mutation := newCmdbAssetsMutation(c.config, OpUpdate)
	return &CmdbAssetsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CmdbAssetsClient) UpdateOne(ca *CmdbAssets) *CmdbAssetsUpdateOne {
	mutation := newCmdbAssetsMutation(c.config, OpUpdateOne, withCmdbAssets(ca))
	return &CmdbAssetsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CmdbAssetsClient) UpdateOneID(id uint64) *CmdbAssetsUpdateOne {
	mutation := newCmdbAssetsMutation(c.config, OpUpdateOne, withCmdbAssetsID(id))
	return &CmdbAssetsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CmdbAssets.
func (c *CmdbAssetsClient) Delete() *CmdbAssetsDelete {
	mutation := newCmdbAssetsMutation(c.config, OpDelete)
	return &CmdbAssetsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CmdbAssetsClient) DeleteOne(ca *CmdbAssets) *CmdbAssetsDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CmdbAssetsClient) DeleteOneID(id uint64) *CmdbAssetsDeleteOne {
	builder := c.Delete().Where(cmdbassets.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CmdbAssetsDeleteOne{builder}
}

// Query returns a query builder for CmdbAssets.
func (c *CmdbAssetsClient) Query() *CmdbAssetsQuery {
	return &CmdbAssetsQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCmdbAssets},
		inters: c.Interceptors(),
	}
}

// Get returns a CmdbAssets entity by its id.
func (c *CmdbAssetsClient) Get(ctx context.Context, id uint64) (*CmdbAssets, error) {
	return c.Query().Where(cmdbassets.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CmdbAssetsClient) GetX(ctx context.Context, id uint64) *CmdbAssets {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CmdbAssetsClient) Hooks() []Hook {
	return c.hooks.CmdbAssets
}

// Interceptors returns the client interceptors.
func (c *CmdbAssetsClient) Interceptors() []Interceptor {
	return c.inters.CmdbAssets
}

func (c *CmdbAssetsClient) mutate(ctx context.Context, m *CmdbAssetsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CmdbAssetsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CmdbAssetsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CmdbAssetsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CmdbAssetsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CmdbAssets mutation op: %q", m.Op())
	}
}

// CmdbIpClient is a client for the CmdbIp schema.
type CmdbIpClient struct {
	config
}

// NewCmdbIpClient returns a client for the CmdbIp from the given config.
func NewCmdbIpClient(c config) *CmdbIpClient {
	return &CmdbIpClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cmdbip.Hooks(f(g(h())))`.
func (c *CmdbIpClient) Use(hooks ...Hook) {
	c.hooks.CmdbIp = append(c.hooks.CmdbIp, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cmdbip.Intercept(f(g(h())))`.
func (c *CmdbIpClient) Intercept(interceptors ...Interceptor) {
	c.inters.CmdbIp = append(c.inters.CmdbIp, interceptors...)
}

// Create returns a builder for creating a CmdbIp entity.
func (c *CmdbIpClient) Create() *CmdbIpCreate {
	mutation := newCmdbIpMutation(c.config, OpCreate)
	return &CmdbIpCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CmdbIp entities.
func (c *CmdbIpClient) CreateBulk(builders ...*CmdbIpCreate) *CmdbIpCreateBulk {
	return &CmdbIpCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CmdbIpClient) MapCreateBulk(slice any, setFunc func(*CmdbIpCreate, int)) *CmdbIpCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CmdbIpCreateBulk{err: fmt.Errorf("calling to CmdbIpClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CmdbIpCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CmdbIpCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CmdbIp.
func (c *CmdbIpClient) Update() *CmdbIpUpdate {
	mutation := newCmdbIpMutation(c.config, OpUpdate)
	return &CmdbIpUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CmdbIpClient) UpdateOne(ci *CmdbIp) *CmdbIpUpdateOne {
	mutation := newCmdbIpMutation(c.config, OpUpdateOne, withCmdbIp(ci))
	return &CmdbIpUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CmdbIpClient) UpdateOneID(id uint64) *CmdbIpUpdateOne {
	mutation := newCmdbIpMutation(c.config, OpUpdateOne, withCmdbIpID(id))
	return &CmdbIpUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CmdbIp.
func (c *CmdbIpClient) Delete() *CmdbIpDelete {
	mutation := newCmdbIpMutation(c.config, OpDelete)
	return &CmdbIpDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CmdbIpClient) DeleteOne(ci *CmdbIp) *CmdbIpDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CmdbIpClient) DeleteOneID(id uint64) *CmdbIpDeleteOne {
	builder := c.Delete().Where(cmdbip.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CmdbIpDeleteOne{builder}
}

// Query returns a query builder for CmdbIp.
func (c *CmdbIpClient) Query() *CmdbIpQuery {
	return &CmdbIpQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCmdbIp},
		inters: c.Interceptors(),
	}
}

// Get returns a CmdbIp entity by its id.
func (c *CmdbIpClient) Get(ctx context.Context, id uint64) (*CmdbIp, error) {
	return c.Query().Where(cmdbip.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CmdbIpClient) GetX(ctx context.Context, id uint64) *CmdbIp {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CmdbIpClient) Hooks() []Hook {
	return c.hooks.CmdbIp
}

// Interceptors returns the client interceptors.
func (c *CmdbIpClient) Interceptors() []Interceptor {
	return c.inters.CmdbIp
}

func (c *CmdbIpClient) mutate(ctx context.Context, m *CmdbIpMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CmdbIpCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CmdbIpUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CmdbIpUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CmdbIpDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CmdbIp mutation op: %q", m.Op())
	}
}

// CmdbIpUserClient is a client for the CmdbIpUser schema.
type CmdbIpUserClient struct {
	config
}

// NewCmdbIpUserClient returns a client for the CmdbIpUser from the given config.
func NewCmdbIpUserClient(c config) *CmdbIpUserClient {
	return &CmdbIpUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cmdbipuser.Hooks(f(g(h())))`.
func (c *CmdbIpUserClient) Use(hooks ...Hook) {
	c.hooks.CmdbIpUser = append(c.hooks.CmdbIpUser, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cmdbipuser.Intercept(f(g(h())))`.
func (c *CmdbIpUserClient) Intercept(interceptors ...Interceptor) {
	c.inters.CmdbIpUser = append(c.inters.CmdbIpUser, interceptors...)
}

// Create returns a builder for creating a CmdbIpUser entity.
func (c *CmdbIpUserClient) Create() *CmdbIpUserCreate {
	mutation := newCmdbIpUserMutation(c.config, OpCreate)
	return &CmdbIpUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CmdbIpUser entities.
func (c *CmdbIpUserClient) CreateBulk(builders ...*CmdbIpUserCreate) *CmdbIpUserCreateBulk {
	return &CmdbIpUserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CmdbIpUserClient) MapCreateBulk(slice any, setFunc func(*CmdbIpUserCreate, int)) *CmdbIpUserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CmdbIpUserCreateBulk{err: fmt.Errorf("calling to CmdbIpUserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CmdbIpUserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CmdbIpUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CmdbIpUser.
func (c *CmdbIpUserClient) Update() *CmdbIpUserUpdate {
	mutation := newCmdbIpUserMutation(c.config, OpUpdate)
	return &CmdbIpUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CmdbIpUserClient) UpdateOne(ciu *CmdbIpUser) *CmdbIpUserUpdateOne {
	mutation := newCmdbIpUserMutation(c.config, OpUpdateOne, withCmdbIpUser(ciu))
	return &CmdbIpUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CmdbIpUserClient) UpdateOneID(id uint64) *CmdbIpUserUpdateOne {
	mutation := newCmdbIpUserMutation(c.config, OpUpdateOne, withCmdbIpUserID(id))
	return &CmdbIpUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CmdbIpUser.
func (c *CmdbIpUserClient) Delete() *CmdbIpUserDelete {
	mutation := newCmdbIpUserMutation(c.config, OpDelete)
	return &CmdbIpUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CmdbIpUserClient) DeleteOne(ciu *CmdbIpUser) *CmdbIpUserDeleteOne {
	return c.DeleteOneID(ciu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CmdbIpUserClient) DeleteOneID(id uint64) *CmdbIpUserDeleteOne {
	builder := c.Delete().Where(cmdbipuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CmdbIpUserDeleteOne{builder}
}

// Query returns a query builder for CmdbIpUser.
func (c *CmdbIpUserClient) Query() *CmdbIpUserQuery {
	return &CmdbIpUserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCmdbIpUser},
		inters: c.Interceptors(),
	}
}

// Get returns a CmdbIpUser entity by its id.
func (c *CmdbIpUserClient) Get(ctx context.Context, id uint64) (*CmdbIpUser, error) {
	return c.Query().Where(cmdbipuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CmdbIpUserClient) GetX(ctx context.Context, id uint64) *CmdbIpUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CmdbIpUserClient) Hooks() []Hook {
	return c.hooks.CmdbIpUser
}

// Interceptors returns the client interceptors.
func (c *CmdbIpUserClient) Interceptors() []Interceptor {
	return c.inters.CmdbIpUser
}

func (c *CmdbIpUserClient) mutate(ctx context.Context, m *CmdbIpUserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CmdbIpUserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CmdbIpUserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CmdbIpUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CmdbIpUserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CmdbIpUser mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Cmdb, CmdbAssets, CmdbIp, CmdbIpUser []ent.Hook
	}
	inters struct {
		Cmdb, CmdbAssets, CmdbIp, CmdbIpUser []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}