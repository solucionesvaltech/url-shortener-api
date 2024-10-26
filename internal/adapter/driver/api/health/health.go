package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"url-shortener/internal/core/usecase"

	"url-shortener/pkg/config"
)

type Handler struct {
	startTime time.Time
	appName   string
}

type health struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	AppName string `json:"appName"`
}

func NewHandler(conf *config.Config) *Handler {
	return &Handler{
		startTime: time.Now(),
		appName:   conf.AppName,
	}
}

func (h *Handler) HandleRequest(context echo.Context) error {
	healthCheck := health{
		Status:  "UP",
		Uptime:  time.Since(h.startTime).String(),
		AppName: h.appName,
	}
	return context.JSON(http.StatusOK, healthCheck)
}

func (h *Handler) Domain() usecase.UseCase {
	return usecase.General
}
