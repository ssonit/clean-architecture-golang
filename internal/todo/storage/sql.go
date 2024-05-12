package storage

import (
	"context"

	"github.com/ssonit/clean-architecture/common"
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
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) GetItem(ctx context.Context, filter map[string]interface{}) (*models.TodoItem, error) {
	var todo models.TodoItem
	if err := s.db.Where(filter).First(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}

	return &todo, nil
}

func (s *sqlStore) Update(ctx context.Context, filter map[string]interface{}, todo *models.TodoItemUpdate) error {
	if err := s.db.Where(filter).Updates(&todo).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) Delete(ctx context.Context, filter map[string]interface{}) error {

	deletedStatus := models.TodoItemStatusDeleted

	if err := s.db.Table(models.TodoItem{}.TableName()).Where(filter).Updates(map[string]interface{}{
		"status": deletedStatus.String(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) ListTodoItem(ctx context.Context, filter *models.Filter, paging *common.Paging, moreKeys ...string) ([]models.TodoItem, error) {
	var todo []models.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(models.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Table(models.TodoItem{}.TableName()).Order("id desc").Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
