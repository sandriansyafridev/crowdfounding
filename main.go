package main

import (
	"github.com/sandriansyafridev/crowdfounding/app"
	"github.com/sandriansyafridev/crowdfounding/controller"
	"github.com/sandriansyafridev/crowdfounding/repository"
	"github.com/sandriansyafridev/crowdfounding/service"
)

var (
	gormDB, sqlDB = app.NewDB()

	jwtService = service.NewJWTServiceImpl()

	authRepository = repository.NewAuthRepositoryImpl(gormDB)
	authService    = service.NewAuthRepositoryImpl(authRepository)
	authController = controller.NewAuthControllerImpl(authService, jwtService)

	userRepository = repository.NewUserRepositoryImpl(gormDB)
	userService    = service.NewUserServiceImpl(userRepository)
	userController = controller.NewUserControllerImpl(userService)
)

func init() {
	app.InitializeMigration(gormDB)
}

func main() {
	defer sqlDB.Close()

	r := app.NewRoute(userController, authController, userRepository, jwtService)
	r.Run(":8080")
}
