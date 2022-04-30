package format

import (
	"time"

	"github.com/sandriansyafridev/crowdfounding/model/entity"
)

type UserResponse struct {
	ID               uint64    `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Token            string    `json:"token,omitempty"`
	PathProfileImage string    `json:"path_profile_image"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func ToUserResponse(user entity.User) (userResponse UserResponse) {
	userResponse.ID = user.ID
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Token = user.Token
	userResponse.PathProfileImage = user.PathProfileImage
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt
	return userResponse
}

func ToUsersResponse(users []entity.User) (usersResponse []UserResponse) {

	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(user))
	}

	return usersResponse
}
