package storage

import (
	"context"

	"github.com/ssonit/clean-architecture/internal/todo/models"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Create(ctx context.Context, todo *models.TodoItemCreation) error {
	if err := s.db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) GetItem(ctx context.Context, filter map[string]interface{}) (*models.TodoItem, error) {
	var todo models.TodoItem
	if err := s.db.Where(filter).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
