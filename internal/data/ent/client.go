// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/menta2l/go-hwc/internal/data/ent/migrate"

	"github.com/menta2l/go-hwc/internal/data/ent/cpu"
	"github.com/menta2l/go-hwc/internal/data/ent/disk"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
	"github.com/menta2l/go-hwc/internal/data/ent/netstat"
	"github.com/menta2l/go-hwc/internal/data/ent/network"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cpu is the client for interacting with the Cpu builders.
	Cpu *CpuClient
	// Disk is the client for interacting with the Disk builders.
	Disk *DiskClient
	// Host is the client for interacting with the Host builders.
	Host *HostClient
	// Netstat is the client for interacting with the Netstat builders.
	Netstat *NetstatClient
	// Network is the client for interacting with the Network builders.
	Network *NetworkClient
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
	c.Cpu = NewCpuClient(c.config)
	c.Disk = NewDiskClient(c.config)
	c.Host = NewHostClient(c.config)
	c.Netstat = NewNetstatClient(c.config)
	c.Network = NewNetworkClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Cpu:     NewCpuClient(cfg),
		Disk:    NewDiskClient(cfg),
		Host:    NewHostClient(cfg),
		Netstat: NewNetstatClient(cfg),
		Network: NewNetworkClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:     ctx,
		config:  cfg,
		Cpu:     NewCpuClient(cfg),
		Disk:    NewDiskClient(cfg),
		Host:    NewHostClient(cfg),
		Netstat: NewNetstatClient(cfg),
		Network: NewNetworkClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cpu.
//		Query().
//		Count(ctx)
//
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
	c.Cpu.Use(hooks...)
	c.Disk.Use(hooks...)
	c.Host.Use(hooks...)
	c.Netstat.Use(hooks...)
	c.Network.Use(hooks...)
}

// CpuClient is a client for the Cpu schema.
type CpuClient struct {
	config
}

