package tasks

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vadskev/go_final_project/internal/lib/api"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/models/task"
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

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	const op = "transport.tasks.Handle"
	searchStr := chi.URLParam(r, "search")
	tasks, err := h.taskRepository.GetTasks(searchStr)
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		api.ResponseError(w, r, err.Error(), http.StatusInternalServerError)
		<-h.ctx.Done()
		return
	}
	response := struct {
		Tasks []task.Task `json:"tasks"`
	}{Tasks: tasks}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}
