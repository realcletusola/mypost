package routes 

import (
	"github.com/go-chi/chi/v5"
	"github.com/realcletusola/mypost/internal/handlers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/posts", handlers.CreatePost)
	r.Get("/posts", handlers.GetPosts)
	r.Get("/posts/{id}", handlers.GetPostByID)
	return r 
}