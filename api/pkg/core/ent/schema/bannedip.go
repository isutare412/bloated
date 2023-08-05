package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BannedIP holds the schema definition for the BannedIP entity.
type BannedIP struct {
	ent.Schema
}

// Fields of the BannedIP.
func (BannedIP) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").
			NotEmpty().
			MaxLen(50).
			Unique(),
		field.String("country").
			Optional().
			NotEmpty().
			MaxLen(60),
	}
}

// Mixin of the BannedIP.
func (BannedIP) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the BannedIP.
func (BannedIP) Edges() []ent.Edge {
	return nil
}
