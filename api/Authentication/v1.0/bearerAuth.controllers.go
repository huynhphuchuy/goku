package authentication

import (
	"net/http"
	"strings"

	authConst "Gogin/api/Authentication/constants"
	authErrors "Gogin/api/Authentication/constants/errors"

	"github.com/gin-gonic/gin"
)

var bearerAuthService IBearerAuthService = &BearerAuthService{}
var _sessionService ISessionService = &SessionService{}

// VerifyToken Controller
func (controller BearerAuthController) VerifyToken(c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, authConst.BEARER_SCHEMA) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": authErrors.TOKEN_REQUIRED,
		})
		return
	}

	token := authHeader[len(authConst.BEARER_SCHEMA):]

	userID, err := bearerAuthService.VerifyToken(token)

	if err != nil || _sessionService.IsSessionExisted(Session{userID, token, 0}) == false {
		c.JSON(http.StatusForbidden, gin.H{
			"message": authErrors.INVALID_TOKEN,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
	})
	return
}
