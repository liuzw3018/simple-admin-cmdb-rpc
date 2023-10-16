// Code generated by ent, DO NOT EDIT.

package cmdbipuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cmdbipuser type in the database.
	Label = "cmdb_ip_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldDepartment holds the string denoting the department field in the database.
	FieldDepartment = "department"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// Table holds the table name of the cmdbipuser in the database.
	Table = "cmdb_ip_users"
)

// Columns holds all SQL columns for cmdbipuser fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUser,
	FieldDepartment,
	FieldMobile,
	FieldRemark,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	MobileValidator func(string) error
)

// OrderOption defines the ordering options for the CmdbIpUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUser orders the results by the user field.
func ByUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUser, opts...).ToFunc()
}

// ByDepartment orders the results by the department field.
func ByDepartment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartment, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}
