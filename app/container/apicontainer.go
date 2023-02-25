package container

import (
	api "boilerplate"
	"boilerplate/controller"
	"boilerplate/repository"
	"boilerplate/service"
)

type API struct {
	CompanyController *controller.CompanyController
}

var APIContainer API

func InitControllers(api api.IAPI) {
	// Repository
	companyRepository := repository.NewCompanyRepository(api.GetDBClient())

	// Services
	companyService := service.NewCompanyService(companyRepository)

	// controllers
	userController := controller.NewCompanyController(companyService)

	APIContainer = API{
		CompanyController: userController,
	}
}
