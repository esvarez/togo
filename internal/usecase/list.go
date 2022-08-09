package usecase

import (
	"context"
	"github.com/esvarez/togo/config"

	"github.com/esvarez/togo/internal/entity"
)

type listRepo interface {
	NewList(ctx context.Context, list *entity.TaskList) (string, error)
	GetLists(ctx context.Context) ([]*entity.TaskList, error)
}

type List struct {
	app  *config.App
	repo listRepo
}

func (l List) Create(ctx context.Context, list *entity.TaskList) (string, error) {
	return l.repo.NewList(ctx, list)
}

func NewList(repo listRepo) *List {
	return &List{repo: repo}
}

func (l List) List(ctx context.Context) ([]*entity.TaskList, error) {
	return l.repo.GetLists(ctx)
}
