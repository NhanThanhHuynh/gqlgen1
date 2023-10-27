package postgres

import (
	"fmt"
	"github.com/NhanThanhHuynh/gqlgen1/graph/model"
	"github.com/go-pg/pg/v10"
)

type TodosRepo struct {
	DB *pg.DB
}

func (u *TodosRepo) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	_, err := u.DB.Model(todo).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(todo, "Todo")
	return todo, nil
}

func (u *TodosRepo) GetTodosByField(field, value string) (*model.Todo, error) {
	var todo model.Todo
	err := u.DB.Model(&todo).Where(fmt.Sprintf("%v = ?", field), value).First()
	if err != nil {
		fmt.Println(err, "err")
	}
	fmt.Println(&todo, "user")
	return &todo, err
}

func (u *TodosRepo) GetTodosById(id string) (*model.Todo, error) {
	return u.GetTodosByField("id", id)
}

func (u *TodosRepo) UpdateTodo(field, value string, input *model.Todo) (*model.Todo, error) {
	_, err := u.DB.Model(input).Where(fmt.Sprintf("%v = ?", field), value).Update()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input, &input, "input")
	return input, err
}

func (u *TodosRepo) DeleteTodo(field, value string) (bool, error) {
	var todo model.Todo
	fmt.Println(field, value, "field - value")
	_, err := u.DB.Model(&todo).Where(fmt.Sprintf("%v = ?", field), value).Delete()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(todo, &todo, "todo")
	return true, err
}

func (u *TodosRepo) GetTodoDetailsByField(input *model.Todo) ([]*model.TodoDetail, error) {
	var todoDetail []*model.TodoDetail
	fmt.Println("GetTodoDetailsByField ----------------------------------------", input.UserID)
	err := u.DB.Model(&todoDetail).Where("user_id = ?", input.UserID).Select()
	if err != nil {
		fmt.Println(err, "err")
	}
	fmt.Println(todoDetail, "user")
	return todoDetail, err
}
