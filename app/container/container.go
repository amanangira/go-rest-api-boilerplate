package container

import (
	"boilerplate/app"
	"boilerplate/controller"
	"boilerplate/service"
	"boilerplate/store/repository"
)

type Container struct {
	UserController *controller.UserController
}

func NewContainer() *Container {
	// TODO - define logic to create db connection under app
	dbClient := app.NewSqlxClient()

	// Repository
	userRepository := repository.NewUserRepository(dbClient)

	// Services
	userService := service.NewUserService(userRepository)

	// controllers
	userController := controller.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
