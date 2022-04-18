package repository

import (
	"github.com/EgorBKV4/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID, listID int) (todo.TodoList, error)
	Delete(userID, listID int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
