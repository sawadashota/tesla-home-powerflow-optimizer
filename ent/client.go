// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargecommandhistory"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargesetting"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargestatecache"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/grant"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/migrate"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/powermetric"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ChargeCommandHistory is the client for interacting with the ChargeCommandHistory builders.
	ChargeCommandHistory *ChargeCommandHistoryClient
	// ChargeSetting is the client for interacting with the ChargeSetting builders.
	ChargeSetting *ChargeSettingClient
	// ChargeStateCache is the client for interacting with the ChargeStateCache builders.
	ChargeStateCache *ChargeStateCacheClient
	// Grant is the client for interacting with the Grant builders.
	Grant *GrantClient
	// PowerMetric is the client for interacting with the PowerMetric builders.
	PowerMetric *PowerMetricClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ChargeCommandHistory = NewChargeCommandHistoryClient(c.config)
	c.ChargeSetting = NewChargeSettingClient(c.config)
	c.ChargeStateCache = NewChargeStateCacheClient(c.config)
	c.Grant = NewGrantClient(c.config)
	c.PowerMetric = NewPowerMetricClient(c.config)
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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
		ctx:                  ctx,
		config:               cfg,
		ChargeCommandHistory: NewChargeCommandHistoryClient(cfg),
		ChargeSetting:        NewChargeSettingClient(cfg),
		ChargeStateCache:     NewChargeStateCacheClient(cfg),
		Grant:                NewGrantClient(cfg),
		PowerMetric:          NewPowerMetricClient(cfg),
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
		ctx:                  ctx,
		config:               cfg,
		ChargeCommandHistory: NewChargeCommandHistoryClient(cfg),
		ChargeSetting:        NewChargeSettingClient(cfg),
		ChargeStateCache:     NewChargeStateCacheClient(cfg),
		Grant:                NewGrantClient(cfg),
		PowerMetric:          NewPowerMetricClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ChargeCommandHistory.
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
	c.ChargeCommandHistory.Use(hooks...)
	c.ChargeSetting.Use(hooks...)
	c.ChargeStateCache.Use(hooks...)
	c.Grant.Use(hooks...)
	c.PowerMetric.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ChargeCommandHistory.Intercept(interceptors...)
	c.ChargeSetting.Intercept(interceptors...)
	c.ChargeStateCache.Intercept(interceptors...)
	c.Grant.Intercept(interceptors...)
	c.PowerMetric.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChargeCommandHistoryMutation:
		return c.ChargeCommandHistory.mutate(ctx, m)
	case *ChargeSettingMutation:
		return c.ChargeSetting.mutate(ctx, m)
	case *ChargeStateCacheMutation:
		return c.ChargeStateCache.mutate(ctx, m)
	case *GrantMutation:
		return c.Grant.mutate(ctx, m)
	case *PowerMetricMutation:
		return c.PowerMetric.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ChargeCommandHistoryClient is a client for the ChargeCommandHistory schema.
type ChargeCommandHistoryClient struct {
	config
}

