package repository_todo

import (
	"golangtesting/internal/domain"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (tr TodoRepository) Getdata() ([]domain.Todo, error) {
	var todo []domain.Todo
	err := tr.db.Find(&todo).Error
	return todo, err
}

func (tr TodoRepository) GetdataByIdActivity(id string) ([]domain.Todo, error) {
	var todo []domain.Todo
	err := tr.db.Where("activity_group_id = ?", id).Find(&todo).Error
	return todo, err
}

func (tr TodoRepository) Getdatabyid(id string) (domain.Todo, error) {
	var todo domain.Todo
	err := tr.db.First(&todo, id).Error
	return todo, err
}

func (tr TodoRepository) Createdata(todo domain.Todo) (domain.Todo, error) {

	err := tr.db.Create(&todo).Error
	return todo, err
}

func (tr TodoRepository) GetActivityId(id string) (domain.Activity, error) {
	var activity domain.Activity
	err := tr.db.First(&activity, id).Error

	return activity, err
}

func (tr TodoRepository) Delete(id string) (domain.Todo, error) {
	var todo domain.Todo
	err := tr.db.Delete(&todo, id).Error
	return todo, err
}

func (tr TodoRepository) Update(id string, todo domain.TodoUpdate) (domain.TodoUpdate, error) {
	err := tr.db.Model(&domain.Todo{}).Where("id = ?", id).Updates(domain.Todo{
		Title:     todo.Title,
		IsActive:  todo.IsActive,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}).Error

	return todo, err

}
