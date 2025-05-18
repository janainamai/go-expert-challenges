package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/cmd/configs"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/cmd/resources"
	"github.com/sirupsen/logrus"
)

func main() {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		logrus.Fatal("WEATHER_API_KEY environment variable is not set")
	}

	r := resources.LoadResources(key)

	configs.SetupLogging()

	mux := http.NewServeMux()
	mux.HandleFunc("/temperature", r.GetTemperatureHandler.Get)

	wrappedMux := loggingMiddleware(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		logrus.Info("Listening on port 8080")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("Could not listen on port 8080, %s", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	err := server.Shutdown(ctx)
	if err != nil {
		logrus.Fatalf("Could not shutdown server, %s", err)
	}

	fmt.Println("Server gracefully stopped")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info(fmt.Sprintf("Request received: %s %s", r.Method, r.URL.Path))

		next.ServeHTTP(w, r)
	})
}
