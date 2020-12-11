package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"time"
)

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("multi").
			Default(true).
			StructTag(`json:"multi"`),
		field.Int("projectId").
			StructTag(`json:"projectId"`).
			StorageKey("projectId"),
		field.Int("namespaceId").
			StructTag(`json:"namespaceId"`).
			StorageKey("namespaceId"),
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

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("namespace",Namespace.Type).
			StructTag(`json:"namespace,omitempty"`).
			Ref("applications").
			Unique(),
		edge.From("project",Project.Type).
			StructTag(`json:"project,omitempty"`).
			Ref("applications").
			Unique(),
		edge.To("instances",Instance.Type).
			StructTag(`json:"instances,omitempty"`),
		edge.To("config",HelmConfig.Type).
			StructTag(`json:"config,omitempty"`).Unique(),
	}
}

func (Application) Indexes() []ent.Index  {
	return []ent.Index{
		index.Edges("namespace","project").Unique(),
	}
}