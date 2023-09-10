package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sk25469/push_noti_service/pkg/config"
	"github.com/sk25469/push_noti_service/pkg/middleware"
	"github.com/sk25469/push_noti_service/pkg/model"
	"github.com/sk25469/push_noti_service/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var user model.UserModel

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error decoding json: [%v]", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	// Hash the user's password (you should use a secure password hashing library)
	hashedPassword, err := middleware.HashPassword(user.Password)
	if err != nil {
		log.Printf("error hashing password: [%v]", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error hashing password"))
		return
	}

	db := config.GetPostgresConnection()

	// Insert the user's information into the database
	_, err = db.Exec(utils.InsertSQL("email", "password"), user.Email, hashedPassword)
	if err != nil {
		log.Printf("error executing sql insert query: [%v]", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating user"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	var user model.UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error decoding json: [%v]", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	db := config.GetPostgresConnection()

	// Query the database to retrieve the user's hashed password
	// TODO: After logging in, update the last login timestamp to now
	var hashedPassword string
	err := db.QueryRow(utils.SelectSQL("password", "email"), user.Email).Scan(&hashedPassword)
	if err != nil {
		log.Printf("error executing sql select query: [%v]", err)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authentication failed"))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error querying database"))
			return
		}
	}

	// Verify the provided password against the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		log.Printf("error comparing hash and password: [%v]", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authentication failed"))
		return
	}

	// If authentication is successful, create a JWT token
	token, err := middleware.CreateJWTToken(user.Email)
	if err != nil {
		log.Printf("error creating JWT Token: [%v]", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating JWT token"))
		return
	}

	// Return the JWT token as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token": "` + token + `"}`))
}
