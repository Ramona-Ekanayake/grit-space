package main

import (
	"fmt"
	"grit-space/handlers"
	"grit-space/models"
	"net/http"
)

func main() {
	models.InitDB()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/create-post", handlers.CreatePostHandler)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
