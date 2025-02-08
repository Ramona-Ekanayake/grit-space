package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ramonaekanayake/grit-space/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	_, err = models.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", req.Username, req.Email, hashedPassword)
	if err != nil {
		http.Error(w, "User already exists!", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var userID int
	var dbPassword string
	err = models.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", req.Email).Scan(&userID, &dbPassword)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	_, err = models.DB.Exec("INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)", userID, sessionToken, expiresAt)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	w.WriteHeader(http.StatusOK)
}
