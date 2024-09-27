package setup

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/janainamai/go-expert-challenges/cmd/resources"
)

func InitServer(resources *resources.Resources) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/order", resources.SaveOrderRestHandler.Create)
	r.Get("/order", resources.ListOrdersRestHandler.List)

	fmt.Printf("REST Server - Listening on port: %d\n", 3000)
	http.ListenAndServe(":3000", r)
}
