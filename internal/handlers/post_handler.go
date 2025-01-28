package handlers

import (
	// "encoding/json"
	"net/http"
	"github.com/realcletusola/mypost/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"sync"
	"time"
)

var (
	// simulate in-memory database 
	posts 	   = make(map[string]*models.Post)
	categories = make(map[string]*models.Category)
	mu		   sync.Mutex 
)

// Create function to create post 
func CreatePost(w http.ResponseWriter, r *http.Request) {
	go func() {
		var post models.Post
		if err := decodeJSONBody(w, r, &post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}
		// make sure only on go routine can access the resourece at a time to avoid race condition
		mu.Lock()
		defer mu.Unlock() // unlock the mutex to allow other goroutine access after the current one is done

		post.ID = uuid.New().String()
		if post.Category != nil {
			post.Category.ID = uuid.New().String()
			categories[post.Category.ID] = post.Category // store in map (database shared memory)
		}
		post.Date = time.Now()
		posts[post.ID] = &post // store to map as a pointer so any update will be reflected

		w.WriteHeader(http.StatusCreated)
	}()
}


// Create function to get all posts 
func GetPosts(w http.ResponseWriter, r *http.Request) {
	go func() {
		mu.Lock()
		defer mu.Unlock()

		// create slice to of 0 initial length and capacity of post length 
		postList := make([]*models.Post, 0, len(posts))
		for _, post := range posts {
			postList = append(postList, post) //append slice with post
		}
		encodeJSONBody(w, postList)
	}()
}


// Create function to get each post by ID 
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	go func() {
		mu.Lock()
		defer mu.Unlock()
		post, exists := post[id]
		if !exists {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		encodeJSONBody(w, post)
	}()
}