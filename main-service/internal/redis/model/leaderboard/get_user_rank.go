package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"main_service/internal/types/database"

	"github.com/redis/go-redis/v9"
)

func (m *Leaderboard) GetUserRank(ctx context.Context, req database.GetUserRankReq) (database.GetUserRankResp, error) {
	log := m.logger.With(slog.String("handler", "GetUserRank"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}
	globalKey := "leaderboard:global"
	globalScore, err := m.db.ZScore(ctx, globalKey, req.GetUserID()).Result()
	if err != nil && err != redis.Nil {
		log.ErrorContext(ctx, "failed to get global score", slog.Any("error", err))
		return nil, err
	}
	globalRank, err := m.db.ZRevRank(ctx, globalKey, req.GetUserID()).Result()
	if err != nil && err != redis.Nil {
		log.ErrorContext(ctx, "failed to get global rank", slog.Any("error", err))
		return nil, err

	}
	if globalRank != 0 {
		globalRank = globalRank + 1
	}
	gameNames, err := m.db.SMembers(ctx, "games").Result()
	if err != nil {
		log.ErrorContext(ctx, "failed to get game list", slog.Any("error", err))
		return nil, err
	}
	var listRankOfGamesDb []listGamesDbResp
	for _, game := range gameNames {
		gameLeaderboardKey := fmt.Sprintf("leaderboard:game:%s", game)

		gameScore, err := m.db.ZScore(ctx, gameLeaderboardKey, req.GetUserID()).Result()
		if err != nil && err != redis.Nil {
			log.ErrorContext(ctx, "failed to get game score", slog.Any("error", err))
			return nil, err
		}

		gameRank, err := m.db.ZRevRank(ctx, gameLeaderboardKey, req.GetUserID()).Result()
		if err != nil && err != redis.Nil {
			log.ErrorContext(ctx, "failed to get game rank", slog.Any("error", err))
			return nil, err
		}
		if gameRank != 0 {
			gameRank = gameRank + 1
			listRankOfGamesDb = append(listRankOfGamesDb, *newListGamesDbResp(int(gameRank), game, gameScore))
		}
	}
	log.InfoContext(
		ctx,
		"success",
	)
	return newListUserRankDbResp(int(globalRank), globalScore, listRankOfGamesDb), nil
}

type listUserRankDbResp struct {
	globalRank  int
	globalScore float64
	listGames   []listGamesDbResp
}

func newListUserRankDbResp(global_rank int, global_score float64, listGames []listGamesDbResp) *listUserRankDbResp {
	return &listUserRankDbResp{
		globalRank:  global_rank,
		globalScore: global_score,
		listGames:   listGames,
	}
}

func (resp *listUserRankDbResp) GetGlobalRank() int {
	return resp.globalRank
}

func (resp *listUserRankDbResp) GetGlobalScore() float64 {
	return resp.globalScore
}

func (resp *listUserRankDbResp) GetScoreList() []database.ScoresOfGames {
	listRankOfGames := make([]database.ScoresOfGames, len(resp.listGames))
	for _, mg := range resp.listGames {
		var rankOfGame database.ScoresOfGames = newListGamesDbResp(mg.rank, mg.game, mg.score)
		listRankOfGames = append(listRankOfGames, rankOfGame)
	}
	return listRankOfGames
}

type listGamesDbResp struct {
	rank  int
	game  string
	score float64
}

func newListGamesDbResp(rank int, game string, score float64) *listGamesDbResp {
	return &listGamesDbResp{
		rank:  rank,
		game:  game,
		score: score,
	}
}

func (g *listGamesDbResp) GetRank() int {
	return g.rank
}

func (g *listGamesDbResp) GetGameName() string {
	return g.game
}

func (g *listGamesDbResp) GetScore() float64 {
	return g.score
}
