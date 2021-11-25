package api

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/matthewswords/solid-modeller/common/config"

	// "github.com/tablenu/servitor/common/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

type Option func(*chi.Mux)

func New(c *config.Config) (http.Handler, error) {
	router := defaultRouter("Pad API", withCORSOptions(c.CORS))

	if c.API.ShowGQLPlayground {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	return router, nil
}

// TODO: MOVE to a more common file
func withCORSOptions(corsOpts cors.Options) func(mux *chi.Mux) {
	return func(r *chi.Mux) {
		r.Use(cors.New(corsOpts).Handler)
	}
}

func defaultRouter(rootMsg string, opts ...Option) *chi.Mux {
	router := chi.NewRouter()

	for _, opt := range opts {
		opt(router)
	}

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":"not found"}`)
	})

	if rootMsg != "" {
		router.Get("/", func(w http.ResponseWriter, route *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Solid Modeller :: %s", rootMsg)
		})
	}

	return router
}
