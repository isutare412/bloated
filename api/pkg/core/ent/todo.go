// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/isutare412/bloated/api/pkg/core/ent/todo"
	"github.com/isutare412/bloated/api/pkg/core/ent/user"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// Task holds the value of the "task" field.
	Task string `json:"task,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges        TodoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			values[i] = new(sql.NullInt64)
		case todo.FieldTask:
			values[i] = new(sql.NullString)
		case todo.FieldCreateTime, todo.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case todo.FieldOwnerID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todo.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case todo.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case todo.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				t.OwnerID = *value
			}
		case todo.FieldTask:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field task", values[i])
			} else if value.Valid {
				t.Task = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Todo.
// This includes values selected through modifiers, order, etc.
func (t *Todo) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Todo entity.
func (t *Todo) QueryOwner() *UserQuery {
	return NewTodoClient(t.config).QueryOwner(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", t.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("task=")
	builder.WriteString(t.Task)
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
