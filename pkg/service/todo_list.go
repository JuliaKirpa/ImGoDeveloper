package service

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list ImGoDeveloper.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
