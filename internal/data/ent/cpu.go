// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/menta2l/go-hwc/internal/data/ent/cpu"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
)

// Cpu is the model entity for the Cpu schema.
type Cpu struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VendorID holds the value of the "vendor_id" field.
	VendorID string `json:"vendor_id,omitempty"`
	// Family holds the value of the "family" field.
	Family string `json:"family,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// ModelName holds the value of the "model_name" field.
	ModelName string `json:"model_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CpuQuery when eager-loading is set.
	Edges       CpuEdges `json:"edges"`
	host_cpu_id *string
}

// CpuEdges holds the relations/edges for other nodes in the graph.
type CpuEdges struct {
	// HostID holds the value of the host_id edge.
	HostID *Host `json:"host_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HostIDOrErr returns the HostID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CpuEdges) HostIDOrErr() (*Host, error) {
	if e.loadedTypes[0] {
		if e.HostID == nil {
			// The edge host_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.HostID, nil
	}
	return nil, &NotLoadedError{edge: "host_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cpu) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cpu.FieldID:
			values[i] = new(sql.NullInt64)
		case cpu.FieldVendorID, cpu.FieldFamily, cpu.FieldModel, cpu.FieldModelName:
			values[i] = new(sql.NullString)
		case cpu.FieldCreatedAt, cpu.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case cpu.ForeignKeys[0]: // host_cpu_id
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cpu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cpu fields.
func (c *Cpu) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cpu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cpu.FieldVendorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vendor_id", values[i])
			} else if value.Valid {
				c.VendorID = value.String
			}
		case cpu.FieldFamily:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family", values[i])
			} else if value.Valid {
				c.Family = value.String
			}
		case cpu.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				c.Model = value.String
			}
		case cpu.FieldModelName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model_name", values[i])
			} else if value.Valid {
				c.ModelName = value.String
			}
		case cpu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cpu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case cpu.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_cpu_id", values[i])
			} else if value.Valid {
				c.host_cpu_id = new(string)
				*c.host_cpu_id = value.String
			}
		}
	}
	return nil
}

// QueryHostID queries the "host_id" edge of the Cpu entity.
func (c *Cpu) QueryHostID() *HostQuery {
	return (&CpuClient{config: c.config}).QueryHostID(c)
}

// Update returns a builder for updating this Cpu.
// Note that you need to call Cpu.Unwrap() before calling this method if this Cpu
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cpu) Update() *CPUUpdateOne {
	return (&CpuClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cpu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cpu) Unwrap() *Cpu {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cpu is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cpu) String() string {
	var builder strings.Builder
	builder.WriteString("Cpu(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", vendor_id=")
	builder.WriteString(c.VendorID)
	builder.WriteString(", family=")
	builder.WriteString(c.Family)
	builder.WriteString(", model=")
	builder.WriteString(c.Model)
	builder.WriteString(", model_name=")
	builder.WriteString(c.ModelName)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Cpus is a parsable slice of Cpu.
type Cpus []*Cpu

func (c Cpus) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
