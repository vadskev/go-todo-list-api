package done

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/vadskev/go_final_project/internal/lib/api"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/lib/utils"
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

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	const op = "transport.handlers.taskItem.HandleGet"
	id := r.URL.Query().Get("id")
	logger.Info(op, zap.String(id, id))
	if len(id) == 0 {
		api.ResponseError(w, r, errors.New("No taskItem id").Error(), http.StatusBadRequest)
		logger.Error(op, zap.Any("No taskItem id, error:", errors.New("No taskItem id").Error()))
		<-h.ctx.Done()
		return
	}

	taskItem, err := h.taskRepository.GetById(id)
	if err != nil {
		api.ResponseError(w, r, errors.New("Task not found").Error(), http.StatusNotFound)
		logger.Error(op, zap.Any("Task not found, error:", errors.New("No taskItem id").Error()))
		<-h.ctx.Done()
		return
	}

	if taskItem.Repeat == "" {
		err = h.taskRepository.Delete(id)
		if err != nil {
			api.ResponseError(w, r, errors.New("Error delete taskItem").Error(), http.StatusInternalServerError)
			logger.Error(op, zap.Any("Error delete taskItem, error:", errors.New("No taskItem id").Error()))
			<-h.ctx.Done()
			return
		}
	} else {
		now := time.Now()
		taskItem.Date, err = utils.NextDate(now, taskItem.Date, taskItem.Repeat)
		if err != nil {
			api.ResponseError(w, r, errors.New("Error repeat format").Error(), http.StatusBadRequest)
			logger.Error(op, zap.Any("Error repeat format:", errors.New("error").Error()))
			<-h.ctx.Done()
			return
		}
		err = h.taskRepository.Update(taskItem)
		if err != nil {
			api.ResponseError(w, r, errors.New("Error update taskItem").Error(), http.StatusInternalServerError)
			logger.Error(op, zap.Any("Error update taskItem:", errors.New("error").Error()))
			<-h.ctx.Done()
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	api.ResponseOK(w, r, task.Response{})
}
