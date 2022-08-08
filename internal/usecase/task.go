package usecase

import (
	"context"
	"github.com/esvarez/togo/internal/entity"
)

type taskRepo interface {
	NewTask(context.Context, *entity.Task) (string, error)
	GetTasks(context.Context, int) ([]*entity.Task, error)
	GetTask(context.Context, string, string) (*entity.Task, error)
	EditTask(context.Context, *entity.Task) error
	CompleteTask(context.Context, string) error
	DeleteTask(context.Context, string) error
}

type Task struct {
	repo taskRepo
}

func NewTask(repo taskRepo) *Task {
	return &Task{repo: repo}
}

func (t *Task) Create(ctx context.Context, task *entity.Task) (string, error) {
	return t.repo.NewTask(ctx, task)
}

func (t *Task) Edit(ctx context.Context, task *entity.Task) error {
	return t.repo.EditTask(ctx, task)
}

func (t *Task) Complete(ctx context.Context, id string) error {
	return t.repo.CompleteTask(ctx, id)
}

func (t *Task) Delete(ctx context.Context, id string) error {
	return t.repo.DeleteTask(ctx, id)
}

func (t *Task) List(ctx context.Context, limit int) ([]*entity.Task, error) {
	return t.repo.GetTasks(ctx, limit)
}
