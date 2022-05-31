// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/menta2l/go-hwc/internal/data/ent/cpu"
	"github.com/menta2l/go-hwc/internal/data/ent/disk"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
	"github.com/menta2l/go-hwc/internal/data/ent/netstat"
	"github.com/menta2l/go-hwc/internal/data/ent/network"
	"github.com/menta2l/go-hwc/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cpuFields := schema.Cpu{}.Fields()
	_ = cpuFields
	// cpuDescCreatedAt is the schema descriptor for created_at field.
	cpuDescCreatedAt := cpuFields[4].Descriptor()
	// cpu.DefaultCreatedAt holds the default value on creation for the created_at field.
	cpu.DefaultCreatedAt = cpuDescCreatedAt.Default.(func() time.Time)
	// cpuDescUpdatedAt is the schema descriptor for updated_at field.
	cpuDescUpdatedAt := cpuFields[5].Descriptor()
	// cpu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cpu.DefaultUpdatedAt = cpuDescUpdatedAt.Default.(func() time.Time)
	diskFields := schema.Disk{}.Fields()
	_ = diskFields
	// diskDescCreatedAt is the schema descriptor for created_at field.
	diskDescCreatedAt := diskFields[4].Descriptor()
	// disk.DefaultCreatedAt holds the default value on creation for the created_at field.
	disk.DefaultCreatedAt = diskDescCreatedAt.Default.(func() time.Time)
	// diskDescUpdatedAt is the schema descriptor for updated_at field.
	diskDescUpdatedAt := diskFields[5].Descriptor()
	// disk.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	disk.DefaultUpdatedAt = diskDescUpdatedAt.Default.(func() time.Time)
	hostFields := schema.Host{}.Fields()
	_ = hostFields
	// hostDescCreatedAt is the schema descriptor for created_at field.
	hostDescCreatedAt := hostFields[10].Descriptor()
	// host.DefaultCreatedAt holds the default value on creation for the created_at field.
	host.DefaultCreatedAt = hostDescCreatedAt.Default.(func() time.Time)
	// hostDescUpdatedAt is the schema descriptor for updated_at field.
	hostDescUpdatedAt := hostFields[11].Descriptor()
	// host.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	host.DefaultUpdatedAt = hostDescUpdatedAt.Default.(func() time.Time)
	netstatFields := schema.Netstat{}.Fields()
	_ = netstatFields
	// netstatDescCreatedAt is the schema descriptor for created_at field.
	netstatDescCreatedAt := netstatFields[4].Descriptor()
	// netstat.DefaultCreatedAt holds the default value on creation for the created_at field.
	netstat.DefaultCreatedAt = netstatDescCreatedAt.Default.(func() time.Time)
	// netstatDescUpdatedAt is the schema descriptor for updated_at field.
	netstatDescUpdatedAt := netstatFields[5].Descriptor()
	// netstat.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	netstat.DefaultUpdatedAt = netstatDescUpdatedAt.Default.(func() time.Time)
	networkFields := schema.Network{}.Fields()
	_ = networkFields
	// networkDescCreatedAt is the schema descriptor for created_at field.
	networkDescCreatedAt := networkFields[6].Descriptor()
	// network.DefaultCreatedAt holds the default value on creation for the created_at field.
	network.DefaultCreatedAt = networkDescCreatedAt.Default.(func() time.Time)
	// networkDescUpdatedAt is the schema descriptor for updated_at field.
	networkDescUpdatedAt := networkFields[7].Descriptor()
	// network.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	network.DefaultUpdatedAt = networkDescUpdatedAt.Default.(func() time.Time)
}
