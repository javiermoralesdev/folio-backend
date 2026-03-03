package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/javiermoralesdev/folio-backend/internal/db"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {

	godotenv.Load()
	dbPath := get_database_file()

	//booksPath := get_books_directory()
	conn, err := sql.Open("sqlite", dbPath)
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

	r.Post("/login", Login(queries))
	r.Post("/users/create", CreateUser(queries))

	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/books", GetBooks(queries))
		r.Get("/books/{id}", GetBook(queries))
		r.Get("/books/{id}/file", ServeBook(queries))
		r.Post("/books/upload", UploadBook(queries))
		r.Delete("/books/{id}", DeleteBook(queries)) // TODO

		r.Post("/bookmarks", UpsertBookmark(queries))            //TODO
		r.Get("/bookmarks", GetUserBookmarks(queries))           //TODO
		r.Get("/bookmarks/{bookId}", GetBookmark(queries))       //TODO
		r.Delete("/bookmarks/{bookId}", DeleteBookmark(queries)) //TODO
	})

	http.ListenAndServe(":1323", r)
}
