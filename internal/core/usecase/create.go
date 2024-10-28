package usecase

import (
	"context"
	"fmt"
	"url-shortener/internal/core/domain"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/helper"
)

// CreateShortURL generate a short ID and persist using the repository
func (u *URLUseCase) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	shortID, err := helper.GenerateShortID()
	if err != nil {
		return "", customerror.SetupError(
			helper.GetDomain(ctx),
			fmt.Sprintf("an error occur trying to create id for URL: %s", originalURL),
			err.Error(),
		)
	}

	if !helper.IsValidURL(originalURL) {
		return "", customerror.ValidationError(
			helper.GetDomain(ctx),
			fmt.Sprintf("invalid format for URL: %s", originalURL),
			"",
		)
	}

	url := domain.URL{Short: shortID, Original: originalURL, Enabled: true}
	if err := u.repo.Save(url); err != nil {
		return "", err
	}

	if err = u.cache.Set(ctx, shortID, url.Original); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", u.domain, shortID), nil
}
