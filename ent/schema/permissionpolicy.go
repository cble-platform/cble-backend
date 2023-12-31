package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionPolicy holds the schema definition for the PermissionPolicy entity.
type PermissionPolicy struct {
	ent.Schema
}

// Fields of the PermissionPolicy.
func (PermissionPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("type").
			Values("ALLOW", "DENY").
			Optional().
			Default("ALLOW"),
		field.Bool("is_inherited").
			Optional().
			Default(false),
	}
}

// Edges of the PermissionPolicy.
func (PermissionPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permission", Permission.Type).
			Unique().
			Required(),
		edge.To("group", Group.Type).
			Unique().
			Required(),
	}
}
