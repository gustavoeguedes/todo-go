package dto

import "github.com/gustavoeguedes/todo-go/internal/entity"

type CreateTodoInput struct {
	Title string `json:"title"`
}

func (t *CreateTodoInput) ToEntity(sub string) (*entity.Todo, error) {

	userId, err := entity.ParseID(sub)
	if err != nil {
		return nil, err
	}
	todo, err := entity.NewTodo(userId, t.Title)
	if err != nil {
		return nil, err
	}

	return todo, nil

}
