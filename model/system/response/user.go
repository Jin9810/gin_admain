package response

import "gin-vue-admin-STL/entity"

type UserResponse struct {
	User entity.User `json:"user"`
}
