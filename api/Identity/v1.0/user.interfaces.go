package account

import "github.com/gin-gonic/gin"

// IUserController Interface
type IUserController interface {
	Register(*gin.Context)
	Verify(*gin.Context)
	Signin(*gin.Context)
	Signout(*gin.Context)
}

// IUserService Interface
type IUserService interface {
	Register(*User) (int, string, error)
	Verify(string) error
	Signin(*UserSignin) (int, string, error)
}

// IUserRepository Interface
type IUserRepository interface {
	Register(User) (int, error)
	Verify(string) error
	ValidateCredentials(UserSignin) bool
	GetUserByIdentity(string, string) *User
}
