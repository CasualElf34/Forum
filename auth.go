package auth

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"myforum/database"
)

// Inscription d'un utilisateur
func RegisterUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err = database.DB.Exec(query, username, email, string(hashedPassword))
	if err != nil {
		return err
	}

	log.Println("Utilisateur enregistré :", username)
	return nil
}

// Connexion d'un utilisateur
func LoginUser(email, password string) (string, error) {
	var storedPassword string
	var userID int

	query := `SELECT id, password FROM users WHERE email = ?`
	err := database.DB.QueryRow(query, email).Scan(&userID, &storedPassword)
	if err == sql.ErrNoRows {
		return "", errors.New("Utilisateur non trouvé")
	} else if err != nil {
		return "", err
	}

	// Vérification du mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return "", errors.New("Mot de passe incorrect")
	}

	// Génération d'un UUID pour la session
	sessionID := uuid.New().String()
	log.Println("Utilisateur connecté avec session :", sessionID)
	return sessionID, nil
}