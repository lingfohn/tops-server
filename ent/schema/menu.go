package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.String("name"),
		field.String("component"),
		field.Int("parentId").
			StructTag(`json:"parentId"`).
			StorageKey("parentId"),
		field.String("redirect").
			Optional(),
		field.Int("weight").
			Optional().
			Default(10),
		field.Int("level"),
		field.String("title").
			Optional(),
		field.String("icon").
			Optional(),
		field.String("target").
			Optional(),
		field.Bool("keepAlive").
			StructTag(`json:"keepAlive"`).
			StorageKey("keepAlive"),
		field.Bool("show").
			Optional().
			Nillable(),
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

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submenus", Menu.Type).
			StructTag(`json:"submenus,omitempty"`).
			From("parent").
			StructTag(`json:"parent,omitempty"`).
			Unique(),
	}
}
