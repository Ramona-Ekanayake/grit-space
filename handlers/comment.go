package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ramonaekanayake/grit-space/models"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if comment.PostID == 0 || comment.Content == "" {
		http.Error(w, "Post ID and content are required", http.StatusBadRequest)
		return
	}

	_, err = models.DB.Exec("INSERT INTO comments (post_id, content) VALUES (?, ?)", comment.PostID, comment.Content)
	if err != nil {
		http.Error(w, "Error adding comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
