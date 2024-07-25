package task

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/vadskev/go_final_project/internal/models"
)

func NewTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var tasks []models.Task
		var t models.Task
		t.ID = "1"
		t.Date = "20240131"
		t.Title = "Фитнес"
		t.Comment = "Мой комментарий"
		t.Repeat = "d 5"

		tasks = append(tasks, t)

		response := struct {
			Tasks []models.Task `json:"tasks"`
		}{Tasks: tasks}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, response)
	}
}
