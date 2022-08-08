package repo

import (
	"context"
	"github.com/esvarez/togo/internal/entity"
	errs "github.com/esvarez/togo/pkg/error"
	"google.golang.org/api/tasks/v1"
)

type ListRepo struct {
	srv *tasks.Service
}

func (l *ListRepo) NewList(ctx context.Context, list *entity.TaskList) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewListRepo(srv *tasks.Service) *ListRepo {
	return &ListRepo{srv: srv}
}

func (l *ListRepo) GetLists(ctx context.Context) ([]*entity.TaskList, error) {
	resp, err := l.srv.Tasklists.List().Do()
	if err != nil {
		return nil, err
	}
	lists := make([]*entity.TaskList, len(resp.Items))
	if len(resp.Items) > 0 {
		for i, item := range resp.Items {
			lists[i] = &entity.TaskList{
				ID:    item.Id,
				Title: item.Title,
			}
		}
	} else {
		return nil, errs.ErrNotFound
	}
	return lists, nil
}
