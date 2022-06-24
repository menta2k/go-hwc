// Code generated by entc, DO NOT EDIT.

package host

import (
	"time"
)

const (
	// Label holds the string label denoting the host type in the database.
	Label = "host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldPlatformFamily holds the string denoting the platform_family field in the database.
	FieldPlatformFamily = "platform_family"
	// FieldPlatformVersion holds the string denoting the platform_version field in the database.
	FieldPlatformVersion = "platform_version"
	// FieldKernelVersion holds the string denoting the kernel_version field in the database.
	FieldKernelVersion = "kernel_version"
	// FieldKernelArch holds the string denoting the kernel_arch field in the database.
	FieldKernelArch = "kernel_arch"
	// FieldVirtualizationSystem holds the string denoting the virtualization_system field in the database.
	FieldVirtualizationSystem = "virtualization_system"
	// FieldVirtualizationRole holds the string denoting the virtualization_role field in the database.
	FieldVirtualizationRole = "virtualization_role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCPU holds the string denoting the cpu edge name in mutations.
	EdgeCPU = "cpu"
	// EdgeNetwork holds the string denoting the network edge name in mutations.
	EdgeNetwork = "network"
	// EdgeNetstat holds the string denoting the netstat edge name in mutations.
	EdgeNetstat = "netstat"
	// EdgeDisk holds the string denoting the disk edge name in mutations.
	EdgeDisk = "disk"
	// Table holds the table name of the host in the database.
	Table = "hosts"
	// CPUTable is the table that holds the cpu relation/edge.
	CPUTable = "cpus"
	// CPUInverseTable is the table name for the Cpu entity.
	// It exists in this package in order to avoid circular dependency with the "cpu" package.
	CPUInverseTable = "cpus"
	// CPUColumn is the table column denoting the cpu relation/edge.
	CPUColumn = "host_cpu"
	// NetworkTable is the table that holds the network relation/edge.
	NetworkTable = "networks"
	// NetworkInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	NetworkInverseTable = "networks"
	// NetworkColumn is the table column denoting the network relation/edge.
	NetworkColumn = "host_network"
	// NetstatTable is the table that holds the netstat relation/edge.
	NetstatTable = "netstats"
	// NetstatInverseTable is the table name for the Netstat entity.
	// It exists in this package in order to avoid circular dependency with the "netstat" package.
	NetstatInverseTable = "netstats"
	// NetstatColumn is the table column denoting the netstat relation/edge.
	NetstatColumn = "host_netstat"
	// DiskTable is the table that holds the disk relation/edge.
	DiskTable = "disks"
	// DiskInverseTable is the table name for the Disk entity.
	// It exists in this package in order to avoid circular dependency with the "disk" package.
	DiskInverseTable = "disks"
	// DiskColumn is the table column denoting the disk relation/edge.
	DiskColumn = "host_disk"
)

// Columns holds all SQL columns for host fields.
var Columns = []string{
	FieldID,
	FieldHostname,
	FieldOs,
	FieldPlatform,
	FieldPlatformFamily,
	FieldPlatformVersion,
	FieldKernelVersion,
	FieldKernelArch,
	FieldVirtualizationSystem,
	FieldVirtualizationRole,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
