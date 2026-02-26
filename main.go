package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/javiermoralesdev/folio-backend/internal/db"

	_ "modernc.org/sqlite"
)

func main() {
	conn, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error opening database!")
	}
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("failed to read schema:", err)
	}

	_, err = conn.Exec(string(schema))
	if err != nil {
		log.Fatal("failed to apply schema:", err)
	}
	queries := db.New(conn)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/users/create", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		_, err := queries.CreateUser(context.Background(), db.CreateUserParams{
			ID:       id, // your uuid
			Username: "john",
			Password: "hashedPassword",
		})
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal("Error creating user")
		}
		w.Write([]byte(id))
	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, err := queries.GetUserByID(context.Background(), id)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(404)
		}

		w.Write([]byte("My name is " + user.Username))

	})
	http.ListenAndServe(":1323", r)
}
