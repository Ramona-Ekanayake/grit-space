package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ramonaekanayake/grit-space/config"
	"github.com/ramonaekanayake/grit-space/handlers"
	"github.com/ramonaekanayake/grit-space/middleware"
)

func main() {
	config.InitDB()

	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.Handle("/api/posts", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreatePostHandler)))
	http.Handle("/api/comments", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateCommentHandler)))
	http.Handle("/api/posts/like", middleware.AuthMiddleware(http.HandlerFunc(handlers.LikePostHandler)))
	http.Handle("/api/posts/dislike", middleware.AuthMiddleware(http.HandlerFunc(handlers.DislikePostHandler)))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	// Serve HTML files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := "./web/templates" + r.URL.Path
		if path == "./web/templates/" {
			path = "./web/templates/index.html"
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, "./web/templates/404.html")
			return
		}
		http.ServeFile(w, r, path)
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
