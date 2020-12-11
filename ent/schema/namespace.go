package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Namespace holds the schema definition for the Namespace entity.
type Namespace struct {
	ent.Schema
}

// Fields of the Namespace.
func (Namespace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			StructTag(`json:"name"`).
			StorageKey("name"),
		field.String("dockerRepo").
			StructTag(`json:"dockerRepo"`).
			StorageKey("dockerRepo"),
		field.String("repoNamespace").
			StructTag(`json:"repoNamespace"`).
			StorageKey("repoNamespace"),
		field.String("active").
			StructTag(`json:"active"`).
			StorageKey("active"),
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

// Edges of the Namespace.
func (Namespace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cluster",K8sCluster.Type).Ref("namespaces").Unique(),
		edge.To("applications",Application.Type),
	}
}
