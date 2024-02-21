package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/engine/models"
	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/google/uuid"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
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
			Values("resource", "data").
			Default("resource"),
		field.String("key").
			Comment("The resource/data key from the blueprint"),
		field.String("resource_type").
			Comment("The resource/data string from the blueprint"),
		field.JSON("features", pgrpc.Features{}).Optional().Default(pgrpc.Features{}),
		field.JSON("object", &models.Object{}).
			Comment("The entire resource/data object from the blueprint"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blueprint", Blueprint.Type).
			Unique().
			Required().
			Comment("Blueprint containing this resource"),
		edge.To("required_by", Resource.Type).
			Comment("Stores all dependents of this resource").
			From("depends_on").
			Comment("Stores all dependencies of this resource"),
	}
}
