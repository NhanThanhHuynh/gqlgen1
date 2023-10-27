package postgres

import (
	"fmt"
	"github.com/NhanThanhHuynh/gqlgen1/graph/model"
	"github.com/go-pg/pg/v10"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	fmt.Println(field, value)
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	if err != nil {
		fmt.Println(err, "err")
	}
	fmt.Println(&user, "user")
	return &user, err
}

func (u *UsersRepo) GetUserByID(id string) (*model.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UsersRepo) CreateUser(user *model.User) (*model.User, error) {
	fmt.Println(u.DB, "Hello world")

	_, err := u.DB.Model(user).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
