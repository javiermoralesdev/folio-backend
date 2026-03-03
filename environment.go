package main

import "os"

func get_database_file() string {
	x := os.Getenv("FOLIO_DB_PATH")
	return x
}

func get_jwt_secret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func get_books_directory() string {
	x := os.Getenv("FOLIO_BOOKS_PATH")
	return x
}
