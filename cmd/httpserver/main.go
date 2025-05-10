package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"user-crud-api/internal/handler"
	"user-crud-api/internal/repository"
	"user-crud-api/internal/service"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users", userHandler.ListUsers)
	r.Get("/users/{id}", userHandler.GetUser)
	r.Put("/users/{id}", userHandler.UpdateUser)
	r.Delete("/users/{id}", userHandler.DeleteUser)

	http.ListenAndServe(":8080", r)
}
