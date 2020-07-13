package authentication

import "github.com/gin-gonic/gin"

// IBearerAuthController Interface
type IBearerAuthController interface {
	VerifyToken(*gin.Context)
}

// IBearerAuthService Interface
type IBearerAuthService interface {
	IssueJwtToken(int, string, string, string) string
	VerifyToken(string) (int, error)
}

// IBearerAuthRepository Interface
type IBearerAuthRepository interface {
}