// NewCpuClient returns a client for the Cpu from the given config.
func NewCpuClient(c config) *CpuClient {
	return &CpuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cpu.Hooks(f(g(h())))`.
func (c *CpuClient) Use(hooks ...Hook) {
	c.hooks.Cpu = append(c.hooks.Cpu, hooks...)
}

// Create returns a create builder for Cpu.
func (c *CpuClient) Create() *CPUCreate {
	mutation := newCPUMutation(c.config, OpCreate)
	return &CPUCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cpu entities.
func (c *CpuClient) CreateBulk(builders ...*CPUCreate) *CPUCreateBulk {
	return &CPUCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cpu.
func (c *CpuClient) Update() *CPUUpdate {
	mutation := newCPUMutation(c.config, OpUpdate)
	return &CPUUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CpuClient) UpdateOne(cp *Cpu) *CPUUpdateOne {
	mutation := newCPUMutation(c.config, OpUpdateOne, withCpu(cp))
	return &CPUUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CpuClient) UpdateOneID(id int) *CPUUpdateOne {
	mutation := newCPUMutation(c.config, OpUpdateOne, withCpuID(id))
	return &CPUUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cpu.
func (c *CpuClient) Delete() *CPUDelete {
	mutation := newCPUMutation(c.config, OpDelete)
	return &CPUDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CpuClient) DeleteOne(cp *Cpu) *CPUDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CpuClient) DeleteOneID(id int) *CPUDeleteOne {
	builder := c.Delete().Where(cpu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CPUDeleteOne{builder}
}

// Query returns a query builder for Cpu.
func (c *CpuClient) Query() *CPUQuery {
	return &CPUQuery{
		config: c.config,
	}
}

// Get returns a Cpu entity by its id.
func (c *CpuClient) Get(ctx context.Context, id int) (*Cpu, error) {
	return c.Query().Where(cpu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CpuClient) GetX(ctx context.Context, id int) *Cpu {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHostID queries the host_id edge of a Cpu.
func (c *CpuClient) QueryHostID(cp *Cpu) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cpu.Table, cpu.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cpu.HostIDTable, cpu.HostIDColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CpuClient) Hooks() []Hook {
	return c.hooks.Cpu
}

// DiskClient is a client for the Disk schema.
type DiskClient struct {
	config
}

// NewDiskClient returns a client for the Disk from the given config.
func NewDiskClient(c config) *DiskClient {
	return &DiskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `disk.Hooks(f(g(h())))`.
func (c *DiskClient) Use(hooks ...Hook) {
	c.hooks.Disk = append(c.hooks.Disk, hooks...)
}

// Create returns a create builder for Disk.
func (c *DiskClient) Create() *DiskCreate {
	mutation := newDiskMutation(c.config, OpCreate)
	return &DiskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Disk entities.
func (c *DiskClient) CreateBulk(builders ...*DiskCreate) *DiskCreateBulk {
	return &DiskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Disk.
func (c *DiskClient) Update() *DiskUpdate {
	mutation := newDiskMutation(c.config, OpUpdate)
	return &DiskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DiskClient) UpdateOne(d *Disk) *DiskUpdateOne {
	mutation := newDiskMutation(c.config, OpUpdateOne, withDisk(d))
	return &DiskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DiskClient) UpdateOneID(id int) *DiskUpdateOne {
	mutation := newDiskMutation(c.config, OpUpdateOne, withDiskID(id))
	return &DiskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Disk.
func (c *DiskClient) Delete() *DiskDelete {
	mutation := newDiskMutation(c.config, OpDelete)
	return &DiskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DiskClient) DeleteOne(d *Disk) *DiskDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DiskClient) DeleteOneID(id int) *DiskDeleteOne {
	builder := c.Delete().Where(disk.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DiskDeleteOne{builder}
}

// Query returns a query builder for Disk.
func (c *DiskClient) Query() *DiskQuery {
	return &DiskQuery{
		config: c.config,
	}
}

// Get returns a Disk entity by its id.
func (c *DiskClient) Get(ctx context.Context, id int) (*Disk, error) {
	return c.Query().Where(disk.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DiskClient) GetX(ctx context.Context, id int) *Disk {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHostID queries the host_id edge of a Disk.
func (c *DiskClient) QueryHostID(d *Disk) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(disk.Table, disk.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, disk.HostIDTable, disk.HostIDColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DiskClient) Hooks() []Hook {
	return c.hooks.Disk
}

// HostClient is a client for the Host schema.
type HostClient struct {
	config
}

// NewHostClient returns a client for the Host from the given config.
func NewHostClient(c config) *HostClient {
	return &HostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `host.Hooks(f(g(h())))`.
func (c *HostClient) Use(hooks ...Hook) {
	c.hooks.Host = append(c.hooks.Host, hooks...)
}

// Create returns a create builder for Host.
func (c *HostClient) Create() *HostCreate {
	mutation := newHostMutation(c.config, OpCreate)
	return &HostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Host entities.
func (c *HostClient) CreateBulk(builders ...*HostCreate) *HostCreateBulk {
	return &HostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Host.
func (c *HostClient) Update() *HostUpdate {
	mutation := newHostMutation(c.config, OpUpdate)
	return &HostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HostClient) UpdateOne(h *Host) *HostUpdateOne {
	mutation := newHostMutation(c.config, OpUpdateOne, withHost(h))
	return &HostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HostClient) UpdateOneID(id string) *HostUpdateOne {
	mutation := newHostMutation(c.config, OpUpdateOne, withHostID(id))
	return &HostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Host.
func (c *HostClient) Delete() *HostDelete {
	mutation := newHostMutation(c.config, OpDelete)
	return &HostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HostClient) DeleteOne(h *Host) *HostDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HostClient) DeleteOneID(id string) *HostDeleteOne {
	builder := c.Delete().Where(host.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HostDeleteOne{builder}
}

// Query returns a query builder for Host.
func (c *HostClient) Query() *HostQuery {
	return &HostQuery{
		config: c.config,
	}
}

// Get returns a Host entity by its id.
func (c *HostClient) Get(ctx context.Context, id string) (*Host, error) {
	return c.Query().Where(host.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HostClient) GetX(ctx context.Context, id string) *Host {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCPU queries the cpu edge of a Host.
func (c *HostClient) QueryCPU(h *Host) *CPUQuery {
	query := &CPUQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(cpu.Table, cpu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.CPUTable, host.CPUColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNetwork queries the network edge of a Host.
func (c *HostClient) QueryNetwork(h *Host) *NetworkQuery {
	query := &NetworkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(network.Table, network.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.NetworkTable, host.NetworkColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNetstat queries the netstat edge of a Host.
func (c *HostClient) QueryNetstat(h *Host) *NetstatQuery {
	query := &NetstatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(netstat.Table, netstat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.NetstatTable, host.NetstatColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDisk queries the disk edge of a Host.
func (c *HostClient) QueryDisk(h *Host) *DiskQuery {
	query := &DiskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(disk.Table, disk.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.DiskTable, host.DiskColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HostClient) Hooks() []Hook {
	return c.hooks.Host
}

// NetstatClient is a client for the Netstat schema.
type NetstatClient struct {
	config
}

// NewNetstatClient returns a client for the Netstat from the given config.
func NewNetstatClient(c config) *NetstatClient {
	return &NetstatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `netstat.Hooks(f(g(h())))`.
func (c *NetstatClient) Use(hooks ...Hook) {
	c.hooks.Netstat = append(c.hooks.Netstat, hooks...)
}

// Create returns a create builder for Netstat.
func (c *NetstatClient) Create() *NetstatCreate {
	mutation := newNetstatMutation(c.config, OpCreate)
	return &NetstatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Netstat entities.
func (c *NetstatClient) CreateBulk(builders ...*NetstatCreate) *NetstatCreateBulk {
	return &NetstatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Netstat.
func (c *NetstatClient) Update() *NetstatUpdate {
	mutation := newNetstatMutation(c.config, OpUpdate)
	return &NetstatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NetstatClient) UpdateOne(n *Netstat) *NetstatUpdateOne {
	mutation := newNetstatMutation(c.config, OpUpdateOne, withNetstat(n))
	return &NetstatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NetstatClient) UpdateOneID(id int) *NetstatUpdateOne {
	mutation := newNetstatMutation(c.config, OpUpdateOne, withNetstatID(id))
	return &NetstatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Netstat.
func (c *NetstatClient) Delete() *NetstatDelete {
	mutation := newNetstatMutation(c.config, OpDelete)
	return &NetstatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NetstatClient) DeleteOne(n *Netstat) *NetstatDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NetstatClient) DeleteOneID(id int) *NetstatDeleteOne {
	builder := c.Delete().Where(netstat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NetstatDeleteOne{builder}
}

// Query returns a query builder for Netstat.
func (c *NetstatClient) Query() *NetstatQuery {
	return &NetstatQuery{
		config: c.config,
	}
}

// Get returns a Netstat entity by its id.
func (c *NetstatClient) Get(ctx context.Context, id int) (*Netstat, error) {
	return c.Query().Where(netstat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NetstatClient) GetX(ctx context.Context, id int) *Netstat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHostID queries the host_id edge of a Netstat.
func (c *NetstatClient) QueryHostID(n *Netstat) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(netstat.Table, netstat.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, netstat.HostIDTable, netstat.HostIDColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NetstatClient) Hooks() []Hook {
	return c.hooks.Netstat
}

// NetworkClient is a client for the Network schema.
type NetworkClient struct {
	config
}

// NewNetworkClient returns a client for the Network from the given config.
func NewNetworkClient(c config) *NetworkClient {
	return &NetworkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `network.Hooks(f(g(h())))`.
func (c *NetworkClient) Use(hooks ...Hook) {
	c.hooks.Network = append(c.hooks.Network, hooks...)
}

// Create returns a create builder for Network.
func (c *NetworkClient) Create() *NetworkCreate {
	mutation := newNetworkMutation(c.config, OpCreate)
	return &NetworkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Network entities.
func (c *NetworkClient) CreateBulk(builders ...*NetworkCreate) *NetworkCreateBulk {
	return &NetworkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Network.
func (c *NetworkClient) Update() *NetworkUpdate {
	mutation := newNetworkMutation(c.config, OpUpdate)
	return &NetworkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NetworkClient) UpdateOne(n *Network) *NetworkUpdateOne {
	mutation := newNetworkMutation(c.config, OpUpdateOne, withNetwork(n))
	return &NetworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NetworkClient) UpdateOneID(id int) *NetworkUpdateOne {
	mutation := newNetworkMutation(c.config, OpUpdateOne, withNetworkID(id))
	return &NetworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Network.
func (c *NetworkClient) Delete() *NetworkDelete {
	mutation := newNetworkMutation(c.config, OpDelete)
	return &NetworkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NetworkClient) DeleteOne(n *Network) *NetworkDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NetworkClient) DeleteOneID(id int) *NetworkDeleteOne {
	builder := c.Delete().Where(network.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NetworkDeleteOne{builder}
}

// Query returns a query builder for Network.
func (c *NetworkClient) Query() *NetworkQuery {
	return &NetworkQuery{
		config: c.config,
	}
}

// Get returns a Network entity by its id.
func (c *NetworkClient) Get(ctx context.Context, id int) (*Network, error) {
	return c.Query().Where(network.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NetworkClient) GetX(ctx context.Context, id int) *Network {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHostID queries the host_id edge of a Network.
func (c *NetworkClient) QueryHostID(n *Network) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(network.Table, network.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, network.HostIDTable, network.HostIDColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NetworkClient) Hooks() []Hook {
	return c.hooks.Network
}
