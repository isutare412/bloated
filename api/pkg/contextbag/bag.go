package contextbag

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/constant"
)

type ctxKeyBag struct{}

func WithBag(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyBag{}, newBag())
}

func Bag(ctx context.Context) *bag {
	b, ok := ctx.Value(ctxKeyBag{}).(*bag)
	if !ok {
		return nil
	}
	return b
}

type bag struct {
	dynamics map[constant.BagKey]any
}

func newBag() *bag {
	return &bag{
		dynamics: make(map[constant.BagKey]any),
	}
}

func (b *bag) Set(key constant.BagKey, val any) {
	if b == nil {
		return
	}
	b.dynamics[key] = val
}

func (b *bag) Get(key constant.BagKey) any {
	if b == nil {
		return nil
	}
	return b.dynamics[key]
}
