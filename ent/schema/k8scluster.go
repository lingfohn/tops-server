package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// K8sCluster holds the schema definition for the K8sCluster entity.
type K8sCluster struct {
	ent.Schema
}

// Fields of the K8sCluster.
func (K8sCluster) Fields() []ent.Field {
	return []ent.Field{
		field.String("cluster"),
		field.String("helmApi").
			StructTag(`json:"helmApi"`).
			StorageKey("helmApi"),
		field.String("accessToken").
			StructTag(`json:"accessToken"`).
			StorageKey("accessToken").
			Optional(),
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

// Edges of the K8sCluster.
func (K8sCluster) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("namespaces",Namespace.Type).
			StructTag(`json:"namespaces,omitempty"`),
	}
}
