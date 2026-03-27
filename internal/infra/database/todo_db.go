package database

import (
	"github.com/gustavoeguedes/todo-go/internal/entity"
	"gorm.io/gorm"
)

type Todo struct {
	DB *gorm.DB
}

func NewTodo(db *gorm.DB) *Todo {
	return &Todo{DB: db}
}

func (t Todo) FindByID(id, userId string) (*entity.Todo, error) {
	var todo entity.Todo
	err := t.DB.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t Todo) Update(todoId, userId string) error {
	todoExists, err := t.FindByID(todoId, userId)
	if err != nil {
		return err
	}
	todoExists.MarkDone()
	return t.DB.Save(todoExists).Error
}

func (t Todo) Delete(todoId, userId string) error {
	todoExists, err := t.FindByID(todoId, userId)
	if err != nil {
		return err
	}

	return t.DB.Delete(todoExists).Error
}

func (t Todo) Create(todo *entity.Todo) error {
	return t.DB.Create(todo).Error
}

func (t Todo) FindAll(page, limit int, userId string) ([]entity.Todo, error) {
	var todos []entity.Todo
	var err error

	if page != 0 && limit != 0 {
		err = t.DB.Limit(limit).Offset((page-1)*limit).Where("user_id = ?", userId).Find(&todos).Error
	} else {
		err = t.DB.Find(&todos).Where("user_id = ?", userId).Error
	}

	return todos, err
}
