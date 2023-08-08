package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TokenHistory holds the schema definition for the TokenHistory entity.
type TokenHistory struct {
	ent.Schema
}

// Fields of the TokenHistory.
func (TokenHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			NotEmpty(),
		field.String("user_name").
			NotEmpty(),
	}
}

// Mixin of the TokenHistory.
func (TokenHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the TokenHistory.
func (TokenHistory) Edges() []ent.Edge {
	return nil
}
