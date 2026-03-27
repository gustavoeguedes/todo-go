package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "password123")
	assert.NoError(t, err)

	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
}
