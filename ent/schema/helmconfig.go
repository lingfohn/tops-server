package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// HelmConfig holds the schema definition for the HelmConfig entity.
type HelmConfig struct {
	ent.Schema
}

// Fields of the HelmConfig.
func (HelmConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("chartVersion").
			StructTag(`json:"chartVersion"`).
			StorageKey("chartVersion"),
		field.String("active").
			StructTag(`json:"active"`).
			StorageKey("active"),
		field.Bool("enableService").
			StructTag(`json:"enableService"`).
			StorageKey("enableServcie").
			Default(false),
		field.String("serviceType").
			StructTag(`json:"serviceType"`).
			StorageKey("serviceType").
			Default("ClusterIP"),
		field.Int("nodePort").
			StructTag(`json:"nodePort"`).
			StorageKey("nodePort"),
		field.String("limitMem").
			StructTag(`json:"limitMem"`).
			StorageKey("limitMem"),
		field.String("limitCPU").
			StructTag(`json:"limitCPU"`).
			StorageKey("limitCPU"),
		field.String("reqCPU").
			StructTag(`json:"reqCPU"`).
			StorageKey("reqCPU"),
		field.String("reqMem").
			StructTag(`json:"reqMem"`).
			StorageKey("reqMem"),
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

// Edges of the HelmConfig.
func (HelmConfig) Edges() []ent.Edge {
	return nil
}
