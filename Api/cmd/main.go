package main

import (
	"Api/internal/domains/models"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {

	answer := &models.Answer{Id: 1}
	data, err := json.Marshal(answer)
	if err != nil {
		log.Fatalf("Error json serialization: %s", err)
	}

	fmt.Println(string(data))

	//ctx := context.Background()
	//application, err := app.NewApp(&ctx)
	//if err != nil {
	//	log.Fatalf("Failed an initialization app: %s", err.Error())
	//}
	//
	//err = application.Run()
	//if err != nil {
	//	log.Fatalf("Failed to run app: %s", err.Error())
	//}
	//
	//r := chi.NewRouter()
	//
	//r.Use(middleware.Logger)
	//r.Use(cors.Handler(cors.Options{
	//	AllowedOrigins:   []string{"https://*", "http://*"},
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//	ExposedHeaders:   []string{"Link"},
	//	AllowCredentials: false,
	//	MaxAge:           300,
	//}))
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
