package setup

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/resources"
)

func InitServer(cfg *configs.Config, resources *resources.Resources) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/order", resources.CreateOrderRestHandler.Create)
	r.Get("/order", resources.ListOrdersRestHandler.List)

	fmt.Printf("REST Server - Listening on port: %s\n", cfg.RestServer.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.RestServer.Port), r)
	if err != nil {
		panic(fmt.Sprintf("error initing rest server: %s", err.Error()))
	}
}
