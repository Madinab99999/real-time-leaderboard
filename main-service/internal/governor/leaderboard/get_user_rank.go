package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"main_service/internal/types/controller"
)

func (r *Leaderboard) GetUserRank(ctx context.Context, req controller.GetUserRankReq) (controller.GetUserRankResp, error) {
	log := r.logger.With(slog.String("handler", "GetUserRank"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	dbReq := newGetUserRankDBReq(req.GetUserID())
	dbResp, err := r.db.GetUserRank(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("db request failed %w", err)
	}
	if dbResp == nil {
		return nil, nil
	}
	var rankOfGames []listGamesResp
	for _, r := range dbResp.GetScoreList() {
		if r != nil {
			rankOfGames = append(rankOfGames, *newListGamesResp(r.GetRank(), r.GetGameName(), r.GetScore()))
		}
	}
	log.InfoContext(
		ctx,
		"success",
	)
	return newListUserRankResp(dbResp.GetGlobalRank(), dbResp.GetGlobalScore(), rankOfGames), nil
}

type getUserRankDBReq struct {
	userId string
}

func newGetUserRankDBReq(userId string) *getUserRankDBReq {
	return &getUserRankDBReq{
		userId: userId,
	}
}

func (req *getUserRankDBReq) GetUserID() string {
	return req.userId
}

type listUserRankResp struct {
	globalRank  int
	globalScore float64
	listGames   []listGamesResp
}

func newListUserRankResp(global_rank int, global_score float64, listGames []listGamesResp) *listUserRankResp {
	return &listUserRankResp{
		globalRank:  global_rank,
		globalScore: global_score,
		listGames:   listGames,
	}
}

func (resp *listUserRankResp) GetGlobalRank() int {
	return resp.globalRank
}

func (resp *listUserRankResp) GetGlobalScore() float64 {
	return resp.globalScore
}

func (resp *listUserRankResp) GetScoreList() []controller.ScoresOfGames {
	listRankOfGames := make([]controller.ScoresOfGames, len(resp.listGames))
	for _, mg := range resp.listGames {
		//var rankOfGame controller.ScoresOfGames = newListGamesResp(mg.rank, mg.game, mg.score)
		listRankOfGames = append(listRankOfGames, newListGamesResp(mg.rank, mg.game, mg.score))
	}
	return listRankOfGames
}

type listGamesResp struct {
	rank  int
	game  string
	score float64
}

func newListGamesResp(rank int, game string, score float64) *listGamesResp {
	return &listGamesResp{
		rank:  rank,
		game:  game,
		score: score,
	}
}

func (g *listGamesResp) GetRank() int {
	return g.rank
}

func (g *listGamesResp) GetGameName() string {
	return g.game
}

func (g *listGamesResp) GetScore() float64 {
	return g.score
}
