package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/javiermoralesdev/folio-backend/internal/db"
)

func UploadBook(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// limit upload size to 100MB
		r.ParseMultipartForm(100 << 20)

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "invalid file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// save to disk
		file_to_save := get_books_directory() + "/" + header.Filename
		dst, err := os.Create(file_to_save)
		if err != nil {
			http.Error(w, "failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		io.Copy(dst, file)

		id := uuid.New().String()

		// parse the JSON body
		title := r.FormValue("title")

		author := r.FormValue("author")
		if title == "" || author == "" {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		_, err = queries.CreateBook(context.Background(), db.CreateBookParams{
			ID:     id,
			Title:  title,
			Author: author,
			Path:   file_to_save,
		})

		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				http.Error(w, "book already exists", http.StatusConflict)
				return
			}
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("The book " + title + " escrito por " + author + " ha sido agregado con el id " + id))
	}
}

func ServeBook(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		book, err := queries.GetBookByID(context.Background(), id)
		if err != nil {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}

		http.ServeFile(w, r, book.Path)

	}
}

func GetBook(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetBooks(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteBook(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
