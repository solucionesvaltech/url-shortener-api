package helper

import (
	"context"
)

type ContextKey string

const (
	domainContextKey ContextKey = "domain"
)

func Set(ctx context.Context, key ContextKey, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func Get(ctx context.Context, key ContextKey) string {
	if value, ok := ctx.Value(key).(string); ok {
		return value
	}
	return ""
}

func SetDomain(ctx context.Context, value string) context.Context {
	return Set(ctx, domainContextKey, value)
}

func GetDomain(ctx context.Context) string {
	value := Get(ctx, domainContextKey)
	if value != "" {
		return value
	}
	return ""
}
