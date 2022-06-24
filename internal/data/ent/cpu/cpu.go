// Code generated by entc, DO NOT EDIT.

package cpu

import (
	"time"
)

const (
	// Label holds the string label denoting the cpu type in the database.
	Label = "cpu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCPU holds the string denoting the cpu field in the database.
	FieldCPU = "idx"
	// FieldVendorID holds the string denoting the vendor_id field in the database.
	FieldVendorID = "vendor_id"
	// FieldFamily holds the string denoting the family field in the database.
	FieldFamily = "family"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldModelName holds the string denoting the model_name field in the database.
	FieldModelName = "model_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeHostID holds the string denoting the host_id edge name in mutations.
	EdgeHostID = "host_id"
	// Table holds the table name of the cpu in the database.
	Table = "cpus"
	// HostIDTable is the table that holds the host_id relation/edge.
	HostIDTable = "cpus"
	// HostIDInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostIDInverseTable = "hosts"
	// HostIDColumn is the table column denoting the host_id relation/edge.
	HostIDColumn = "host_cpu"
)

// Columns holds all SQL columns for cpu fields.
var Columns = []string{
	FieldID,
	FieldCPU,
	FieldVendorID,
	FieldFamily,
	FieldModel,
	FieldModelName,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cpus"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"host_cpu",
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
