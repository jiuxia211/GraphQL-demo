package middleware

import (
	"context"
	"net/http"

	"github.com/jiuxia211/GraphQL-demo/pkg/utils"
)

type contextKey string

const userCtxKey contextKey = "user"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// 可以不验证
			next.ServeHTTP(w, r)
			return
		}

		claims, err := utils.CheckToken(authHeader)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func ForContext(ctx context.Context) *utils.Claims {
	raw, _ := ctx.Value(userCtxKey).(*utils.Claims)
	return raw
}
