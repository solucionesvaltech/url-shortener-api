package usecase

import (
	"context"
	"fmt"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/helper"
)

// ToggleURLStatus change the enabled/disabled status of a short URL
func (u *URLUseCase) ToggleURLStatus(ctx context.Context, shortID string, enable bool) error {
	url, err := u.findURL(ctx, shortID, false)
	if err != nil {
		return err
	}

	url.Enabled = enable
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
