package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			StructTag(`json:"username"`),
		field.String("password").
			Optional().
			Sensitive(),
		field.Bool("initial").
			Default(true).
			StructTag(`json:"initial"`),
		field.String("oauser").
			Optional(),
		field.String("name").
			Optional(),
		field.String("telephone").
			Optional(),
		field.String("email").
			StructTag(`json:"email"`),
		field.String("avatar").
			Optional(),
		field.Bool("super").
			Default(false).
			StructTag(`json:"super"`),
		field.Int("status").
			Optional(),
		field.Time("lastLogin").
			StorageKey("lastLogin").
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
