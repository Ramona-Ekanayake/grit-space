package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ramonaekanayake/grit-space/models"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var like models.Like
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = models.DB.Exec("INSERT INTO likes (post_id, user_id) VALUES (?, ?)", like.PostID, like.UserID)
	if err != nil {
		http.Error(w, "Error liking post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dislike models.Dislike
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = models.DB.Exec("INSERT INTO dislikes (post_id, user_id) VALUES (?, ?)", dislike.PostID, dislike.UserID)
	if err != nil {
		http.Error(w, "Error disliking post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
