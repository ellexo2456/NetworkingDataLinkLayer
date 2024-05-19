package app

import (
	"context"
	"github.com/ellexo2456/NetworkingDataLinkLayer/internal/config"
	"github.com/ellexo2456/NetworkingDataLinkLayer/internal/http"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Config  *config.Config
	Handler *http.Handler
	log     *logrus.Logger
}

func New(ctx context.Context, log *logrus.Logger) (*Application, error) {
	cfg := config.MustLoad()
	log.WithField("config", cfg).Info("config parsed")

	h := http.NewHandler(cfg.BaseURL, log)
	app := &Application{
		Config:  cfg,
		Handler: h,
		log:     log,
	}

	return app, nil
}
