package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ssonit/clean-architecture/common"
	gin_http "github.com/ssonit/clean-architecture/internal/todo/transport/gin"
	"github.com/ssonit/clean-architecture/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	dsn := "user:pass@tcp(localhost:3307)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.Use(middleware.Recovery())

	v1 := r.Group("/v1/api")
	{
		todo := v1.Group("/todo")
		{

			gin_http.RegisterTodoRoutes(todo, db)

			todo.GET("/ping", func(c *gin.Context) {
				go func() {
					defer common.Recovery()
					fmt.Println([]int{}[0])
				}()

				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}
	}

	r.Run()
}
