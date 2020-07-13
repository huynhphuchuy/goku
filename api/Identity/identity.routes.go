package accountroutes

import (
	authv1_0 "Gogin/api/Authentication/v1.0"
	accountv1_0 "Gogin/api/Identity/v1.0"

	"github.com/gin-gonic/gin"
)

var userControllerV1_0 accountv1_0.IUserController = accountv1_0.UserController{}
var sessionControllerV1_0 authv1_0.ISessionController = authv1_0.SessionController{}

// Routes Group
func Routes(r *gin.Engine) {
	client := r.Group("/identity")
	{
		client.POST("/user/v1.0/register", userControllerV1_0.Register, sessionControllerV1_0.CreateSession)
		client.POST("/user/v1.0/signin", userControllerV1_0.Signin, sessionControllerV1_0.CreateSession)
		client.GET("/user/v1.0/signout", userControllerV1_0.Signout, sessionControllerV1_0.DeleteSession)
		client.GET("/user/v1.0/verify/:token", userControllerV1_0.Verify)
	}
}
