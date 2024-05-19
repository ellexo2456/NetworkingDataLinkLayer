package main

import (
	"context"
	"github.com/ellexo2456/NetworkingDataLinkLayer/internal/lib/logger/handlers/logruspretty"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/ellexo2456/NetworkingDataLinkLayer/internal/app"
)

// @title DataLinkLayer API
// @version 1.0
// @description API server for DataLinkLayer application

// @host http://localhost:8081
// @BasePath /

func main() {
	log := setupLogger()

	log.Info("Application start")
	ctx := context.Background()

	application, err := app.New(ctx, log)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.Run()
	log.Info("Application terminated")
}

func setupLogger() *logrus.Logger {
	var log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	return setupPrettyLogrus(log)
}

func setupPrettyLogrus(log *logrus.Logger) *logrus.Logger {
	prettyHandler := logruspretty.NewPrettyHandler(os.Stdout)
	log.SetFormatter(prettyHandler)
	return log
}
