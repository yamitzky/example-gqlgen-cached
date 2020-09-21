// better solution: https://github.com/99designs/gqlgen/pull/1048
package expiration

import (
	"context"
	"net/http"
)

var ctxKey = &contextKey{"expire"}

type contextKey struct {
	name string
}

type expiration struct {
	time int64
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// put it in context
		exp := expiration{}
		ctx := context.WithValue(r.Context(), ctxKey, &exp)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetExpire(ctx context.Context) *expiration {
	currentExpire, _ := ctx.Value(ctxKey).(*expiration)
	return currentExpire
}

func SetExpire(ctx context.Context, exp int64) {
	currentExpire := GetExpire(ctx)
	if currentExpire.time == 0 || exp < currentExpire.time {
		currentExpire.time = exp
	}
}
