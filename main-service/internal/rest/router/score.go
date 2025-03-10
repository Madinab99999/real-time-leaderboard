package router

import (
	"context"
	"net/http"
)

func (r *Router) score(ctx context.Context) {
	r.router.Handle("POST /scores/new", http.HandlerFunc(r.handler.PostScore))
}
