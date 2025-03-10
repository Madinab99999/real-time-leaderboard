package router

import (
	"context"
	"net/http"

	"main_service/internal/rest/handler"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
}

func New(handler *handler.Handler) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {

	r.leaderboard(ctx)
	r.score(ctx)

	return r.router
}
