// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/menta2l/go-hwc/internal/data/ent/cpu"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
)

// CPUCreate is the builder for creating a Cpu entity.
type CPUCreate struct {
	config
	mutation *CPUMutation
	hooks    []Hook
}

// SetVendorID sets the "vendor_id" field.
func (cc *CPUCreate) SetVendorID(s string) *CPUCreate {
	cc.mutation.SetVendorID(s)
	return cc
}

// SetFamily sets the "family" field.
func (cc *CPUCreate) SetFamily(s string) *CPUCreate {
	cc.mutation.SetFamily(s)
	return cc
}

// SetModel sets the "model" field.
func (cc *CPUCreate) SetModel(s string) *CPUCreate {
	cc.mutation.SetModel(s)
	return cc
}

// SetModelName sets the "model_name" field.
func (cc *CPUCreate) SetModelName(s string) *CPUCreate {
	cc.mutation.SetModelName(s)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CPUCreate) SetCreatedAt(t time.Time) *CPUCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CPUCreate) SetNillableCreatedAt(t *time.Time) *CPUCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CPUCreate) SetUpdatedAt(t time.Time) *CPUCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CPUCreate) SetNillableUpdatedAt(t *time.Time) *CPUCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetHostIDID sets the "host_id" edge to the Host entity by ID.
func (cc *CPUCreate) SetHostIDID(id string) *CPUCreate {
	cc.mutation.SetHostIDID(id)
	return cc
}

// SetNillableHostIDID sets the "host_id" edge to the Host entity by ID if the given value is not nil.
func (cc *CPUCreate) SetNillableHostIDID(id *string) *CPUCreate {
	if id != nil {
		cc = cc.SetHostIDID(*id)
	}
	return cc
}

// SetHostID sets the "host_id" edge to the Host entity.
func (cc *CPUCreate) SetHostID(h *Host) *CPUCreate {
	return cc.SetHostIDID(h.ID)
}

// Mutation returns the CPUMutation object of the builder.
func (cc *CPUCreate) Mutation() *CPUMutation {
	return cc.mutation
}

// Save creates the Cpu in the database.
func (cc *CPUCreate) Save(ctx context.Context) (*Cpu, error) {
	var (
		err  error
		node *Cpu
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CPUMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CPUCreate) SaveX(ctx context.Context) *Cpu {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CPUCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CPUCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CPUCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := cpu.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := cpu.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CPUCreate) check() error {
	if _, ok := cc.mutation.VendorID(); !ok {
		return &ValidationError{Name: "vendor_id", err: errors.New(`ent: missing required field "Cpu.vendor_id"`)}
	}
	if _, ok := cc.mutation.Family(); !ok {
		return &ValidationError{Name: "family", err: errors.New(`ent: missing required field "Cpu.family"`)}
	}
	if _, ok := cc.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Cpu.model"`)}
	}
	if _, ok := cc.mutation.ModelName(); !ok {
		return &ValidationError{Name: "model_name", err: errors.New(`ent: missing required field "Cpu.model_name"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Cpu.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Cpu.updated_at"`)}
	}
	return nil
}

func (cc *CPUCreate) sqlSave(ctx context.Context) (*Cpu, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CPUCreate) createSpec() (*Cpu, *sqlgraph.CreateSpec) {
	var (
		_node = &Cpu{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cpu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cpu.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.VendorID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cpu.FieldVendorID,
		})
		_node.VendorID = value
	}
	if value, ok := cc.mutation.Family(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cpu.FieldFamily,
		})
		_node.Family = value
	}
	if value, ok := cc.mutation.Model(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cpu.FieldModel,
		})
		_node.Model = value
	}
	if value, ok := cc.mutation.ModelName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cpu.FieldModelName,
		})
		_node.ModelName = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cpu.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cpu.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.HostIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cpu.HostIDTable,
			Columns: []string{cpu.HostIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.host_cpu_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CPUCreateBulk is the builder for creating many Cpu entities in bulk.
type CPUCreateBulk struct {
	config
	builders []*CPUCreate
}

// Save creates the Cpu entities in the database.
func (ccb *CPUCreateBulk) Save(ctx context.Context) ([]*Cpu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cpu, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CPUMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CPUCreateBulk) SaveX(ctx context.Context) []*Cpu {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CPUCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CPUCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
