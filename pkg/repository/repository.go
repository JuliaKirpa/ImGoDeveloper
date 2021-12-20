package repository

import (
	"ImGoDeveloper"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ImGoDeveloper.User) (int, error)
	GetUser(username, password string) (ImGoDeveloper.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
