package graph

import (
	"database/sql"
	"ozon-test/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Storage byte

const (
	InMemory Storage = 0
	Postgres Storage = 1
)

type Resolver struct {
	Storage    Storage
	Connection *sql.DB
	authors    []*model.Author
	comments   []*model.Comment
	posts      []*model.Post
}
