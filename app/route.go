package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/controller"
	"github.com/sandriansyafridev/crowdfounding/repository"
	"github.com/sandriansyafridev/crowdfounding/service"
)

func NewRoute(
	userController controller.UserController,
	authController controller.AuthController,
	userRepository repository.UserRepository,
	jwtService service.JWTService,
) *gin.Engine {

	r := gin.New()
	v1 := r.Group("/api/v1/")

	authRoutes_v1 := v1.Group("auth")
	authRoutes_v1.POST("/login", authController.Login)
	authRoutes_v1.POST("/register", authController.Register)

	userRoutes_v1 := v1.Group("users")
	userRoutes_v1.GET("/", userController.GetUsers)
	userRoutes_v1.GET("/:id", userController.GetUserByID)
	userRoutes_v1.DELETE("/:id", userController.Delete)
	userRoutes_v1.POST("/profile-image", userController.UploadProfileImage)

	return r
}
