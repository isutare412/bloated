package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/isutare412/bloated/api/pkg/core/enum"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("email").
			NotEmpty().
			MaxLen(256),
		field.String("user_name").
			NotEmpty().
			MaxLen(800),
		field.String("given_name").
			NotEmpty().
			MaxLen(800),
		field.String("family_name").
			NotEmpty().
			MaxLen(800),
		field.String("photo_url").
			NotEmpty(),
		field.Enum("origin").
			GoType(enum.Issuer("")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
