package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/gustavoeguedes/todo-go/configs"
	"github.com/gustavoeguedes/todo-go/internal/entity"
	"github.com/gustavoeguedes/todo-go/internal/infra/database"
	"github.com/gustavoeguedes/todo-go/internal/infra/webserver/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(configs.DBUrl))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Todo{})
	todoDB := database.NewTodo(db)
	userDB := database.NewUser(db)

	userHandler := handlers.NewUserHandler(userDB)
	todoHandler := handlers.NewTodoHandler(todoDB)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JWTExpiresIn))

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)

		r.Route("/todos", func(r chi.Router) {
			r.Use(jwtauth.Verifier(configs.TokenAuth))
			r.Use(jwtauth.Authenticator)

			r.Get("/", todoHandler.FindAll)
			r.Get("/{id}", todoHandler.FindByID)
			r.Post("/", todoHandler.Create)
			r.Put("/{id}", todoHandler.Update)
			r.Delete("/{id}", todoHandler.Delete)
		})
	})

	if err := http.ListenAndServe(configs.WebServerPort, r); err != nil {
		log.Fatal(err)
	}

}
