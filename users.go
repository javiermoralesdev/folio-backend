package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/javiermoralesdev/folio-backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// strip "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return get_jwt_secret(), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CreateUser(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// parse the JSON body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// now use body.Username and body.Password
		hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		fmt.Println(body.Username, string(hashed))
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		_, err = queries.CreateUser(context.Background(), db.CreateUserParams{
			ID:       id, // your uuid
			Username: body.Username,
			Password: string(hashed),
		})
		if err != nil {
			http.Error(w, "user already exists", http.StatusBadRequest)
		} else {
			w.Write([]byte(id))
		}
	}
}

func Login(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
