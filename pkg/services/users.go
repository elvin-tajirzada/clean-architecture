package services

import (
	"github.com/elvin-tacirzade/clean-architecture/pkg/models"
	"github.com/elvin-tacirzade/clean-architecture/pkg/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type (
	UsersService interface {
		GetAllUsers() ([]models.Users, error)
		FindById(r *http.Request) (models.Users, error)
		InsertUser(r *http.Request, validate *validator.Validate) error
		DeleteUser(r *http.Request) error
	}

	usersService struct {
		UsersRepository repositories.UsersRepository
	}
)

func NewUsersServices(u repositories.UsersRepository) UsersService {
	return &usersService{
		UsersRepository: u,
	}
}

func (u *usersService) GetAllUsers() ([]models.Users, error) {
	users, err := u.UsersRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *usersService) FindById(r *http.Request) (models.Users, error) {
	vars := mux.Vars(r)
	user, err := u.UsersRepository.FindById(vars["id"])
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *usersService) InsertUser(r *http.Request, validate *validator.Validate) error {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	currentTime := time.Now()
	user := models.Users{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	//start validate
	err := validate.Struct(user)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return err
		}
	}
	//end validate
	err = u.UsersRepository.InsertUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u *usersService) DeleteUser(r *http.Request) error {
	vars := mux.Vars(r)
	err := u.UsersRepository.DeleteUser(vars["id"])
	if err != nil {
		return err
	}
	return nil
}
