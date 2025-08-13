package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"task-api/internal/domain"
	"task-api/internal/logger"
	"task-api/internal/repository"
)

type TaskHandler struct {
	repo   *repository.InMemoryRepo
	logger *logger.Logger
}

func NewTaskHandler(repo *repository.InMemoryRepo, logger *logger.Logger) *TaskHandler {
	return &TaskHandler{repo: repo, logger: logger}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	statusStr := r.URL.Query().Get("status")
	var status *domain.Status
	if statusStr != "" {
		s := domain.Status(statusStr)
		status = &s
	}
	tasks := h.repo.GetAll(status)
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, ok := h.repo.GetByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	h.repo.Create(task)
	h.logger.Log("Created task: " + task.ID)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}
