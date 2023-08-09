package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			NotEmpty().
			MaxLen(36).
			Optional(),
		field.UUID("owner_id", uuid.UUID{}),
		field.String("task").
			NotEmpty().
			MaxLen(3000),
	}
}

// Mixin of the Todo.
func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("todos").
			Unique().
			Required().
			Field("owner_id"),
	}
}

// Indexes of the Todo.
func (Todo) Indexes() []ent.Index {
	return nil
}
