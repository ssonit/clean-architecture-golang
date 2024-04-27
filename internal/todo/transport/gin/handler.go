package gin_http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssonit/clean-architecture/common"
	"github.com/ssonit/clean-architecture/internal/todo/biz"
	"github.com/ssonit/clean-architecture/internal/todo/models"
	"github.com/ssonit/clean-architecture/internal/todo/storage"
	"gorm.io/gorm"
)

func CreateTodo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var todo models.TodoItemCreation

		if err := c.ShouldBind(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := storage.NewSQLStore(db)

		biz := biz.NewTodoBiz(storage)

		if err := biz.Create(c.Request.Context(), &todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(todo.ID))
	}
}

func GetListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var todo []models.TodoItem
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		var filter models.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewTodoBiz(store)

		todo, err := biz.ListTodoItem(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(todo, paging, filter))
	}
}

func GetTodoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewTodoBiz(store)

		data, err := biz.GetItemById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func UpdateTodoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var todo models.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		biz := biz.NewTodoBiz(store)

		if err := biz.UpdateTodo(c.Request.Context(), id, &todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func DeleteTodoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewTodoBiz(store)

		if err := biz.Delete(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
