package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	userId := NewID()
	todo, err := NewTodo(userId, "Test Todo")
	assert.NoError(t, err)
	assert.Equal(t, "Test Todo", todo.Title)
	assert.Equal(t, userId, todo.UserID)
	assert.False(t, todo.Done)
}

func TestTodo_Validate(t *testing.T) {
	userId := NewID()
	_, err := NewTodo(userId, "")
	assert.ErrorIs(t, err, ErrInvalidTitle)
}
