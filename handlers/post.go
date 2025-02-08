package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ramonaekanayake/grit-space/models"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if post.Title == "" || post.Content == "" || post.Category == "" {
		http.Error(w, "Title, content, and category are required", http.StatusBadRequest)
		return
	}

	_, err = models.DB.Exec("INSERT INTO posts (title, content, category) VALUES (?, ?, ?)", post.Title, post.Content, post.Category)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
