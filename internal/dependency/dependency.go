package dependency

import (
	"context"
	"errors"
	"fmt"
	"url-shortener/internal/core/port"

	"url-shortener/pkg/config"
	"url-shortener/pkg/log"
	"url-shortener/pkg/metric"
	"url-shortener/pkg/server"
)

type Dependency struct {
	Conf         *config.Config
	server       *server.Server
	metricClient metric.Metric
	cache        port.URLCache
	repo         port.URLRepository
}

func NewDependency(
	cfg *config.Config,
	server *server.Server,
	metricClient metric.Metric,
	cache port.URLCache,
	repo port.URLRepository,
) *Dependency {
	return &Dependency{
		Conf:         cfg,
		server:       server,
		metricClient: metricClient,
		cache:        cache,
		repo:         repo,
	}
}

func (d *Dependency) Start(ctx context.Context) (err error) {
	log.Log.Info("starting dependency")
	defer func() {
		if r := recover(); r != nil {
			errorMessage := fmt.Sprintf("%v", r)
			err = errors.New(errorMessage)
		}
	}()

	if err = d.metricClient.Start(); err != nil {
		return
	}

	if err = d.cache.Ping(ctx); err != nil {
		return
	}

	if err = d.repo.Init(); err != nil {
		return
	}

	go d.server.Start()
	return
}

func (d *Dependency) Shutdown() {
	log.Log.Info("starting dependency shutdown...")
	d.server.Shutdown()
	if d.cache.Shutdown() != nil {
		log.Log.Error("failed to shutdown cache")
	}
	log.Log.Info("dependency shutdown complete")
}
