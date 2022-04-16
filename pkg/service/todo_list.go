package service

import (
	"github.com/EgorBKV4/todo-app"
	"github.com/EgorBKV4/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}
