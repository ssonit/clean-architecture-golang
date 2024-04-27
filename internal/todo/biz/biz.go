package biz

import (
	"context"
	"strings"

	"github.com/ssonit/clean-architecture/common"
	error_todo "github.com/ssonit/clean-architecture/internal/todo"
	"github.com/ssonit/clean-architecture/internal/todo/models"
)

type TodoStorage interface {
	Create(ctx context.Context, todo *models.TodoItemCreation) error
	GetItem(ctx context.Context, filter map[string]interface{}) (*models.TodoItem, error)
	Update(ctx context.Context, filter map[string]interface{}, todo *models.TodoItemUpdate) error
	Delete(ctx context.Context, filter map[string]interface{}) error
	ListTodoItem(ctx context.Context, filter *models.Filter, paging *common.Paging, moreKeys ...string) ([]models.TodoItem, error)
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

func (biz *todoBiz) UpdateTodo(ctx context.Context, id int, todo *models.TodoItemUpdate) error {
	filter := map[string]interface{}{"id": id}

	data, err := biz.store.GetItem(ctx, filter)

	if err != nil {
		return err
	}

	if data.Status == nil || *data.Status == models.TodoItemStatusDeleted {
		return error_todo.ErrTodoItemIsDeleted
	}

	if err := biz.store.Update(ctx, filter, todo); err != nil {
		return err
	}
	return nil
}

func (biz *todoBiz) Delete(ctx context.Context, id int) error {
	filter := map[string]interface{}{"id": id}

	data, err := biz.store.GetItem(ctx, filter)

	if err != nil {
		return err
	}

	if data.Status == nil || *data.Status == models.TodoItemStatusDeleted {
		return error_todo.ErrTodoItemIsDeleted
	}

	if err := biz.store.Delete(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (biz *todoBiz) ListTodoItem(ctx context.Context, filter *models.Filter, paging *common.Paging) ([]models.TodoItem, error) {
	data, err := biz.store.ListTodoItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
