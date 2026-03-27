package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/gustavoeguedes/todo-go/internal/dto"
	"github.com/gustavoeguedes/todo-go/internal/infra/database"
)

type TodoHandler struct {
	TodoDB database.TodoInterface
}

func NewTodoHandler(todoDB database.TodoInterface) *TodoHandler {
	return &TodoHandler{
		TodoDB: todoDB,
	}
}

func (t *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo dto.CreateTodoInput
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sub := claims["sub"].(string)

	todoEntity, err := todo.ToEntity(sub)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = t.TodoDB.Create(todoEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (t *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sub := claims["sub"].(string)
	id := chi.URLParam(r, "id")

	err = t.TodoDB.Update(id, sub)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (t *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sub := claims["sub"].(string)
	id := chi.URLParam(r, "id")

	err = t.TodoDB.Delete(id, sub)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *TodoHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}

	sub := claims["sub"].(string)

	log.Printf("page %d limit %d", page, limit)
	todos, err := t.TodoDB.FindAll(page, limit, sub)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (t *TodoHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sub := claims["sub"].(string)
	id := chi.URLParam(r, "id")

	todo, err := t.TodoDB.FindByID(id, sub)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
