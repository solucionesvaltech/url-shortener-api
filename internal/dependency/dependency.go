package dependency

import (
	"context"
	"errors"
	"fmt"
	"url-shortener/internal/adapter/driven/dynamo"
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
	dynamoClient dynamo.ClientInterface
	cache        port.URLCache
}

func NewDependency(
	cfg *config.Config,
	server *server.Server,
	metricClient metric.Metric,
	cache port.URLCache,
) *Dependency {
	return &Dependency{
		Conf:         cfg,
		server:       server,
		metricClient: metricClient,
		cache:        cache,
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

	err = d.metricClient.Start()
	if err != nil {
		return
	}

	err = d.cache.Ping(ctx)
	if err != nil {
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
