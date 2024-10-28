package usecase

import (
	"context"
	"url-shortener/pkg/log"
)

// ResolveURL search for the original URL using the short ID
func (u *URLUseCase) ResolveURL(ctx context.Context, shortID string) (string, error) {
	cachedURL, err := u.cache.Get(ctx, shortID)
	if err == nil && cachedURL != "" {
		return cachedURL, nil
	}

	log.Log.Infof("shortID: %s not found in cache, going to search in DB", shortID)
	url, err := u.findURL(ctx, shortID, true)
	if err != nil {
		return "", err
	}

	if err = u.cache.Set(ctx, shortID, url.Original); err != nil {
		log.Log.Errorf("cache failed saving shortID: %s", shortID)
	}

	return url.Original, nil
}
