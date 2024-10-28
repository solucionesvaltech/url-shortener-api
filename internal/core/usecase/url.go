package usecase

import (
	"context"
	"fmt"
	"url-shortener/internal/core/domain"
	"url-shortener/internal/core/port"
	"url-shortener/pkg/config"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/helper"
)

// URLUseCase implements the business logic of the URLShortenerService interface
type URLUseCase struct {
	repo   port.URLRepository
	cache  port.URLCache
	domain string
}

// NewURLUseCase create a new instance of URLUseCase
func NewURLUseCase(repo port.URLRepository, cache port.URLCache, conf *config.Config) *URLUseCase {
	return &URLUseCase{repo: repo, cache: cache, domain: conf.Domain}
}

func (u *URLUseCase) findURL(ctx context.Context, shortID string, strict bool) (*domain.URL, error) {
	url, err := u.repo.Find(shortID)
	if err != nil {
		return nil, customerror.DatabaseError(
			helper.GetDomain(ctx),
			fmt.Sprintf("original url not found for id: %s", shortID),
			err.Error(),
		)
	}

	if url == nil || (strict && !url.Enabled) {
		return nil, customerror.BusinessError(
			helper.GetDomain(ctx),
			fmt.Sprintf("original url not found or disabled for id: %s", shortID),
			"",
		)
	}
	return url, nil
}
