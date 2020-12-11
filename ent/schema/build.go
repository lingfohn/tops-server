package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Build holds the schema definition for the Build entity.
type Build struct {
	ent.Schema
}

// Fields of the Build.
func (Build) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag").
			StructTag(`json:"tag"`).
			StorageKey("tag"),
		field.String("name").
			StructTag(`json:"name"`).
			StorageKey("name"),
		field.Int("status").
			StructTag(`json:"status"`).
			StorageKey("status").
			Default(0),
		field.String("commitId").
			StructTag(`json:"commitId"`).
			StorageKey("commitId"),
		field.String("shortId").
			StructTag(`json:"shortId"`).
			StorageKey("shortId"),
		field.String("message").
			StructTag(`json:"message"`).
			StorageKey("message"),
		field.String("committerName").
			StructTag(`json:"committerName"`).
			StorageKey("committerName"),
		field.Time("committedData").
			StructTag(`json:"committedData"`).
			StorageKey("committedData"),
		field.Time("buildTime").
			StructTag(`json:"buildTime"`).
			StorageKey("buildTime"),
		field.Int("jenkinsQueueId").
			StructTag(`json:"jenkinsQueueId"'`).
				StorageKey("jenkinsQueueId"),
		field.Int("jenkinsBuildId").
			Optional().
			StructTag(`json:"jenkinsBuildId"`).
			StorageKey("jenkinsBuildId"),
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

// Edges of the Build.
func (Build) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instance",Instance.Type).
			Ref("builds").
			Unique(),
	}
}
