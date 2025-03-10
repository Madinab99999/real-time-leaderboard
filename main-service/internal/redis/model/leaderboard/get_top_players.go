package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"main_service/internal/types/database"
)

func (m *Leaderboard) GetTopPlayers(ctx context.Context, req database.GetTopPlayersReq) (database.GetTopPlayersResp, error) {
	log := m.logger.With(slog.String("handler", "GetTopPlayers"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}
	key := "leaderboard:global"
	if *req.GetGameName() != "" {
		key = fmt.Sprintf("leaderboard:game:%s", *req.GetGameName())
	}
	topPlayers, err := m.db.ZRevRangeWithScores(ctx, key, 0, int64(req.GetTopCount())-1).Result()
	if err != nil {
		log.ErrorContext(ctx, "failed get top players", slog.Any("error", err))
	}
	var playersDb []playerDbResp
	for i, v := range topPlayers {
		playersDb = append(playersDb, *newPlayerDbResp(i+1, v.Member.(string), v.Score))
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return newGetTopPlayersResp(playersDb), nil
}

type getTopPlayersResp struct {
	list []playerDbResp
}

func newGetTopPlayersResp(list []playerDbResp) *getTopPlayersResp {
	return &getTopPlayersResp{list: list}
}

func (resp *getTopPlayersResp) GetList() []database.ItemGetTopPlayersResp {
	listPlayers := make([]database.ItemGetTopPlayersResp, len(resp.list))
	for _, mg := range resp.list {
		var player database.ItemGetTopPlayersResp = newPlayerDbResp(mg.rank, mg.userId, mg.score)
		listPlayers = append(listPlayers, player)
	}
	return listPlayers
}

type playerDbResp struct {
	rank   int
	userId string
	score  float64
}

func newPlayerDbResp(rank int, userId string, score float64) *playerDbResp {
	return &playerDbResp{
		rank:   rank,
		userId: userId,
		score:  score,
	}
}

func (pl *playerDbResp) GetRank() int {
	return pl.rank
}

func (pl *playerDbResp) GetUserID() string {
	return pl.userId
}

func (pl *playerDbResp) GetScore() float64 {
	return pl.score
}
