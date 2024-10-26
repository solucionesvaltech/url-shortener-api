package api

import (
	"url-shortener/internal/adapter/driver/api/health"
	"url-shortener/internal/adapter/driver/api/url/create"
	"url-shortener/internal/adapter/driver/api/url/details"
	"url-shortener/internal/adapter/driver/api/url/get"
	"url-shortener/internal/adapter/driver/api/url/toggle"
	"url-shortener/internal/adapter/driver/api/url/update"
	"url-shortener/pkg/config"
)

func NewHandlers(
	conf *config.Config,
	health *health.Handler,
	create *create.Handler,
	get *get.Handler,
	update *update.Handler,
	toggle *toggle.Handler,
	details *details.Handler,
) []*RequestHandler {
	handlers := make([]*RequestHandler, 0)

	// Lista de definiciones de handlers
	handlerDefinitions := []struct {
		Method, Path string
		Handler      Handler
	}{
		{conf.ServerConfig.Routes.Health.Method, conf.ServerConfig.Routes.Health.Path, health},
		{conf.ServerConfig.Routes.Create.Method, conf.ServerConfig.Routes.Create.Path, create},
		{conf.ServerConfig.Routes.Get.Method, conf.ServerConfig.Routes.Get.Path, get},
		{conf.ServerConfig.Routes.Update.Method, conf.ServerConfig.Routes.Update.Path, update},
		{conf.ServerConfig.Routes.Toggle.Method, conf.ServerConfig.Routes.Toggle.Path, toggle},
		{conf.ServerConfig.Routes.Details.Method, conf.ServerConfig.Routes.Details.Path, details},
	}

	for _, def := range handlerDefinitions {
		handlers = append(handlers, &RequestHandler{
			Method:  def.Method,
			Path:    def.Path,
			Handler: def.Handler,
		})
	}

	return handlers
}
