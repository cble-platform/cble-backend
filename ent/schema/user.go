package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
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
		field.String("username").
			Unique(),
		field.String("email"),
		field.String("password").
			Sensitive(),
		field.String("first_name"),
		field.String("last_name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),

		edge.From("deployments", Deployment.Type).
			Ref("requester"),
	}
}
