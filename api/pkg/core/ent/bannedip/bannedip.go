// Code generated by ent, DO NOT EDIT.

package bannedip

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bannedip type in the database.
	Label = "banned_ip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// Table holds the table name of the bannedip in the database.
	Table = "banned_ips"
)

// Columns holds all SQL columns for bannedip fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldIP,
	FieldCountry,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// IPValidator is a validator for the "ip" field. It is called by the builders before save.
	IPValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
)

// OrderOption defines the ordering options for the BannedIP queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}
