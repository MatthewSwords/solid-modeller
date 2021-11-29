//go:generate go run github.com/99designs/gqlgen
package solid

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/matthewswords/solid-modeller/common/config"
	"github.com/matthewswords/solid-modeller/modeller/api/solid/generated"
	"github.com/matthewswords/solid-modeller/modeller/api/solid/model"
)

// Resolver struct that is included with gql gen
type Resolver struct {
	config *config.Config
}

// Handler takes in config and tags on to the resolver
func Handler(c *config.Config) *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &Resolver{
				config: c,
			},
		}),
	)
}

// type queryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *queryResolver) Solid(ctx context.Context, id *int) (*model.Solid, error) {
	panic(fmt.Errorf("not implemented"))
}

// // Mutation returns generated.MutationResolver implementation.
// func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
