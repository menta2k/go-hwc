// Code generated by entc, DO NOT EDIT.

package network

import (
	"time"
)

const (
	// Label holds the string label denoting the network type in the database.
	Label = "network"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "idx"
	// FieldMTU holds the string denoting the mtu field in the database.
	FieldMTU = "mtu"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHardwareAddr holds the string denoting the hardwareaddr field in the database.
	FieldHardwareAddr = "mac"
	// FieldFlags holds the string denoting the flags field in the database.
	FieldFlags = "flags"
	// FieldAddrs holds the string denoting the addrs field in the database.
	FieldAddrs = "addrs"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeHostID holds the string denoting the host_id edge name in mutations.
	EdgeHostID = "host_id"
	// Table holds the table name of the network in the database.
	Table = "networks"
	// HostIDTable is the table that holds the host_id relation/edge.
	HostIDTable = "networks"
	// HostIDInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostIDInverseTable = "hosts"
	// HostIDColumn is the table column denoting the host_id relation/edge.
	HostIDColumn = "host_network"
)

// Columns holds all SQL columns for network fields.
var Columns = []string{
	FieldID,
	FieldIndex,
	FieldMTU,
	FieldName,
	FieldHardwareAddr,
	FieldFlags,
	FieldAddrs,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "networks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"host_network",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
