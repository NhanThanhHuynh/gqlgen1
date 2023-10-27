package domain

import "github.com/NhanThanhHuynh/gqlgen1/graph/postgres"

type Domain struct {
	UsersRepo postgres.UsersRepo
	TodoRepo  postgres.TodosRepo
}

func NewDomain(usersRepo postgres.UsersRepo, todosRepo postgres.TodosRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, TodoRepo: todosRepo}
}
