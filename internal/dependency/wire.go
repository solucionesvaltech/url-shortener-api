//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"url-shortener/internal/adapter/driven/dynamo"
	"url-shortener/internal/adapter/driven/dynamo/url"
	"url-shortener/internal/adapter/driven/redis"
	"url-shortener/internal/adapter/driver/api"
	"url-shortener/internal/adapter/driver/api/health"
	"url-shortener/internal/adapter/driver/api/url/create"
	"url-shortener/internal/adapter/driver/api/url/details"
	"url-shortener/internal/adapter/driver/api/url/get"
	"url-shortener/internal/adapter/driver/api/url/toggle"
	"url-shortener/internal/adapter/driver/api/url/update"
	"url-shortener/internal/core/port"
	"url-shortener/internal/core/usecase"
	"url-shortener/pkg/config/configyaml"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/metric"
	"url-shortener/pkg/metric/prometheus"
	"url-shortener/pkg/server"
)

var ProviderSet = wire.NewSet(
	configyaml.NewConfigYaml,
	customerror.NewErrorHandler,
	dynamo.NewDynamoDBClient,
	wire.Bind(new(dynamo.ClientInterface), new(*dynamo.Client)),
	redis.NewRedisClient,
	wire.Bind(new(port.URLCache), new(*redis.Client)),
	url.NewURLRepository,
	wire.Bind(new(port.URLRepository), new(*url.Repository)),
	usecase.NewURLUseCase,
	wire.Bind(new(port.URLShortenerUseCase), new(*usecase.URLUseCase)),
	health.NewHandler,
	create.NewHandler,
	get.NewHandler,
	update.NewHandler,
	toggle.NewHandler,
	details.NewHandler,
	api.NewHandlers,
	echo.New,
	server.NewServer,
	prometheus.NewClient,
	wire.Bind(new(metric.Metric), new(*prometheus.Client)),
	NewDependency,
)

func Init() (*Dependency, error) {
	wire.Build(ProviderSet)
	return &Dependency{}, nil
}
