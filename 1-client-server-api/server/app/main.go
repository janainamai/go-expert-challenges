package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/core"
	"server/internal/handler"
	"server/internal/infra"
	"server/internal/infra/database"
	"server/internal/infra/rest"
	"syscall"
	"time"
)

type Resources struct {
	QuotationHandler handler.QuotationHandler
}

func main() {
	r := setupResources()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", r.QuotationHandler.GetCurrentQuotation)
	mux.HandleFunc("/cotacao/audit", r.QuotationHandler.GetRegistereQuotations)

	wrappedMux := loggingMiddleware(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("Listening on port 8080\n")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not start server on port: 8080, error: %s", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Could not shutdown server, %s", err)
	}

	fmt.Println("Server gracefully stopped")
}

func setupResources() *Resources {
	quotationRest := rest.NewQuotationRest()
	quotationRepository := database.NewQuotationRepository()

	orchestrator := infra.NewQuotationOrchestrator(quotationRepository, quotationRest)

	quotationCore := core.NewQuotationCore(orchestrator)

	quotationHandler := handler.NewQuotationHandler(quotationCore)

	return &Resources{
		QuotationHandler: quotationHandler,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
