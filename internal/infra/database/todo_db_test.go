package database

import (
	"strconv"
	"testing"

	"github.com/gustavoeguedes/todo-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTodo_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Todo{})
	userId := entity.NewID()
	todo, _ := entity.NewTodo(userId, "Todo 1")
	todoDb := NewTodo(db)
	err = todoDb.Create(todo)
	if err != nil {
		t.Error(err)
	}

	var todoFound entity.Todo
	err = db.First(&todoFound, todo.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, todo.ID, todoFound.ID)
	assert.Equal(t, todo.UserID, todoFound.UserID)
	assert.Equal(t, todo.Title, todoFound.Title)
	assert.Equal(t, todo.Done, todoFound.Done)

}

func TestTodo_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Todo{})
	userId := entity.NewID()
	todo, _ := entity.NewTodo(userId, "Todo 1")
	todoDb := NewTodo(db)
	err = todoDb.Create(todo)
	if err != nil {
		t.Error(err)
	}

	err = todoDb.Update(todo.ID.String(), userId.String())
	assert.NoError(t, err)
}

func TestTodo_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Todo{})
	userId := entity.NewID()
	todo, _ := entity.NewTodo(userId, "Todo 1")
	todoDb := NewTodo(db)
	err = todoDb.Create(todo)
	if err != nil {
		t.Error(err)
	}

	err = todoDb.Delete(todo.ID.String(), userId.String())
	assert.NoError(t, err)

	var todoFound entity.Todo
	err = db.First(&todoFound, todo.ID).Error
	assert.Error(t, err)

	err = todoDb.Create(todo)
	assert.NoError(t, err)

	todo.UserID = entity.NewID()

	err = todoDb.Delete(todo.ID.String(), todo.UserID.String())
	assert.Error(t, err)
}

func TestTodo_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Todo{})

	todoDb := NewTodo(db)
	userId := entity.NewID()
	for i := 0; i <= 25; i++ {
		todo, _ := entity.NewTodo(userId, "Todo "+strconv.FormatInt(int64(i), 10))
		err = todoDb.Create(todo)
		if err != nil {
			t.Error(err)
		}
	}

	todos, err := todoDb.FindAll(2, 10, userId.String())
	assert.NoError(t, err)
	assert.Len(t, todos, 10)
}

func TestTodo_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Todo{})
	userId := entity.NewID()
	todo, _ := entity.NewTodo(userId, "Todo 1")
	todoDb := NewTodo(db)
	err = todoDb.Create(todo)
	if err != nil {
		t.Error(err)
	}

	todoFound, err := todoDb.FindByID(todo.ID.String(), userId.String())
	assert.NoError(t, err)
	assert.Equal(t, todo.ID, todoFound.ID)
	assert.Equal(t, todo.UserID, todoFound.UserID)
	assert.Equal(t, todo.Title, todoFound.Title)
	assert.Equal(t, todo.Done, todoFound.Done)
}
