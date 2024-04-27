package biz

import (
	"context"
	"strings"

	error_todo "github.com/ssonit/clean-architecture/internal/todo"
	"github.com/ssonit/clean-architecture/internal/todo/models"
)

type TodoStorage interface {
	Create(ctx context.Context, todo *models.TodoItemCreation) error
	GetItem(ctx context.Context, filter map[string]interface{}) (*models.TodoItem, error)
}

type todoBiz struct {
	store TodoStorage
}

func NewTodoBiz(store TodoStorage) *todoBiz {
	return &todoBiz{
		store: store,
	}
}

func (biz *todoBiz) Create(ctx context.Context, todo *models.TodoItemCreation) error {
	title := strings.TrimSpace(todo.Title)

	if title == "" {
		return error_todo.ErrTitleIsBlank
	}

	if err := biz.store.Create(ctx, todo); err != nil {
		return err
	}

	return nil
}

func (biz *todoBiz) GetItemById(ctx context.Context, id int) (*models.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
