package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/controller"
	"github.com/sandriansyafridev/crowdfounding/middleware"
	"github.com/sandriansyafridev/crowdfounding/repository"
	"github.com/sandriansyafridev/crowdfounding/service"
)

func NewRoute(
	jwtService service.JWTService,
	userRepository repository.UserRepository,
	authController controller.AuthController,
	userController controller.UserController,
	campaignController controller.CampaignController,
) *gin.Engine {

	r := gin.New()

	var rootStaticPath = "public/assets/images/"
	staticDirImage := r.Group(rootStaticPath)
	staticDirImageCampaign := staticDirImage.Group("campaigns")
	staticDirImageUser := staticDirImage.Group("users")

	staticDirImageCampaign.Static("/", fmt.Sprintf("./%s/campaigns/", rootStaticPath))
	staticDirImageUser.Static("/", fmt.Sprintf("./%s/users/", rootStaticPath))

	v1 := r.Group("/api/v1/")

	authRoutes_v1 := v1.Group("auth")
	authRoutes_v1.POST("/login", authController.Login)
	authRoutes_v1.POST("/register", authController.Register)

	userRoutes_v1 := v1.Group("users")
	userRoutes_v1.GET("/", userController.GetUsers)
	userRoutes_v1.GET("/:id", userController.GetUserByID)
	userRoutes_v1.DELETE("/:id", userController.Delete)
	userRoutes_v1.POST("/profile-image", userController.UploadProfileImage)

	campaignRoutes_v1 := v1.Group("campaigns")
	campaignRoutes_v1.Use(middleware.AuthorizationMiddleware(userRepository, jwtService))
	campaignRoutes_v1.GET("/", campaignController.GetCampaigns)
	campaignRoutes_v1.POST("/", campaignController.CreateCampaign)
	campaignRoutes_v1.GET("/:id", campaignController.GetCampaign)
	campaignRoutes_v1.PUT("/:id", campaignController.UpdateCampaign)

	return r
}
