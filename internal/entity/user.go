package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       ID     `json:"id" db:"id" gorm:"primaryKey"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
