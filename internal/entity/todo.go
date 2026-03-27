package entity

import "errors"

var (
	ErrInvalidTitle = errors.New("invalid title")
)

type Todo struct {
	ID     ID     `json:"id"`
	UserID ID     `json:"user_id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
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
