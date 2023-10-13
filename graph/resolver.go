package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/cble-platform/backend/ent"
	"github.com/cble-platform/backend/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	ent *ent.Client
	// redis *redis.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: &Resolver{
			ent: client,
			// rdb:           rdb,
		},
	}
	// c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*model.UserRole) (res interface{}, err error) {
	// 	currentUser, err := auth.ForContext(ctx)
	// 	if err != nil {
	// 		return nil, auth.AUTH_REQUIRED_GQL_ERROR
	// 	}
	// 	for _, role := range roles {
	// 		if role.String() == string(currentUser.Role) {
	// 			return next(ctx)
	// 		}
	// 	}
	// 	return nil, auth.PERMISSION_DENIED_GQL_ERROR
	// }
	return generated.NewExecutableSchema(c)
}