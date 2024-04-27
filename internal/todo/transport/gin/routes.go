package gin_http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTodoRoutes(group *gin.RouterGroup, db *gorm.DB) {
	group.POST("/", CreateTodo(db))
	group.GET("/", GetListItem(db))
	group.GET("/:id", GetTodoItem(db))
}
