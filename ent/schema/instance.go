package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("application",Application.Type).
			Ref("instances").
			Unique(),
		edge.To("builds",Build.Type),
		edge.To("config",HelmConfig.Type).Unique(),
	}
}
