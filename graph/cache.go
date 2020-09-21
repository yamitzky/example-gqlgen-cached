package graph

import (
	"context"
	"io/ioutil"

	"github.com/google/martian/v3/log"
	lru "github.com/hashicorp/golang-lru"
)

type QueryCache struct {
	lru *lru.Cache
}

func NewQueryCache() *QueryCache {
	cache, err := lru.New(1024)
	if err != nil {
		panic(err)
	}
	return &QueryCache{cache}
}

func (c *QueryCache) Get(ctx context.Context, key string) (value interface{}, ok bool) {
	if r, ok := c.lru.Get(key); ok {
		return r, ok
	}
	bytes, err := ioutil.ReadFile("./graph/queries/" + key + ".graphql")
	if err != nil {
		return nil, false
	}
	query := string(bytes)
	c.lru.Add(key, query)
	return c.lru.Get(key)
}

func (c *QueryCache) Add(ctx context.Context, key string, value interface{}) {
	// do nothing, warn it
	log.Errorf("Unexpected APQ registration: %s", value)
}
