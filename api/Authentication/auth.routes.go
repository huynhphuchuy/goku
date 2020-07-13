package authroutes

import (
	authv1_0 "Gogin/api/Authentication/v1.0"
	"github.com/gin-gonic/gin"
)

var bearerAuthControllerV1_0 authv1_0.IBearerAuthController = authv1_0.BearerAuthController{}

// Routes Group
func Routes(r *gin.Engine) {
	client := r.Group("/auth")
	{
		client.GET("/bearer/v1.0/verify", bearerAuthControllerV1_0.VerifyToken)
	}
}
