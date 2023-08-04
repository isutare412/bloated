package contextbag

import "context"

type ctxKeyBag struct{}

func WithBag(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyBag{}, newBag())
}

func Bag(ctx context.Context) *bag {
	return ctx.Value(ctxKeyBag{}).(*bag)
}

type bag struct {
	dynamics map[string]any
}

func newBag() *bag {
	return &bag{
		dynamics: make(map[string]any),
	}
}

func (b *bag) Set(key string, val any) {
	if b == nil {
		return
	}
	b.dynamics[key] = val
}

func (b *bag) Get(key string) any {
	if b == nil {
		return nil
	}
	return b.dynamics[key]
}
