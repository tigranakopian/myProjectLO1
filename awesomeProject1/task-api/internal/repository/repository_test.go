package repository

import (
	"task-api/internal/domain"
	"testing"
)

func TestCreateAndGetByID(t *testing.T) {
	repo := NewInMemoryRepo()
	task := domain.Task{ID: "123", Title: "Тест", Status: "pending"}

	repo.Create(task)
	got, ok := repo.GetByID("123")

	if !ok || got.ID != "123" {
		t.Errorf("expected task with ID 123, got %+v", got)
	}
}
