package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("projectName").
			Unique().
			StructTag(`json:"projectName"`).
			StorageKey("projectName"),
		field.String("proType").
			StorageKey("proType"),
		field.String("description").
			Optional(),
		field.String("gitlab").
			Unique().
			StructTag("gitlab"),
		field.Int("port").
			StructTag("port"),
		field.Int("debugPort").
			StructTag(`json:"debugPort"`).
			StorageKey("debugPort"),
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

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("applications",Application.Type),
	}
}
