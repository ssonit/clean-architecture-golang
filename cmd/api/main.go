package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gin_http "github.com/ssonit/clean-architecture/internal/todo/transport/gin"
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

	v1 := r.Group("/v1/api")
	{
		todo := v1.Group("/todo")
		{

			gin_http.RegisterTodoRoutes(todo, db)
		}
	}

	r.Run()
}
