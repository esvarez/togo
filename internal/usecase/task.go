package usecase

import (
	"context"
	"github.com/esvarez/togo/config"
	"github.com/esvarez/togo/internal/entity"
)

type taskRepo interface {
	NewTask(context.Context, string, *entity.Task) (string, error)
	GetTasks(context.Context, string, int) ([]*entity.Task, error)
	GetTask(context.Context, string, string) (*entity.Task, error)
	EditTask(context.Context, *entity.Task) error
	CompleteTask(context.Context, string, string) error
	DeleteTask(context.Context, string, string) error
}

type Task struct {
	app  *config.App
	repo taskRepo
}

func NewTask(repo taskRepo) *Task {
	return &Task{repo: repo}
}

func (t *Task) Create(ctx context.Context, listID string, task *entity.Task) (string, error) {
	return t.repo.NewTask(ctx, listID, task)
}

func (t *Task) Edit(ctx context.Context, task *entity.Task) error {
	return t.repo.EditTask(ctx, task)
}

func (t *Task) Complete(ctx context.Context, listID, id string) error {
	return t.repo.CompleteTask(ctx, listID, id)
}

func (t *Task) Delete(ctx context.Context, idList, id string) error {
	return t.repo.DeleteTask(ctx, idList, id)
}

func (t *Task) List(ctx context.Context, listID string, limit int) ([]*entity.Task, error) {
	return t.repo.GetTasks(ctx, listID, limit)
}
