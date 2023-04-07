package utils

import (
	"context"
	"time"
)

func newSecContext(nb int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*time.Duration(nb))
}

func wc(ctx context.Context, cancel context.CancelFunc) context.Context {
	return ctx
}
