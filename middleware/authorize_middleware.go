package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/repository"
	"github.com/sandriansyafridev/crowdfounding/service"
)

func AuthorizationMiddleware(userRepository repository.UserRepository, jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {

		//check authoriazation header
		if Authorization := c.GetHeader("Authorization"); Authorization == "" {
			responseFail := response.ResponseFail("Unauthorize", errors.New("Not allow to access this page"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseFail)
		} else {

			//parse token
			if token, err := jwtService.ParseToken(Authorization); err != nil {
				responseFail := response.ResponseFail("Unauthorize", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, responseFail)
			} else {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					UserID := int64(claims["user_id"].(float64))
					if user, err := userRepository.FindByID(uint64(UserID)); err != nil {
						responseFail := response.ResponseFail("User not found", err)
						c.AbortWithStatusJSON(http.StatusInternalServerError, responseFail)
					} else {
						c.Set("currentUser", user)
					}
				} else {
					responseFail := response.ResponseFail("Invalid token", err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, responseFail)
				}
			}

		}

	}
}
