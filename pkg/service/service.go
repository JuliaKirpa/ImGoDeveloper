package service

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/repository"
)

type Authorization interface {
	CreateUser(user ImGoDeveloper.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list ImGoDeveloper.TodoList) (int, error)
	GetAll(userId int) ([]ImGoDeveloper.TodoList, error)
	GetById(userId, listId int) (ImGoDeveloper.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
