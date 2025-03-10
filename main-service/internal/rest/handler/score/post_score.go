package score

import (
	"log/slog"
	"net/http"

	"main_service/internal/rest/pkg/httperror"
	"main_service/pkg/httputils/request"
	"main_service/pkg/httputils/response"
)

func (h *Score) PostScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "PostScore")

	reqBody := new(postScoreReq)
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	_, err := h.ctrl.NewScore(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid values", "", "").
			HandleError(w, err)
		return
	}

	if err := response.JSON(
		w,
		http.StatusCreated,
		nil,
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

type postScoreReq struct {
	Data *postScoreReqData `json:"data"`
}

func (s *postScoreReq) GetUserID() string {
	return s.Data.User_id
}

func (s *postScoreReq) GetGameName() string {
	return s.Data.Game
}

func (s *postScoreReq) GetScore() float64 {
	return s.Data.Score
}

type postScoreReqData struct {
	User_id string  `json:"user_id"`
	Game    string  `json:"game"`
	Score   float64 `json:"score"`
}
