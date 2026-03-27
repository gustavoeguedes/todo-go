package database

import "github.com/gustavoeguedes/todo-go/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}

type TodoInterface interface {
	Create(todo *entity.Todo) error
	FindByID(id, userId string) (*entity.Todo, error)
	FindAll(page, limit int, userId string) ([]entity.Todo, error)
	Update(todoId, userId string) error
	Delete(todoId, userId string) error
}
