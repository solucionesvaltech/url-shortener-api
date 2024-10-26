package details

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/internal/adapter/driver/api/common"
	"url-shortener/internal/core/port"
	"url-shortener/internal/core/usecase"
	"url-shortener/pkg/helper"
)

type Handler struct {
	useCase port.URLShortenerUseCase
}

func NewHandler(useCase port.URLShortenerUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) HandleRequest(serverCtx echo.Context) error {
	shortID := serverCtx.Param("shortID")
	if shortID == "" {
		return common.DeserializingError(serverCtx, errors.New("shortID is empty"), h.Domain().String())
	}
	ctx := context.Background()
	ctx = helper.SetDomain(ctx, h.Domain().String())
	url, err := h.useCase.DetailURL(ctx, shortID)
	if err != nil {
		return common.NotFoundError(serverCtx, err, h.Domain().String())
	}
	return serverCtx.JSON(http.StatusOK, url)
}

func (h *Handler) Domain() usecase.UseCase {
	return usecase.DetailsURL
}
