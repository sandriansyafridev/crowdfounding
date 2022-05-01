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

	campaignRepository = repository.NewCampaignRepositoryImpl(gormDB)
	campaignService    = service.NewCampaignServiceImpl(campaignRepository)
	campaignController = controller.NewCampaignControllerImpl(campaignService)
)

func init() {
	app.InitializeMigration(gormDB)
}

func main() {
	defer sqlDB.Close()

	r := app.NewRoute(jwtService, userRepository, authController, userController, campaignController)
	r.Run(":8080")
}
