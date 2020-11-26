package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Namespace holds the schema definition for the Namespace entity.
type Namespace struct {
	ent.Schema
}

// Fields of the Namespace.
func (Namespace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
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
