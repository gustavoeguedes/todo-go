package entity

import "errors"

var (
	ErrInvalidTitle = errors.New("invalid title")
)

type Todo struct {
	ID     ID     `json:"id" db:"id"`
	UserID ID     `json:"userId" db:"user_id" gorm:"column:user_id;not null;index"`
	Title  string `json:"title" db:"title"`
	Done   bool   `json:"done" db:"done"`
}

func NewTodo(userID ID, title string) (*Todo, error) {
	t := &Todo{
		ID:     NewID(),
		UserID: userID,
		Title:  title,
		Done:   false,
	}

	err := t.Validate()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Todo) MarkDone() {
	t.Done = true
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return ErrInvalidTitle
	}
	return nil
}
