package controllers

import (
	"github.com/elvin-tacirzade/clean-architecture/pkg/helpers"
	"github.com/elvin-tacirzade/clean-architecture/pkg/services"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type (
	UsersController interface {
		GetAllUsers(w http.ResponseWriter, r *http.Request)
		FindById(w http.ResponseWriter, r *http.Request)
		InsertUser(w http.ResponseWriter, r *http.Request)
		DeleteUser(w http.ResponseWriter, r *http.Request)
	}

	usersController struct {
		UsersService services.UsersService
	}
)

var validate = validator.New()

func NewUsersController(u services.UsersService) UsersController {
	return &usersController{
		UsersService: u,
	}
}

func (u *usersController) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := u.UsersService.GetAllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	helpers.JsonNewEncoder(w, http.StatusOK, users)
}

func (u *usersController) FindById(w http.ResponseWriter, r *http.Request) {
	user, err := u.UsersService.FindById(r)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	helpers.JsonNewEncoder(w, http.StatusOK, user)
}

func (u *usersController) InsertUser(w http.ResponseWriter, r *http.Request) {
	response := u.UsersService.InsertUser(r, validate)
	if response.Error != nil {
		log.Println(response.Error)
		helpers.JsonNewEncoder(w, response.StatusCode, map[string]string{
			"status":  "error",
			"message": response.Error.Error(),
		})
		return
	}
	helpers.JsonNewEncoder(w, response.StatusCode, map[string]string{
		"status":  "ok",
		"message": "User successfully added",
	})
}

func (u *usersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := u.UsersService.DeleteUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	helpers.JsonNewEncoder(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "User successfully deleted",
	})
}
