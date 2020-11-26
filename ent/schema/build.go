package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Build holds the schema definition for the Build entity.
type Build struct {
	ent.Schema
}

// Fields of the Build.
func (Build) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Build.
func (Build) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instance",Instance.Type).
			Ref("builds").
			Unique(),
	}
}
