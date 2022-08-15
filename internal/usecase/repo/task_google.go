package repo

import (
	"context"
	"github.com/esvarez/togo/internal/entity"
	errs "github.com/esvarez/togo/pkg/error"
	"google.golang.org/api/tasks/v1"
)

const (
	completed = "completed"
)

type TaskRepo struct {
	srv *tasks.Service
}

func NewTaskRepo(srv *tasks.Service) *TaskRepo {
	return &TaskRepo{srv: srv}
}

func (t TaskRepo) NewTask(ctx context.Context, listID string, task *entity.Task) (string, error) {
	resp, err := t.srv.Tasks.Insert(listID, &tasks.Task{
		Notes: task.Description,
		Title: task.Name,
		Due:   task.DueDate,
	}).Do()
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (t TaskRepo) GetTasks(ctx context.Context, listID string, limit int) ([]*entity.Task, error) {
	resp, err := t.srv.Tasks.List(listID).MaxResults(int64(limit)).Do()
	if err != nil {
		return nil, err
	}

	tasks := make([]*entity.Task, 0)

	if len(resp.Items) > 0 {
		for _, item := range resp.Items {
			if item.Status != completed {
				tasks = append(tasks, &entity.Task{
					ID:          item.Id,
					Name:        item.Title,
					Description: item.Notes,
				})
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

func (t TaskRepo) CompleteTask(ctx context.Context, listID, taskID string) error {
	item, err := t.srv.Tasks.Get(listID, taskID).Do()
	if err != nil {
		return err
	}
	item.Status = completed

	_, err = t.srv.Tasks.Update(listID, taskID, item).Do()
	if err != nil {
		return err
	}
	return nil
}

func (t TaskRepo) DeleteTask(ctx context.Context, listID, taskID string) error {
	err := t.srv.Tasks.Delete(listID, taskID).Do()
	if err != nil {
		return err
	}
	return nil
}
