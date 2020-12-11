package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("method"),
		field.String("fullpath"),
		field.String("action"),
		field.String("summary"),
		field.Int("controlLevel").
			StructTag(`json:"controlLevel"`).
			StorageKey("controlLevel"),
		field.Int("status").
			Default(1).StructTag(`json:"status"`),
		field.Time("createdAt").
			Default(time.Now).
			StorageKey("createdAt").
			Immutable(),
		field.Time("updatedAt").
			StorageKey("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
