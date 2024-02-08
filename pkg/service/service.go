package service

import "github.com/Blxssy/todo-app/pkg/repository"

type Authorization interface {
}

type TodoList interface {
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
		Authorization: nil,
		TodoList:      nil,
		TodoItem:      nil,
	}
}
