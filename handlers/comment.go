func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postID := r.FormValue("post_id")
		content := r.FormValue("content")

		_, err := models.DB.Exec("INSERT INTO comments (post_id, content) VALUES (?, ?)", postID, content)
		if err != nil {
			http.Error(w, "Error adding comment", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/post?id="+postID, http.StatusSeeOther)
	}
}
