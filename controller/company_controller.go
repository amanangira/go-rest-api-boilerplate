package controller

import (
	"boilerplate/app"
	"boilerplate/app/app_error"
	"boilerplate/persistence/dtos"
	"boilerplate/service"
	"encoding/json"
	"net/http"
)

// TODO - add validation for incoming requests

type CompanyController struct {
	companyService service.ICompanyService
}

func NewCompanyController(
	companyService service.ICompanyService) *CompanyController {
	return &CompanyController{
		companyService: companyService,
	}
}

func (p CompanyController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input dtos.CompanyRequestDTO

	// TODO - decode to view equivalent and run validations on it
	decodeErr := json.NewDecoder(r.Body).Decode(&input)
	if decodeErr != nil {
		app_error.NewError(decodeErr, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	err := p.companyService.Create(ctx, &input)
	if err != nil {
		app_error.NewError(err, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	app.WriteJSON(w, http.StatusOK, nil)
}
