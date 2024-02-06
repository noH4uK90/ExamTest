package main

import (
	"Api/internal/app"
	"context"
	"log"
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
}
