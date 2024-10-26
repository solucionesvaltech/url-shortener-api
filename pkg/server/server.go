package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
	"url-shortener/internal/adapter/driver/api"
	"url-shortener/pkg/config"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/log"
	"url-shortener/pkg/metric"
)

type Server struct {
	server       *echo.Echo
	conf         *config.Config
	routes       []*api.RequestHandler
	metricClient metric.Metric
	errorHandler *customerror.ErrorHandler
}

func NewServer(
	server *echo.Echo,
	conf *config.Config,
	routes []*api.RequestHandler,
	metricClient metric.Metric,
	errorHandler *customerror.ErrorHandler,
) (*Server, error) {
	s := &Server{
		server:       server,
		conf:         conf,
		routes:       routes,
		metricClient: metricClient,
		errorHandler: errorHandler,
	}
	s.server.HideBanner = true
	err := s.setupRoutes(routes)
	return s, err
}

func (s *Server) Start() {
	port := ":" + s.conf.ServerConfig.Port
	timeout := time.Duration(s.conf.ServerConfig.TimeoutMinutes) * time.Minute
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	log.Log.Fatal(s.server.StartServer(server))
}

func (s *Server) Shutdown() {
	var shutdownTime time.Duration = 10
	currentContext, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)
	defer cancel()
	if err := s.server.Shutdown(currentContext); err != nil {
		log.Log.Fatal("error shutting down web server")
	}
	log.Log.Info("shutdown web server")
}

func (s *Server) setupRoutes(routes []*api.RequestHandler) error {
	for _, route := range routes {
		wrapper := api.NewRequestWrapper(route, s.metricClient, s.errorHandler)
		switch route.Method {
		case http.MethodGet:
			s.server.GET(route.Path, route.Handler.HandleRequest, wrapper.ApplyCommon)
		case http.MethodPost:
			s.server.POST(route.Path, route.Handler.HandleRequest, wrapper.ApplyCommon)
		case http.MethodPut:
			s.server.PUT(route.Path, route.Handler.HandleRequest, wrapper.ApplyCommon)
		case http.MethodPatch:
			s.server.PATCH(route.Path, route.Handler.HandleRequest, wrapper.ApplyCommon)
		default:
			return fmt.Errorf("invalid method: %s for path: %s", route.Method, route.Path)
		}
	}
	s.server.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	return nil
}
