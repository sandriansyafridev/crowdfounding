package main

import (
	"github.com/sandriansyafridev/crowdfounding/app"
	"github.com/sandriansyafridev/crowdfounding/controller"
	"github.com/sandriansyafridev/crowdfounding/repository"
	"github.com/sandriansyafridev/crowdfounding/service"
)

var (
	gormDB, sqlDB  = app.NewDB()
	userRepository = repository.NewUserRepositoryImpl(gormDB)
	userService    = service.NewUserServiceImpl(userRepository)
	userController = controller.NewUserControllerImpl(userService)
)

func main() {
	app.InitializeMigration(gormDB)
	defer sqlDB.Close()

	r := app.NewRoute(userController)
	r.Run(":8080")

}
