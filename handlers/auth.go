package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"myforum/database"
	"myforum/utils"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Inscription
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err = database.DB.Exec(query, user.Username, user.Email, hashedPassword)

	if err != nil {
		http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inscription réussie"})
}

// Connexion
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	var storedUser User
	err = database.DB.QueryRow(`SELECT id, password FROM users WHERE email = ?`, user.Email).Scan(&storedUser.ID, &storedUser.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	sessionID := utils.CreateSession(storedUser.ID)
	http.SetCookie(w, &http.Cookie{Name: "session", Value: sessionID, HttpOnly: true})

	json.NewEncoder(w).Encode(map[string]string{"message": "Connexion réussie"})
}