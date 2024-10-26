package create

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/internal/adapter/driver/api/common"
	"url-shortener/internal/core/port"
	"url-shortener/internal/core/usecase"
	"url-shortener/pkg/helper"
)

type request struct {
	URL string `json:"url" validate:"required,url"`
}

type Handler struct {
	useCase port.URLShortenerUseCase
}

func NewHandler(useCase port.URLShortenerUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) HandleRequest(serverCtx echo.Context) error {
	var payload request
	err := common.Deserialize(serverCtx, &payload)
	if err != nil {
		return common.DeserializingError(serverCtx, err, h.Domain().String())
	}

	ctx := context.Background()
	ctx = helper.SetDomain(ctx, h.Domain().String())
	newURL, err := h.useCase.CreateShortURL(ctx, payload.URL)
	if err != nil {
		return common.InternalError(serverCtx, err, h.Domain().String())
	}
	return serverCtx.JSON(http.StatusCreated, newURL)
}

func (h *Handler) Domain() usecase.UseCase {
	return usecase.ShortURL
}