// NewChargeCommandHistoryClient returns a client for the ChargeCommandHistory from the given config.
func NewChargeCommandHistoryClient(c config) *ChargeCommandHistoryClient {
	return &ChargeCommandHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chargecommandhistory.Hooks(f(g(h())))`.
func (c *ChargeCommandHistoryClient) Use(hooks ...Hook) {
	c.hooks.ChargeCommandHistory = append(c.hooks.ChargeCommandHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chargecommandhistory.Intercept(f(g(h())))`.
func (c *ChargeCommandHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChargeCommandHistory = append(c.inters.ChargeCommandHistory, interceptors...)
}

// Create returns a builder for creating a ChargeCommandHistory entity.
func (c *ChargeCommandHistoryClient) Create() *ChargeCommandHistoryCreate {
	mutation := newChargeCommandHistoryMutation(c.config, OpCreate)
	return &ChargeCommandHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChargeCommandHistory entities.
func (c *ChargeCommandHistoryClient) CreateBulk(builders ...*ChargeCommandHistoryCreate) *ChargeCommandHistoryCreateBulk {
	return &ChargeCommandHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChargeCommandHistoryClient) MapCreateBulk(slice any, setFunc func(*ChargeCommandHistoryCreate, int)) *ChargeCommandHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChargeCommandHistoryCreateBulk{err: fmt.Errorf("calling to ChargeCommandHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChargeCommandHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChargeCommandHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChargeCommandHistory.
func (c *ChargeCommandHistoryClient) Update() *ChargeCommandHistoryUpdate {
	mutation := newChargeCommandHistoryMutation(c.config, OpUpdate)
	return &ChargeCommandHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChargeCommandHistoryClient) UpdateOne(cch *ChargeCommandHistory) *ChargeCommandHistoryUpdateOne {
	mutation := newChargeCommandHistoryMutation(c.config, OpUpdateOne, withChargeCommandHistory(cch))
	return &ChargeCommandHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChargeCommandHistoryClient) UpdateOneID(id int) *ChargeCommandHistoryUpdateOne {
	mutation := newChargeCommandHistoryMutation(c.config, OpUpdateOne, withChargeCommandHistoryID(id))
	return &ChargeCommandHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChargeCommandHistory.
func (c *ChargeCommandHistoryClient) Delete() *ChargeCommandHistoryDelete {
	mutation := newChargeCommandHistoryMutation(c.config, OpDelete)
	return &ChargeCommandHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChargeCommandHistoryClient) DeleteOne(cch *ChargeCommandHistory) *ChargeCommandHistoryDeleteOne {
	return c.DeleteOneID(cch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChargeCommandHistoryClient) DeleteOneID(id int) *ChargeCommandHistoryDeleteOne {
	builder := c.Delete().Where(chargecommandhistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChargeCommandHistoryDeleteOne{builder}
}

// Query returns a query builder for ChargeCommandHistory.
func (c *ChargeCommandHistoryClient) Query() *ChargeCommandHistoryQuery {
	return &ChargeCommandHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChargeCommandHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a ChargeCommandHistory entity by its id.
func (c *ChargeCommandHistoryClient) Get(ctx context.Context, id int) (*ChargeCommandHistory, error) {
	return c.Query().Where(chargecommandhistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChargeCommandHistoryClient) GetX(ctx context.Context, id int) *ChargeCommandHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChargeCommandHistoryClient) Hooks() []Hook {
	return c.hooks.ChargeCommandHistory
}

// Interceptors returns the client interceptors.
func (c *ChargeCommandHistoryClient) Interceptors() []Interceptor {
	return c.inters.ChargeCommandHistory
}

func (c *ChargeCommandHistoryClient) mutate(ctx context.Context, m *ChargeCommandHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChargeCommandHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChargeCommandHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChargeCommandHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChargeCommandHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChargeCommandHistory mutation op: %q", m.Op())
	}
}

// ChargeSettingClient is a client for the ChargeSetting schema.
type ChargeSettingClient struct {
	config
}

// NewChargeSettingClient returns a client for the ChargeSetting from the given config.
func NewChargeSettingClient(c config) *ChargeSettingClient {
	return &ChargeSettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chargesetting.Hooks(f(g(h())))`.
func (c *ChargeSettingClient) Use(hooks ...Hook) {
	c.hooks.ChargeSetting = append(c.hooks.ChargeSetting, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chargesetting.Intercept(f(g(h())))`.
func (c *ChargeSettingClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChargeSetting = append(c.inters.ChargeSetting, interceptors...)
}

// Create returns a builder for creating a ChargeSetting entity.
func (c *ChargeSettingClient) Create() *ChargeSettingCreate {
	mutation := newChargeSettingMutation(c.config, OpCreate)
	return &ChargeSettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChargeSetting entities.
func (c *ChargeSettingClient) CreateBulk(builders ...*ChargeSettingCreate) *ChargeSettingCreateBulk {
	return &ChargeSettingCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChargeSettingClient) MapCreateBulk(slice any, setFunc func(*ChargeSettingCreate, int)) *ChargeSettingCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChargeSettingCreateBulk{err: fmt.Errorf("calling to ChargeSettingClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChargeSettingCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChargeSettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChargeSetting.
func (c *ChargeSettingClient) Update() *ChargeSettingUpdate {
	mutation := newChargeSettingMutation(c.config, OpUpdate)
	return &ChargeSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChargeSettingClient) UpdateOne(cs *ChargeSetting) *ChargeSettingUpdateOne {
	mutation := newChargeSettingMutation(c.config, OpUpdateOne, withChargeSetting(cs))
	return &ChargeSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChargeSettingClient) UpdateOneID(id int) *ChargeSettingUpdateOne {
	mutation := newChargeSettingMutation(c.config, OpUpdateOne, withChargeSettingID(id))
	return &ChargeSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChargeSetting.
func (c *ChargeSettingClient) Delete() *ChargeSettingDelete {
	mutation := newChargeSettingMutation(c.config, OpDelete)
	return &ChargeSettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChargeSettingClient) DeleteOne(cs *ChargeSetting) *ChargeSettingDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChargeSettingClient) DeleteOneID(id int) *ChargeSettingDeleteOne {
	builder := c.Delete().Where(chargesetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChargeSettingDeleteOne{builder}
}

// Query returns a query builder for ChargeSetting.
func (c *ChargeSettingClient) Query() *ChargeSettingQuery {
	return &ChargeSettingQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChargeSetting},
		inters: c.Interceptors(),
	}
}

// Get returns a ChargeSetting entity by its id.
func (c *ChargeSettingClient) Get(ctx context.Context, id int) (*ChargeSetting, error) {
	return c.Query().Where(chargesetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChargeSettingClient) GetX(ctx context.Context, id int) *ChargeSetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChargeSettingClient) Hooks() []Hook {
	return c.hooks.ChargeSetting
}

// Interceptors returns the client interceptors.
func (c *ChargeSettingClient) Interceptors() []Interceptor {
	return c.inters.ChargeSetting
}

func (c *ChargeSettingClient) mutate(ctx context.Context, m *ChargeSettingMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChargeSettingCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChargeSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChargeSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChargeSettingDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChargeSetting mutation op: %q", m.Op())
	}
}

// ChargeStateCacheClient is a client for the ChargeStateCache schema.
type ChargeStateCacheClient struct {
	config
}

// NewChargeStateCacheClient returns a client for the ChargeStateCache from the given config.
func NewChargeStateCacheClient(c config) *ChargeStateCacheClient {
	return &ChargeStateCacheClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chargestatecache.Hooks(f(g(h())))`.
func (c *ChargeStateCacheClient) Use(hooks ...Hook) {
	c.hooks.ChargeStateCache = append(c.hooks.ChargeStateCache, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chargestatecache.Intercept(f(g(h())))`.
func (c *ChargeStateCacheClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChargeStateCache = append(c.inters.ChargeStateCache, interceptors...)
}

// Create returns a builder for creating a ChargeStateCache entity.
func (c *ChargeStateCacheClient) Create() *ChargeStateCacheCreate {
	mutation := newChargeStateCacheMutation(c.config, OpCreate)
	return &ChargeStateCacheCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChargeStateCache entities.
func (c *ChargeStateCacheClient) CreateBulk(builders ...*ChargeStateCacheCreate) *ChargeStateCacheCreateBulk {
	return &ChargeStateCacheCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChargeStateCacheClient) MapCreateBulk(slice any, setFunc func(*ChargeStateCacheCreate, int)) *ChargeStateCacheCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChargeStateCacheCreateBulk{err: fmt.Errorf("calling to ChargeStateCacheClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChargeStateCacheCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChargeStateCacheCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChargeStateCache.
func (c *ChargeStateCacheClient) Update() *ChargeStateCacheUpdate {
	mutation := newChargeStateCacheMutation(c.config, OpUpdate)
	return &ChargeStateCacheUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChargeStateCacheClient) UpdateOne(csc *ChargeStateCache) *ChargeStateCacheUpdateOne {
	mutation := newChargeStateCacheMutation(c.config, OpUpdateOne, withChargeStateCache(csc))
	return &ChargeStateCacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChargeStateCacheClient) UpdateOneID(id int) *ChargeStateCacheUpdateOne {
	mutation := newChargeStateCacheMutation(c.config, OpUpdateOne, withChargeStateCacheID(id))
	return &ChargeStateCacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChargeStateCache.
func (c *ChargeStateCacheClient) Delete() *ChargeStateCacheDelete {
	mutation := newChargeStateCacheMutation(c.config, OpDelete)
	return &ChargeStateCacheDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChargeStateCacheClient) DeleteOne(csc *ChargeStateCache) *ChargeStateCacheDeleteOne {
	return c.DeleteOneID(csc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChargeStateCacheClient) DeleteOneID(id int) *ChargeStateCacheDeleteOne {
	builder := c.Delete().Where(chargestatecache.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChargeStateCacheDeleteOne{builder}
}

// Query returns a query builder for ChargeStateCache.
func (c *ChargeStateCacheClient) Query() *ChargeStateCacheQuery {
	return &ChargeStateCacheQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChargeStateCache},
		inters: c.Interceptors(),
	}
}

// Get returns a ChargeStateCache entity by its id.
func (c *ChargeStateCacheClient) Get(ctx context.Context, id int) (*ChargeStateCache, error) {
	return c.Query().Where(chargestatecache.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChargeStateCacheClient) GetX(ctx context.Context, id int) *ChargeStateCache {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChargeStateCacheClient) Hooks() []Hook {
	return c.hooks.ChargeStateCache
}

// Interceptors returns the client interceptors.
func (c *ChargeStateCacheClient) Interceptors() []Interceptor {
	return c.inters.ChargeStateCache
}

func (c *ChargeStateCacheClient) mutate(ctx context.Context, m *ChargeStateCacheMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChargeStateCacheCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChargeStateCacheUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChargeStateCacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChargeStateCacheDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChargeStateCache mutation op: %q", m.Op())
	}
}

// GrantClient is a client for the Grant schema.
type GrantClient struct {
	config
}

// NewGrantClient returns a client for the Grant from the given config.
func NewGrantClient(c config) *GrantClient {
	return &GrantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `grant.Hooks(f(g(h())))`.
func (c *GrantClient) Use(hooks ...Hook) {
	c.hooks.Grant = append(c.hooks.Grant, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `grant.Intercept(f(g(h())))`.
func (c *GrantClient) Intercept(interceptors ...Interceptor) {
	c.inters.Grant = append(c.inters.Grant, interceptors...)
}

// Create returns a builder for creating a Grant entity.
func (c *GrantClient) Create() *GrantCreate {
	mutation := newGrantMutation(c.config, OpCreate)
	return &GrantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Grant entities.
func (c *GrantClient) CreateBulk(builders ...*GrantCreate) *GrantCreateBulk {
	return &GrantCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *GrantClient) MapCreateBulk(slice any, setFunc func(*GrantCreate, int)) *GrantCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &GrantCreateBulk{err: fmt.Errorf("calling to GrantClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*GrantCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &GrantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Grant.
func (c *GrantClient) Update() *GrantUpdate {
	mutation := newGrantMutation(c.config, OpUpdate)
	return &GrantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GrantClient) UpdateOne(gr *Grant) *GrantUpdateOne {
	mutation := newGrantMutation(c.config, OpUpdateOne, withGrant(gr))
	return &GrantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GrantClient) UpdateOneID(id int) *GrantUpdateOne {
	mutation := newGrantMutation(c.config, OpUpdateOne, withGrantID(id))
	return &GrantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Grant.
func (c *GrantClient) Delete() *GrantDelete {
	mutation := newGrantMutation(c.config, OpDelete)
	return &GrantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GrantClient) DeleteOne(gr *Grant) *GrantDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GrantClient) DeleteOneID(id int) *GrantDeleteOne {
	builder := c.Delete().Where(grant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GrantDeleteOne{builder}
}

// Query returns a query builder for Grant.
func (c *GrantClient) Query() *GrantQuery {
	return &GrantQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGrant},
		inters: c.Interceptors(),
	}
}

// Get returns a Grant entity by its id.
func (c *GrantClient) Get(ctx context.Context, id int) (*Grant, error) {
	return c.Query().Where(grant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GrantClient) GetX(ctx context.Context, id int) *Grant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GrantClient) Hooks() []Hook {
	return c.hooks.Grant
}

// Interceptors returns the client interceptors.
func (c *GrantClient) Interceptors() []Interceptor {
	return c.inters.Grant
}

func (c *GrantClient) mutate(ctx context.Context, m *GrantMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GrantCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GrantUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GrantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GrantDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Grant mutation op: %q", m.Op())
	}
}

// PowerMetricClient is a client for the PowerMetric schema.
type PowerMetricClient struct {
	config
}

// NewPowerMetricClient returns a client for the PowerMetric from the given config.
func NewPowerMetricClient(c config) *PowerMetricClient {
	return &PowerMetricClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `powermetric.Hooks(f(g(h())))`.
func (c *PowerMetricClient) Use(hooks ...Hook) {
	c.hooks.PowerMetric = append(c.hooks.PowerMetric, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `powermetric.Intercept(f(g(h())))`.
func (c *PowerMetricClient) Intercept(interceptors ...Interceptor) {
	c.inters.PowerMetric = append(c.inters.PowerMetric, interceptors...)
}

// Create returns a builder for creating a PowerMetric entity.
func (c *PowerMetricClient) Create() *PowerMetricCreate {
	mutation := newPowerMetricMutation(c.config, OpCreate)
	return &PowerMetricCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PowerMetric entities.
func (c *PowerMetricClient) CreateBulk(builders ...*PowerMetricCreate) *PowerMetricCreateBulk {
	return &PowerMetricCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PowerMetricClient) MapCreateBulk(slice any, setFunc func(*PowerMetricCreate, int)) *PowerMetricCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PowerMetricCreateBulk{err: fmt.Errorf("calling to PowerMetricClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PowerMetricCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PowerMetricCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PowerMetric.
func (c *PowerMetricClient) Update() *PowerMetricUpdate {
	mutation := newPowerMetricMutation(c.config, OpUpdate)
	return &PowerMetricUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PowerMetricClient) UpdateOne(pm *PowerMetric) *PowerMetricUpdateOne {
	mutation := newPowerMetricMutation(c.config, OpUpdateOne, withPowerMetric(pm))
	return &PowerMetricUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PowerMetricClient) UpdateOneID(id int) *PowerMetricUpdateOne {
	mutation := newPowerMetricMutation(c.config, OpUpdateOne, withPowerMetricID(id))
	return &PowerMetricUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PowerMetric.
func (c *PowerMetricClient) Delete() *PowerMetricDelete {
	mutation := newPowerMetricMutation(c.config, OpDelete)
	return &PowerMetricDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PowerMetricClient) DeleteOne(pm *PowerMetric) *PowerMetricDeleteOne {
	return c.DeleteOneID(pm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PowerMetricClient) DeleteOneID(id int) *PowerMetricDeleteOne {
	builder := c.Delete().Where(powermetric.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PowerMetricDeleteOne{builder}
}

// Query returns a query builder for PowerMetric.
func (c *PowerMetricClient) Query() *PowerMetricQuery {
	return &PowerMetricQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePowerMetric},
		inters: c.Interceptors(),
	}
}

// Get returns a PowerMetric entity by its id.
func (c *PowerMetricClient) Get(ctx context.Context, id int) (*PowerMetric, error) {
	return c.Query().Where(powermetric.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PowerMetricClient) GetX(ctx context.Context, id int) *PowerMetric {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PowerMetricClient) Hooks() []Hook {
	return c.hooks.PowerMetric
}

// Interceptors returns the client interceptors.
func (c *PowerMetricClient) Interceptors() []Interceptor {
	return c.inters.PowerMetric
}

func (c *PowerMetricClient) mutate(ctx context.Context, m *PowerMetricMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PowerMetricCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PowerMetricUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PowerMetricUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PowerMetricDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PowerMetric mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ChargeCommandHistory, ChargeSetting, ChargeStateCache, Grant,
		PowerMetric []ent.Hook
	}
	inters struct {
		ChargeCommandHistory, ChargeSetting, ChargeStateCache, Grant,
		PowerMetric []ent.Interceptor
	}
)
