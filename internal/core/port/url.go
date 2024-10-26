package port

import (
	"context"
	"url-shortener/internal/core/domain"
)

//go:generate mockgen -source=url.go -destination=../../../mocks/url.go -package=mock

// URLCache defines cache methods for URLs
type URLCache interface {
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{}) error
	Clean(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (string, error)
	Shutdown() error
}

// URLRepository is an interface for interacting with Original related data.
type URLRepository interface {
	Save(URL domain.URL) error
	Find(shortID string) (*domain.URL, error)
	Update(URL domain.URL) error
}

// URLShortenerUseCase is an interface for interacting with Original related business logic.
type URLShortenerUseCase interface {
	CreateShortURL(ctx context.Context, originalURL string) (string, error)
	ResolveURL(ctx context.Context, shortID string) (string, error)
	UpdateURL(ctx context.Context, shortID string, newOriginalURL string) error
	ToggleURLStatus(ctx context.Context, shortID string, enable bool) error
	DetailURL(ctx context.Context, shortID string) (*domain.URL, error)
}
