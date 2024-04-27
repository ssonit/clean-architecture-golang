package models

import (
	"database/sql/driver"
	"strings"

	"github.com/ssonit/clean-architecture/common"
	error_todo "github.com/ssonit/clean-architecture/internal/todo"
)

type TodoItemStatus int

const (
	TodoItemStatusDoing TodoItemStatus = iota
	TodoItemStatusDone
	TodoItemStatusDeleted
)

var allTodoItemStatus = []string{"Doing", "Done", "Deleted"}

func parseStringToTodoItemStatus(status string) (TodoItemStatus, error) {
	for i := range allTodoItemStatus {
		if allTodoItemStatus[i] == status {
			return TodoItemStatus(i), nil
		}
	}

	return TodoItemStatus(0), error_todo.ErrInvalidTodoItemStatus
}

func (s *TodoItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return error_todo.ErrScanData
	}

	v, err := parseStringToTodoItemStatus(string(bytes))

	if err != nil {
		return err
	}

	*s = v

	return nil
}

func (s *TodoItemStatus) Value() (driver.Value, error) {
	return s.String(), nil
}

func (s *TodoItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	v, err := parseStringToTodoItemStatus(str)
	if err != nil {
		return err
	}

	*s = v

	return nil

}

func (s *TodoItemStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *TodoItemStatus) String() string {
	return allTodoItemStatus[*s]
}

type TodoItem struct {
	common.SQLModel
	Title       string          `json:"title" gorm:"column:title"`
	Description string          `json:"description" gorm:"column:description"`
	Status      *TodoItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemCreation struct {
	ID          int             `json:"-" gorm:"column:id;primaryKey"`
	Title       string          `json:"title" gorm:"column:title"`
	Description string          `json:"description" gorm:"column:description"`
	Status      *TodoItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       *string         `json:"title" gorm:"column:title"`
	Description *string         `json:"description" gorm:"column:description"`
	Status      *TodoItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todos"
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
