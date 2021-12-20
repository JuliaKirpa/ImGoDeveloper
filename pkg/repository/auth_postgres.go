package repository

import (
	"ImGoDeveloper"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user ImGoDeveloper.User) (int, error) {
	var id int
	qwery := fmt.Sprintf("INSERT INTO %s (name, usermame, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(qwery, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
