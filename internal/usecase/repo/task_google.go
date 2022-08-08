package repo

import (
	"context"
	"github.com/esvarez/togo/internal/entity"
	errs "github.com/esvarez/togo/pkg/error"

	"google.golang.org/api/tasks/v1"
)

type TaskRepo struct {
	srv *tasks.Service
}

func NewTaskRepo(srv *tasks.Service) *TaskRepo {
	return &TaskRepo{srv: srv}
}

func (t TaskRepo) NewTask(ctx context.Context, task *entity.Task) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) GetTasks(ctx context.Context, listID string, limit int) ([]*entity.Task, error) {
	resp, err := t.srv.Tasks.List(listID).MaxResults(int64(limit)).Do()
	if err != nil {
		return nil, err
	}

	tasks := make([]*entity.Task, len(resp.Items))
	if len(resp.Items) > 0 {
		for i, item := range resp.Items {
			tasks[i] = &entity.Task{
				ID:          item.Id,
				Name:        item.Title,
				Description: item.Notes,
			}
		}
	} else {
		return nil, errs.ErrNotFound
	}
	return tasks, nil
}

func (t TaskRepo) GetTask(ctx context.Context, s string, s2 string) (*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) EditTask(ctx context.Context, task *entity.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) CompleteTask(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) DeleteTask(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}
