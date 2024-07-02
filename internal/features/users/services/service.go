package services

import (
	"apibe23/internal/features/users"
	"apibe23/internal/utils"
)

type UserServices struct {
	qry users.Query
}

func NewUserService(q users.Query) users.Services {
	return &UserServices{
		qry: q,
	}
}

func (us *UserServices) Register(newData users.User) error {
	processPw, err := utils.GeneratePassword(newData.Password)

	if err != nil {
		return err
	}

	newData.Password = string(processPw)

	err = us.qry.Register(newData)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserServices) Login(email string, password string) (users.User, string, error) {
	result, err := us.qry.Login(email)
	if err != nil {
		return users.User{}, "", err
	}

	us.kerjaLogin()

	err = utils.CheckPassword([]byte(password), []byte(result.Password))

	if err != nil {
		return users.User{}, "", err
	}

	token, err := utils.GenerateToken(result.ID)
	if err != nil {
		return users.User{}, "", err
	}

	return result, token, nil
}

func (us *UserServices) kerjaLogin() {}
