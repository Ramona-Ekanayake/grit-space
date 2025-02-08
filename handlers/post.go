func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := models.DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
		if err != nil {
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
