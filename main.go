package main

import (
	"fmt"
	"golangtesting/internal/activity/handler"
	"golangtesting/internal/activity/repository"
	"golangtesting/internal/activity/service"
	"golangtesting/internal/database"
	"golangtesting/internal/todo/handler_todo"
	"golangtesting/internal/todo/repository_todo"
	"golangtesting/internal/todo/service_todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading env ")
	}
	router := gin.Default()
	db := database.Connetion()
	ActivityRepo := repository.NewActivityRepository(db)
	ActivitySer := service.NewActivityService(ActivityRepo)
	ActivityHand := handler.NewActivityHandler(ActivitySer)

	router.GET("activity-groups", ActivityHand.Getdata)
	router.GET("activity-groups/:id", ActivityHand.GetdataById)
	router.POST("activity-groups", ActivityHand.Createdata)
	router.DELETE("activity-groups/:id", ActivityHand.DeleteData)
	router.PATCH("activity-groups/:id", ActivityHand.UpdateData)

	//modul todo
	TodoRepo := repository_todo.NewTodoRepository(db)
	TodoSer := service_todo.NewTodoService(TodoRepo)
	TodoHand := handler_todo.NewTodoHandler(TodoSer)

	router.GET("todo-items", TodoHand.Getdata)
	router.GET("/todo-items/:id", TodoHand.GetdataById)
	router.POST("/todo-items", TodoHand.Createdata)
	router.DELETE("/todo-items/:id", TodoHand.Deletedata)
	router.PATCH("/todo-items/:id", TodoHand.Updatedata)

	router.Run(":8090")

}
