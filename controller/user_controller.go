package controller

import (
	"boilerplate/app"
	"boilerplate/app/app_error"
	"boilerplate/service"
	"boilerplate/store/dtos"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// TODO - add validation for incoming requests

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u UserController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input dtos.UserDTO

	// TODO - decode to view equivalent and run validations on it
	decodeErr := json.NewDecoder(r.Body).Decode(&input)
	if decodeErr != nil {
		app_error.NewError(decodeErr, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	err := u.userService.Create(ctx, &input)
	if err != nil {
		app_error.NewError(err, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	app.WriteJSON(w, http.StatusOK, struct {
		ID string
	}{ID: input.ID})
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := chi.URLParam(r, "user_id")
	result, err := u.userService.Get(ctx, userID)
	if err != nil {
		app_error.NewError(err, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	app.WriteJSON(w, http.StatusOK, result)
}
