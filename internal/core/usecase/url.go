package usecase

import (
	"context"
	"fmt"
	"url-shortener/internal/core/domain"
	"url-shortener/internal/core/port"
	"url-shortener/pkg/customerror"
	"url-shortener/pkg/helper"
)

// URLUseCase implements the business logic of the URLShortenerService interface
type URLUseCase struct {
	repo  port.URLRepository
	cache port.URLCache
}

// NewURLUseCase create a new instance of URLUseCase
func NewURLUseCase(repo port.URLRepository, cache port.URLCache) *URLUseCase {
	return &URLUseCase{repo: repo, cache: cache}
}

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

	url := domain.URL{Short: shortID, Original: originalURL, Enabled: true}
	if err := u.repo.Save(url); err != nil {
		return "", err
	}

	if err = u.cache.Set(ctx, shortID, url.Original); err != nil {
		return "", err
	}

	return shortID, nil
}

// ResolveURL search for the original URL using the short ID
func (u *URLUseCase) ResolveURL(ctx context.Context, shortID string) (string, error) {
	cachedURL, err := u.cache.Get(ctx, shortID)
	if err == nil && cachedURL != "" {
		return cachedURL, nil
	}

	url, err := u.findURL(ctx, shortID, true)
	if err != nil {
		return "", err
	}

	if err = u.cache.Set(ctx, shortID, url.Original); err != nil {
		return "", err
	}

	return url.Original, nil
}

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
