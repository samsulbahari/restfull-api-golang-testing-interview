package handler_todo

import (
	"golangtesting/internal/domain"

	"github.com/gin-gonic/gin"
)

type TodoService interface {
	GetDataService(ctx *gin.Context) ([]domain.Todo, int, error)
	GetdataByidService(ctx *gin.Context) (domain.Todo, int, error)
	CreateService(todo domain.Todo, ctx *gin.Context) (domain.Todo, int, error)
	DeleteService(ctx *gin.Context) (domain.Todo, int, error)
	UpdateService(ctx *gin.Context, todoupdate domain.TodoUpdate) (domain.TodoUpdate, int, error)
}

type TodoHandler struct {
	TodoHan TodoService
}

func NewTodoHandler(ts TodoService) *TodoHandler {
	return &TodoHandler{ts}
}

func (th TodoHandler) Getdata(ctx *gin.Context) {
	todo, code, err := th.TodoHan.GetDataService(ctx)

	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func (th TodoHandler) GetdataById(ctx *gin.Context) {
	todo, code, err := th.TodoHan.GetdataByidService(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Not Found",
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})

}

func (th TodoHandler) Createdata(ctx *gin.Context) {
	var todo domain.Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		if todo.Title == "" {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "title cannot be null",
				"data":    nil,
			})
			return
		}

		if todo.IdActivity == "" {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "activity_group_id cannot be null",
				"data":    nil,
			})
			return
		}
	}

	todo, code, err := th.TodoHan.CreateService(todo, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Not Found",
			"message": err.Error(),
			"data":    err,
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})

}

func (th TodoHandler) Deletedata(ctx *gin.Context) {
	_, code, err := th.TodoHan.DeleteService(ctx)

	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Not Found",
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    err,
	})

}

func (th TodoHandler) Updatedata(ctx *gin.Context) {
	var todo domain.TodoUpdate
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":  "Bad Request",
			"message": "Error",
			"data":    nil,
		})
		return
	}

	todo, code, err := th.TodoHan.UpdateService(ctx, todo)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Failed",
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}
