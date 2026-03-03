package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/javiermoralesdev/folio-backend/internal/db"
)

func UpsertBookmark(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			BookID string `json:"book"`
			UserID string `json:"user"`
			Page   int    `json:"page"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		id := uuid.New().String()

		err := queries.UpsertBookmark(context.Background(), db.UpsertBookmarkParams{
			ID:     id,
			UserID: body.UserID,
			BookID: body.BookID,
			Page:   int64(body.Page),
		})

		if err != nil {
			http.Error(w, "Error saving bookmark", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetUserBookmarks(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			UserID string `json:"user"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		bookmarks, err := queries.GetUserBookmarks(context.Background(), body.UserID)

		if err != nil {
			http.Error(w, "Error getting bookmarks", http.StatusInternalServerError)
			return
		}

		if bookmarks == nil {
			bookmarks = []db.Bookmark{}
		}

		json.NewEncoder(w).Encode(bookmarks)
	}
}

func GetBookmark(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			UserID string `json:"user"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		bookID := chi.URLParam(r, "bookId")

		book, err := queries.GetBookmark(context.Background(), db.GetBookmarkParams{
			BookID: bookID,
			UserID: body.UserID,
		})

		if err != nil {
			http.Error(w, "Error getting bookmark", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(book)

	}
}

func DeleteBookmark(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			UserID string `json:"user"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		bookID := chi.URLParam(r, "bookId")

		err := queries.DeleteBookmark(context.Background(), db.DeleteBookmarkParams{
			UserID: body.UserID,
			BookID: bookID,
		})

		if err != nil {
			http.Error(w, "Failed to delete bookmark", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
