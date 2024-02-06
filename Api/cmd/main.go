package main

import (
	"Api/internal/app"
	"context"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	ctx := context.Background()
	application, err := app.NewApp(&ctx)
	if err != nil {
		log.Fatalf("Failed an initialization app: %s", err.Error())
	}

	err = application.Run()
	if err != nil {
		log.Fatalf("Failed to run app: %s", err.Error())
	}

	//r := chi.NewRouter()
	//

	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	//
	//http.ListenAndServe(":3000", r)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
