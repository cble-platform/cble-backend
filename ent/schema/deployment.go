package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Deployment holds the schema definition for the Deployment entity.
type Deployment struct {
	ent.Schema
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
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
		field.String("description").
			Optional().
			Default(""),
		field.JSON("template_vars", map[string]interface{}{}).
			Default(make(map[string]interface{})),
		field.JSON("deployment_vars", map[string]interface{}{}).
			Default(make(map[string]interface{})),
		field.JSON("deployment_state", map[string]string{}).
			Default(make(map[string]string)),
		field.Enum("state").
			Values("UNKNOWN", "INPROGRESS", "ACTIVE", "DESTROYED").
			Default("UNKNOWN"),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blueprint", Blueprint.Type).
			Unique().
			Required(),
		edge.To("requester", User.Type).
			Unique().
			Required(),
	}
}
