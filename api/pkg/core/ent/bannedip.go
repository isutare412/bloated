// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/isutare412/bloated/api/pkg/core/ent/bannedip"
)

// BannedIP is the model entity for the BannedIP schema.
type BannedIP struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Country holds the value of the "country" field.
	Country      string `json:"country,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BannedIP) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bannedip.FieldID:
			values[i] = new(sql.NullInt64)
		case bannedip.FieldIP, bannedip.FieldCountry:
			values[i] = new(sql.NullString)
		case bannedip.FieldCreateTime, bannedip.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BannedIP fields.
func (bi *BannedIP) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bannedip.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bi.ID = int(value.Int64)
		case bannedip.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				bi.CreateTime = value.Time
			}
		case bannedip.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				bi.UpdateTime = value.Time
			}
		case bannedip.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				bi.IP = value.String
			}
		case bannedip.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				bi.Country = value.String
			}
		default:
			bi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BannedIP.
// This includes values selected through modifiers, order, etc.
func (bi *BannedIP) Value(name string) (ent.Value, error) {
	return bi.selectValues.Get(name)
}

// Update returns a builder for updating this BannedIP.
// Note that you need to call BannedIP.Unwrap() before calling this method if this BannedIP
// was returned from a transaction, and the transaction was committed or rolled back.
func (bi *BannedIP) Update() *BannedIPUpdateOne {
	return NewBannedIPClient(bi.config).UpdateOne(bi)
}

// Unwrap unwraps the BannedIP entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bi *BannedIP) Unwrap() *BannedIP {
	_tx, ok := bi.config.driver.(*txDriver)
	if !ok {
		panic("ent: BannedIP is not a transactional entity")
	}
	bi.config.driver = _tx.drv
	return bi
}

// String implements the fmt.Stringer.
func (bi *BannedIP) String() string {
	var builder strings.Builder
	builder.WriteString("BannedIP(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bi.ID))
	builder.WriteString("create_time=")
	builder.WriteString(bi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(bi.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(bi.IP)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(bi.Country)
	builder.WriteByte(')')
	return builder.String()
}

// BannedIPs is a parsable slice of BannedIP.
type BannedIPs []*BannedIP