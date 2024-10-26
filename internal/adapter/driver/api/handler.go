package api

import (
	"github.com/labstack/echo/v4"
	"url-shortener/internal/core/usecase"
)

type Handler interface {
	HandleRequest(ctx echo.Context) error
	Domain() usecase.UseCase
}

type RequestHandler struct {
	Method  string
	Path    string
	Handler Handler
}
