package leaderboard

import (
	"log/slog"
	"net/http"

	"main_service/internal/rest/pkg/httperror"
	"main_service/pkg/httputils/response"
)

func (h *Leaderboard) GetUserRank(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetUserRank")

	id_user := r.PathValue("id")
	reqBody := newGetUserRankReq(id_user)
	ctrlResp, err := h.ctrl.GetUserRank(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid values", "", "").
			HandleError(w, err)
		return
	}
	var games []GameRank
	for _, v := range ctrlResp.GetScoreList() {
		if v != nil {
			games = append(games, *newGameRank(v.GetRank(), v.GetGameName(), v.GetScore()))
		}
	}
	respBody := newGetUserRankResp(*newUserRank(ctrlResp.GetGlobalRank(), ctrlResp.GetGlobalScore(), games))
	if err := response.JSON(
		w,
		http.StatusOK,
		respBody,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			slog.Any("error", err),
		)
		return
	}

	log.InfoContext(
		ctx,
		"success",
	)
}

type getUserRankReq struct {
	userId string
}

func newGetUserRankReq(userId string) *getUserRankReq {
	return &getUserRankReq{
		userId: userId,
	}
}

func (r *getUserRankReq) GetUserID() string {
	return r.userId
}

type GetUserRankResp struct {
	Data UserRank `json:"data"`
}

func newGetUserRankResp(data UserRank) *GetUserRankResp {
	return &GetUserRankResp{
		Data: data,
	}
}

type UserRank struct {
	GlobalRank  int        `json:"global_rank"`
	GlobalScore float64    `json:"global_score"`
	GamesRank   []GameRank `json:"games_rank"`
}

func newUserRank(global_rank int, global_score float64, gamesrank []GameRank) *UserRank {
	return &UserRank{
		GlobalRank:  global_rank,
		GlobalScore: global_score,
		GamesRank:   gamesrank,
	}
}

type GameRank struct {
	Rank  int     `json:"rank"`
	Game  string  `json:"game"`
	Score float64 `json:"score"`
}

func newGameRank(rank int, game string, score float64) *GameRank {
	return &GameRank{
		Rank:  rank,
		Game:  game,
		Score: score,
	}
}
