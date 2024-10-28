package usecase

import (
	"context"
	"fmt"
	"url-shortener/internal/core/domain"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/helper"
)

func (u *URLUseCase) DetailURL(ctx context.Context, shortID string) (*domain.URL, error) {
	url, err := u.repo.Find(shortID)
	if err != nil {
		return nil, customerror.DatabaseError(
			helper.GetDomain(ctx),
			fmt.Sprintf("original url not found for id: %s", shortID),
			err.Error(),
		)
	}

	return url, nil
}
