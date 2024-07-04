package services_test

import (
	"apibe23/internal/features/users"
	"apibe23/internal/features/users/services"
	"apibe23/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	qry := mocks.NewQuery(t)
	pu := mocks.NewPasswordUtilityInterface(t)
	srv := services.NewUserService(qry, pu)
	input := users.User{Name: "Jerry", Password: "altaalta", Email: "jerry@alterra.id", Phone: "12345"}

	t.Run("Success Register", func(t *testing.T) {
		inputQry := users.User{Name: "Jerry", Password: "somepassword", Email: "jerry@alterra.id", Phone: "12345"}

		pu.On("GeneratePassword", input.Password).Return([]byte("somepassword"), nil).Once()
		qry.On("Register", inputQry).Return(nil).Once()

		err := srv.Register(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error Hash Password", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()

		err := srv.Register(input)

		pu.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "input data tidak valid, data tidak bisa diproses")
	})

	t.Run("Error From Query", func(t *testing.T) {
		inputQry := users.User{Name: "Jerry", Password: "goodpassword", Email: "jerry@alterra.id", Phone: "12345"}
		pu.On("GeneratePassword", input.Password).Return([]byte("goodpassword"), nil).Once()
		qry.On("Register", inputQry).Return(gorm.ErrInvalidData).Once()

		err := srv.Register(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengolah data")

	})
}

func TestLogin(t *testing.T) {

}
