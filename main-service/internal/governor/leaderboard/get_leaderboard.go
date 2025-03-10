package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"main_service/internal/types/controller"
)

func (r *Leaderboard) GetLeaderboard(ctx context.Context, req controller.GetLeaderboardReq) (controller.GetLeaderboardResp, error) {
	log := r.logger.With(slog.String("handler", "GetLeaderboard"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	dbReq := newGetLeaderboardDBReq(req.GetTopCount(), req.GetGameName())
	dbResp, err := r.db.GetTopPlayers(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("db request failed %w", err)
	}
	if dbResp == nil {
		return nil, nil
	}
	var players []playersResp
	for _, v := range dbResp.GetList() {
		if v != nil {
			players = append(players, *newPlayersResp(v.GetRank(), v.GetUserID(), v.GetScore()))
		}
	}
	log.InfoContext(
		ctx,
		"success",
	)
	return newListPlayersResp(players), nil
}

type getLeaderboardDBReq struct {
	topCount int
	game     *string
}

func newGetLeaderboardDBReq(topCount int, game *string) *getLeaderboardDBReq {
	return &getLeaderboardDBReq{
		topCount: topCount,
		game:     game,
	}
}

func (req *getLeaderboardDBReq) GetTopCount() int {
	return req.topCount
}

func (req *getLeaderboardDBReq) GetGameName() *string {
	if req.game == nil {
		return nil
	}
	return req.game
}

type listPlayersResp struct {
	list []playersResp
}

func newListPlayersResp(list []playersResp) *listPlayersResp {
	return &listPlayersResp{list: list}
}

func (resp *listPlayersResp) GetList() []controller.ItemGetLeaderboardResp {
	listPlayers := make([]controller.ItemGetLeaderboardResp, len(resp.list))
	for _, mg := range resp.list {
		var player controller.ItemGetLeaderboardResp = newPlayersResp(mg.rank, mg.userId, mg.score)
		listPlayers = append(listPlayers, player)
	}
	return listPlayers
}

type playersResp struct {
	rank   int
	userId string
	score  float64
}

func newPlayersResp(rank int, userId string, score float64) *playersResp {
	return &playersResp{
		rank:   rank,
		userId: userId,
		score:  score,
	}
}

func (pl *playersResp) GetRank() int {
	return pl.rank
}

func (pl *playersResp) GetUserID() string {
	return pl.userId
}

func (pl *playersResp) GetScore() float64 {
	return pl.score
}
