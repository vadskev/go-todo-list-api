package signin

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/vadskev/go_final_project/internal/lib/api"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/lib/utils"
	"github.com/vadskev/go_final_project/internal/storage/db"
	"go.uber.org/zap"
)

type Handler struct {
	taskRepository db.Repository
	ctx            context.Context
}

func New(ctx context.Context, taskRepository db.Repository) *Handler {
	return &Handler{
		taskRepository: taskRepository,
		ctx:            ctx,
	}
}

type SignRequest struct {
	Password string `json:"password"`
}

type SignResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	const op = "transport.handlers.signin.HandlePost"

	var req SignRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		api.ResponseError(w, r, err.Error(), http.StatusBadRequest)
		logger.Error(op, zap.Any("Decode json, error:", err.Error()))
		<-h.ctx.Done()
		return
	}

	password := os.Getenv("TODO_PASSWORD")

	if req.Password != password {
		api.ResponseError(w, r, errors.New("Wrong password").Error(), http.StatusBadRequest)
		logger.Error(op, zap.Any("error:", errors.New("Wrong password").Error()))
		<-h.ctx.Done()
		return
	}

	tokenString := utils.CreateHash(req.Password)
	resp := SignResponse{Token: tokenString}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, resp)
}
