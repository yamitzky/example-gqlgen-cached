package expiration

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	lru "github.com/hashicorp/golang-lru"
)

type Extension struct {
	lru *lru.Cache
}

func NewExtension() *Extension {
	cache, err := lru.New(1024)
	if err != nil {
		panic(err)
	}
	return &Extension{
		lru: cache,
	}
}

func (e Extension) ExtensionName() string {
	return "Expiration"
}

func (e Extension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (e Extension) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	path := graphql.GetPath(ctx).String() // should vary depending on params
	if v, ok := e.lru.Get(path); ok {
		return v, nil
	}

	expire := GetExpire(ctx)
	before := len(expire.times)
	res, err = next(ctx)
	expire = GetExpire(ctx)
	after := len(expire.times)

	if err == nil && after > before {
		last := expire.times[len(expire.times)-1]
		log.Printf("TODO: must expire in %d secs", last)
		e.lru.Add(path, res)
	}
	return res, err
}
