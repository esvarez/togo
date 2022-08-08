package usecase

import (
	"context"

	"github.com/esvarez/togo/internal/entity"
)

type listRepo interface {
	// NewList(ctx context.Context, list *entity.TaskList) (string, error)
	GetLists(ctx context.Context) ([]*entity.TaskList, error)
}

type List struct {
	repo listRepo
}

func NewList(repo listRepo) *List {
	return &List{repo: repo}
}

func (l List) List(ctx context.Context) ([]*entity.TaskList, error) {
	return l.repo.GetLists(ctx)
}
