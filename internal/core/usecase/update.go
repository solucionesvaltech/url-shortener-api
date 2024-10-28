package usecase

import (
	"context"
	"fmt"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/helper"
)

// UpdateURL updates the original URL of an existing short URL
func (u *URLUseCase) UpdateURL(ctx context.Context, shortID string, newOriginalURL string) error {
	if !helper.IsValidURL(newOriginalURL) {
		return customerror.ValidationError(
			helper.GetDomain(ctx),
			fmt.Sprintf("invalid format for URL: %s", newOriginalURL),
			"",
		)
	}

	url, err := u.findURL(ctx, shortID, true)
	if err != nil {
		return err
	}

	url.Original = newOriginalURL
	if err := u.repo.Update(*url); err != nil {
		return customerror.SavingError(
			helper.GetDomain(ctx),
			fmt.Sprintf("an error occur trying to update for id: %s", shortID),
			err.Error(),
		)
	}

	if err := u.cache.Clean(ctx, url.Short); err != nil {
		return err
	}

	return nil
}
