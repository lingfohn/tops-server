package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("applicationId").
			StructTag(`json:"applicationId"`).
			StorageKey("applicationId"),
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

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("application",Application.Type).
			StructTag(`json:"application,omitempty"`).
			Ref("instances").
			Unique(),
		edge.To("builds",Build.Type).
			StructTag(`json:"builds,omitempty"`),
		edge.To("config",HelmConfig.Type).
			StructTag(`json:"config,omitempty"`).
			Unique(),
	}
}
