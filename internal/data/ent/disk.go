// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/menta2l/go-hwc/internal/data/ent/disk"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
)

// Disk is the model entity for the Disk schema.
type Disk struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Device holds the value of the "device" field.
	Device string `json:"device,omitempty"`
	// Mountpoint holds the value of the "Mountpoint" field.
	Mountpoint string `json:"Mountpoint,omitempty"`
	// Fstype holds the value of the "Fstype" field.
	Fstype string `json:"Fstype,omitempty"`
	// Opts holds the value of the "opts" field.
	Opts []string `json:"opts,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiskQuery when eager-loading is set.
	Edges     DiskEdges `json:"edges"`
	host_disk *string
}

// DiskEdges holds the relations/edges for other nodes in the graph.
type DiskEdges struct {
	// HostID holds the value of the host_id edge.
	HostID *Host `json:"host_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HostIDOrErr returns the HostID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiskEdges) HostIDOrErr() (*Host, error) {
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
func (*Disk) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case disk.FieldOpts:
			values[i] = new([]byte)
		case disk.FieldID:
			values[i] = new(sql.NullInt64)
		case disk.FieldDevice, disk.FieldMountpoint, disk.FieldFstype:
			values[i] = new(sql.NullString)
		case disk.FieldCreatedAt, disk.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case disk.ForeignKeys[0]: // host_disk
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Disk", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Disk fields.
func (d *Disk) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case disk.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case disk.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				d.Device = value.String
			}
		case disk.FieldMountpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Mountpoint", values[i])
			} else if value.Valid {
				d.Mountpoint = value.String
			}
		case disk.FieldFstype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Fstype", values[i])
			} else if value.Valid {
				d.Fstype = value.String
			}
		case disk.FieldOpts:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field opts", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.Opts); err != nil {
					return fmt.Errorf("unmarshal field opts: %w", err)
				}
			}
		case disk.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case disk.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case disk.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_disk", values[i])
			} else if value.Valid {
				d.host_disk = new(string)
				*d.host_disk = value.String
			}
		}
	}
	return nil
}

// QueryHostID queries the "host_id" edge of the Disk entity.
func (d *Disk) QueryHostID() *HostQuery {
	return (&DiskClient{config: d.config}).QueryHostID(d)
}

// Update returns a builder for updating this Disk.
// Note that you need to call Disk.Unwrap() before calling this method if this Disk
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Disk) Update() *DiskUpdateOne {
	return (&DiskClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Disk entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Disk) Unwrap() *Disk {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Disk is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Disk) String() string {
	var builder strings.Builder
	builder.WriteString("Disk(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", device=")
	builder.WriteString(d.Device)
	builder.WriteString(", Mountpoint=")
	builder.WriteString(d.Mountpoint)
	builder.WriteString(", Fstype=")
	builder.WriteString(d.Fstype)
	builder.WriteString(", opts=")
	builder.WriteString(fmt.Sprintf("%v", d.Opts))
	builder.WriteString(", created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Disks is a parsable slice of Disk.
type Disks []*Disk

func (d Disks) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
