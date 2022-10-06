package service_todo

import (
	"errors"
	"fmt"
	"golangtesting/internal/domain"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoRepo interface {
	Getdata() ([]domain.Todo, error)
	GetdataByIdActivity(id string) ([]domain.Todo, error)
	Getdatabyid(id string) (domain.Todo, error)
	Createdata(todo domain.Todo) (domain.Todo, error)
	GetActivityId(id string) (domain.Activity, error)
	Delete(id string) (domain.Todo, error)
	Update(id string, todo domain.TodoUpdate) (domain.TodoUpdate, error)
}

type TodoService struct {
	TodoSer TodoRepo
}

func NewTodoService(tr TodoRepo) *TodoService {
	return &TodoService{tr}
}

func (ts TodoService) GetDataService(ctx *gin.Context) ([]domain.Todo, int, error) {
	query := ctx.Request.URL.Query().Get("activity_group_id")
	if query == "" {
		todo, err := ts.TodoSer.Getdata()
		if err != nil {
			return nil, 500, errors.New("error get data from database")
		}
		return todo, 200, nil
	} else {
		todo, err := ts.TodoSer.GetdataByIdActivity(query)
		if err != nil {
			return nil, 200, errors.New("Success")
		}
		return todo, 200, nil
	}

}

func (ts TodoService) GetdataByidService(ctx *gin.Context) (domain.Todo, int, error) {
	id := ctx.Param("id")

	todo, err := ts.TodoSer.Getdatabyid(id)
	if err != nil {
		msg := fmt.Sprintf("Todo with ID %s Not Found", id)
		return todo, 404, errors.New(msg)
	}
	return todo, 200, nil
}

func (ts TodoService) CreateService(todo domain.Todo, ctx *gin.Context) (domain.Todo, int, error) {
	id := todo.IdActivity
	_, err := ts.TodoSer.GetActivityId(id)
	if err != nil {
		msg := fmt.Sprintf("Activity with activity_group_id %s Not Found", id)
		return domain.Todo{}, 404, errors.New(msg)
	} else {
		todo, err := ts.TodoSer.Createdata(todo)
		if err != nil {
			return todo, 500, errors.New("error create data todo")
		}
		return todo, 200, nil
	}

}

func (ts TodoService) DeleteService(ctx *gin.Context) (domain.Todo, int, error) {
	id := ctx.Param("id")
	todo, err := ts.TodoSer.Getdatabyid(id)
	if err != nil {
		msg := fmt.Sprintf("Todo with ID %s Not Found", id)
		return todo, 404, errors.New(msg)
	} else {
		todo, err := ts.TodoSer.Delete(id)
		if err != nil {
			return todo, 500, errors.New("errors delete data")
		}
		return todo, 200, nil
	}

}
func (ts TodoService) UpdateService(ctx *gin.Context, todoupdate domain.TodoUpdate) (domain.TodoUpdate, int, error) {
	id := ctx.Param("id")
	_, err := ts.TodoSer.Getdatabyid(id)
	if err != nil {
		msg := fmt.Sprintf("Todo with ID %s Not Found", id)
		return domain.TodoUpdate{}, 404, errors.New(msg)
	} else {
		todo, _ := ts.TodoSer.Update(id, todoupdate)
		id_primary, _ := strconv.Atoi(id)
		todo.ID = id_primary
		todo.CreatedAt = time.Now()
		todo.UpdatedAt = time.Now()
		return todo, 200, nil
	}
}
