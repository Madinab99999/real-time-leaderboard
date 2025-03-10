package leaderboard

import (
	"log/slog"
	"net/http"
	"strconv"

	"main_service/internal/rest/pkg/httperror"
	"main_service/pkg/httputils/response"
)

func (h *Leaderboard) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetLeaderboard")

	game := r.URL.Query().Get("game")
	top := r.URL.Query().Get("top")
	if top == "" {
		top = "10"
	}
	topN, err := strconv.Atoi(top)
	if err != nil {
		log.ErrorContext(
			ctx, "failed to parse string to int",
			slog.Any("error", err))
		return
	}
	reqBody := newGetLeaderboardReq(topN, &game)
	ctrlResp, err := h.ctrl.GetLeaderboard(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid values", "", "").
			HandleError(w, err)
		return
	}
	var players []*PlayersResp
	for _, pl := range ctrlResp.GetList() {
		if pl != nil {
			players = append(players, &PlayersResp{
				Rank:   pl.GetRank(),
				UserId: pl.GetUserID(),
				Score:  pl.GetScore(),
			})
		}
	}
	respBody := newGetLeaderboardResp(players)
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

type getLeaderboardReq struct {
	topCount int
	game     *string
}

func newGetLeaderboardReq(topCount int, game *string) *getLeaderboardReq {
	return &getLeaderboardReq{
		topCount: topCount,
		game:     game,
	}
}

func (s *getLeaderboardReq) GetTopCount() int {
	return s.topCount
}

func (s *getLeaderboardReq) GetGameName() *string {
	if s.game == nil {
		return nil
	}
	return s.game
}

type GetLeaderboardResp struct {
	Data []*PlayersResp `json:"data"`
}

func newGetLeaderboardResp(data []*PlayersResp) *GetLeaderboardResp {
	return &GetLeaderboardResp{
		Data: data,
	}
}

type PlayersResp struct {
	Rank   int     `json:"rank"`
	UserId string  `json:"user_id"`
	Score  float64 `json:"score"`
}
