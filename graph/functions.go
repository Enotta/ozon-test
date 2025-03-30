package graph

import (
	"database/sql"
	"reflect"
)

func validateLocal(list interface{}, id string) bool {
	val := reflect.ValueOf(list)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		idField := elem.FieldByName("ID")
		if !idField.IsValid() || idField.Kind() != reflect.String {
			continue
		}

		if idField.String() == id {
			return true
		}
	}

	return false
}
func secureAuthor(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS authors (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL
		);
	`)

	return err
}
func securePost(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content VARCHAR(2000) NOT NULL,
			author INT NOT NULL REFERENCES authors(id),
			comments_enabled BOOLEAN NOT NULL DEFAULT TRUE,
			created_at TIMESTAMP
		);
	`)

	return err
}
func secureComment(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			content VARCHAR(2000) NOT NULL,
			author INT NOT NULL REFERENCES authors(id),
			post_id INT NOT NULL REFERENCES posts(id),
			parent_id INT,
			created_at TIMESTAMP,
			FOREIGN KEY (parent_id) REFERENCES comments(id)
		);
	`)

	return err
}
