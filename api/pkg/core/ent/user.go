// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/isutare412/bloated/api/pkg/core/ent/user"
	"github.com/isutare412/bloated/api/pkg/core/enum"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// GivenName holds the value of the "given_name" field.
	GivenName string `json:"given_name,omitempty"`
	// FamilyName holds the value of the "family_name" field.
	FamilyName string `json:"family_name,omitempty"`
	// PhotoURL holds the value of the "photo_url" field.
	PhotoURL string `json:"photo_url,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin enum.Issuer `json:"origin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Todos holds the value of the todos edge.
	Todos []*Todo `json:"todos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodosOrErr returns the Todos value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TodosOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todos, nil
	}
	return nil, &NotLoadedError{edge: "todos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmail, user.FieldUserName, user.FieldGivenName, user.FieldFamilyName, user.FieldPhotoURL, user.FieldOrigin:
			values[i] = new(sql.NullString)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				u.UserName = value.String
			}
		case user.FieldGivenName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field given_name", values[i])
			} else if value.Valid {
				u.GivenName = value.String
			}
		case user.FieldFamilyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family_name", values[i])
			} else if value.Valid {
				u.FamilyName = value.String
			}
		case user.FieldPhotoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url", values[i])
			} else if value.Valid {
				u.PhotoURL = value.String
			}
		case user.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				u.Origin = enum.Issuer(value.String)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryTodos queries the "todos" edge of the User entity.
func (u *User) QueryTodos() *TodoQuery {
	return NewUserClient(u.config).QueryTodos(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(u.UserName)
	builder.WriteString(", ")
	builder.WriteString("given_name=")
	builder.WriteString(u.GivenName)
	builder.WriteString(", ")
	builder.WriteString("family_name=")
	builder.WriteString(u.FamilyName)
	builder.WriteString(", ")
	builder.WriteString("photo_url=")
	builder.WriteString(u.PhotoURL)
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(fmt.Sprintf("%v", u.Origin))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User