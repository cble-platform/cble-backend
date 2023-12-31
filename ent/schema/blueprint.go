package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Blueprint holds the schema definition for the Blueprint entity.
type Blueprint struct {
	ent.Schema
}

// Fields of the Blueprint.
func (Blueprint) Fields() []ent.Field {
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
		field.String("name"),
		field.String("description"),
		field.Bytes("blueprint_template"),
	}
}

// Edges of the Blueprint.
func (Blueprint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent_group", Group.Type).
			Unique().
			Required(),
		edge.To("provider", Provider.Type).
			Unique().
			Required(),

		edge.From("deployments", Deployment.Type).
			Ref("blueprint"),
	}
}
