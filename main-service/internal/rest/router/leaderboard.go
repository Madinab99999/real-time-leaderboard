package router

import (
	"context"
	"net/http"
)

func (r *Router) leaderboard(ctx context.Context) {
	// получение таблицы лидеров
	r.router.Handle("GET /leaderboard", http.HandlerFunc(r.handler.GetLeaderboard))
	//получение ранга пользователя
	r.router.Handle("GET /leaderboard/rank/{id}", http.HandlerFunc(r.handler.GetUserRank))

	//r.router.Handle("GET /leaderboard/report", )
}
