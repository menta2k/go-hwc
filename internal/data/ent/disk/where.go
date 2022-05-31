// Code generated by entc, DO NOT EDIT.

package disk

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jackc/pgtype"
	"github.com/menta2l/go-hwc/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Device applies equality check predicate on the "device" field. It's identical to DeviceEQ.
func Device(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDevice), v))
	})
}

// Mount applies equality check predicate on the "mount" field. It's identical to MountEQ.
func Mount(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMount), v))
	})
}

// FsType applies equality check predicate on the "fs_type" field. It's identical to FsTypeEQ.
func FsType(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFsType), v))
	})
}

// Opts applies equality check predicate on the "opts" field. It's identical to OptsEQ.
func Opts(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpts), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeviceEQ applies the EQ predicate on the "device" field.
func DeviceEQ(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDevice), v))
	})
}

// DeviceNEQ applies the NEQ predicate on the "device" field.
func DeviceNEQ(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDevice), v))
	})
}

// DeviceIn applies the In predicate on the "device" field.
func DeviceIn(vs ...string) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDevice), v...))
	})
}

// DeviceNotIn applies the NotIn predicate on the "device" field.
func DeviceNotIn(vs ...string) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDevice), v...))
	})
}

// DeviceGT applies the GT predicate on the "device" field.
func DeviceGT(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDevice), v))
	})
}

// DeviceGTE applies the GTE predicate on the "device" field.
func DeviceGTE(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDevice), v))
	})
}

// DeviceLT applies the LT predicate on the "device" field.
func DeviceLT(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDevice), v))
	})
}

// DeviceLTE applies the LTE predicate on the "device" field.
func DeviceLTE(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDevice), v))
	})
}

// DeviceContains applies the Contains predicate on the "device" field.
func DeviceContains(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDevice), v))
	})
}

// DeviceHasPrefix applies the HasPrefix predicate on the "device" field.
func DeviceHasPrefix(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDevice), v))
	})
}

// DeviceHasSuffix applies the HasSuffix predicate on the "device" field.
func DeviceHasSuffix(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDevice), v))
	})
}

// DeviceEqualFold applies the EqualFold predicate on the "device" field.
func DeviceEqualFold(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDevice), v))
	})
}

// DeviceContainsFold applies the ContainsFold predicate on the "device" field.
func DeviceContainsFold(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDevice), v))
	})
}

// MountEQ applies the EQ predicate on the "mount" field.
func MountEQ(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMount), v))
	})
}

// MountNEQ applies the NEQ predicate on the "mount" field.
func MountNEQ(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMount), v))
	})
}

// MountIn applies the In predicate on the "mount" field.
func MountIn(vs ...string) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMount), v...))
	})
}

// MountNotIn applies the NotIn predicate on the "mount" field.
func MountNotIn(vs ...string) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMount), v...))
	})
}

// MountGT applies the GT predicate on the "mount" field.
func MountGT(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMount), v))
	})
}

// MountGTE applies the GTE predicate on the "mount" field.
func MountGTE(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMount), v))
	})
}

// MountLT applies the LT predicate on the "mount" field.
func MountLT(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMount), v))
	})
}

// MountLTE applies the LTE predicate on the "mount" field.
func MountLTE(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMount), v))
	})
}

// MountContains applies the Contains predicate on the "mount" field.
func MountContains(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMount), v))
	})
}

// MountHasPrefix applies the HasPrefix predicate on the "mount" field.
func MountHasPrefix(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMount), v))
	})
}

// MountHasSuffix applies the HasSuffix predicate on the "mount" field.
func MountHasSuffix(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMount), v))
	})
}

// MountEqualFold applies the EqualFold predicate on the "mount" field.
func MountEqualFold(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMount), v))
	})
}

// MountContainsFold applies the ContainsFold predicate on the "mount" field.
func MountContainsFold(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMount), v))
	})
}

// FsTypeEQ applies the EQ predicate on the "fs_type" field.
func FsTypeEQ(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFsType), v))
	})
}

// FsTypeNEQ applies the NEQ predicate on the "fs_type" field.
func FsTypeNEQ(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFsType), v))
	})
}

// FsTypeIn applies the In predicate on the "fs_type" field.
func FsTypeIn(vs ...string) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFsType), v...))
	})
}

// FsTypeNotIn applies the NotIn predicate on the "fs_type" field.
func FsTypeNotIn(vs ...string) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFsType), v...))
	})
}

// FsTypeGT applies the GT predicate on the "fs_type" field.
func FsTypeGT(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFsType), v))
	})
}

// FsTypeGTE applies the GTE predicate on the "fs_type" field.
func FsTypeGTE(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFsType), v))
	})
}

// FsTypeLT applies the LT predicate on the "fs_type" field.
func FsTypeLT(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFsType), v))
	})
}

// FsTypeLTE applies the LTE predicate on the "fs_type" field.
func FsTypeLTE(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFsType), v))
	})
}

// FsTypeContains applies the Contains predicate on the "fs_type" field.
func FsTypeContains(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFsType), v))
	})
}

// FsTypeHasPrefix applies the HasPrefix predicate on the "fs_type" field.
func FsTypeHasPrefix(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFsType), v))
	})
}

// FsTypeHasSuffix applies the HasSuffix predicate on the "fs_type" field.
func FsTypeHasSuffix(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFsType), v))
	})
}

// FsTypeEqualFold applies the EqualFold predicate on the "fs_type" field.
func FsTypeEqualFold(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFsType), v))
	})
}

// FsTypeContainsFold applies the ContainsFold predicate on the "fs_type" field.
func FsTypeContainsFold(v string) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFsType), v))
	})
}

// OptsEQ applies the EQ predicate on the "opts" field.
func OptsEQ(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpts), v))
	})
}

// OptsNEQ applies the NEQ predicate on the "opts" field.
func OptsNEQ(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOpts), v))
	})
}

// OptsIn applies the In predicate on the "opts" field.
func OptsIn(vs ...*pgtype.TextArray) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOpts), v...))
	})
}

// OptsNotIn applies the NotIn predicate on the "opts" field.
func OptsNotIn(vs ...*pgtype.TextArray) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOpts), v...))
	})
}

// OptsGT applies the GT predicate on the "opts" field.
func OptsGT(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOpts), v))
	})
}

// OptsGTE applies the GTE predicate on the "opts" field.
func OptsGTE(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOpts), v))
	})
}

// OptsLT applies the LT predicate on the "opts" field.
func OptsLT(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOpts), v))
	})
}

// OptsLTE applies the LTE predicate on the "opts" field.
func OptsLTE(v *pgtype.TextArray) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOpts), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Disk {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disk(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasHostID applies the HasEdge predicate on the "host_id" edge.
func HasHostID() predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostIDTable, HostIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostIDWith applies the HasEdge predicate on the "host_id" edge with a given conditions (other predicates).
func HasHostIDWith(preds ...predicate.Host) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostIDTable, HostIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Disk) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Disk) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Disk) predicate.Disk {
	return predicate.Disk(func(s *sql.Selector) {
		p(s.Not())
	})
}
