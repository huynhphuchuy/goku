package authentication

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var sessionService ISessionService = &SessionService{}

// CreateSession Controller
func (controller SessionController) CreateSession(c *gin.Context) {

	userID, exist := c.Get("UserID")
	userToken, _ := c.Get("UserToken")

	if !exist {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Create new session
	session := Session{UserID: userID.(int), Token: userToken.(string), Timestamp: time.Now().Unix()}
	sessionService.CreateSession(session)

	c.AbortWithStatus(http.StatusOK)
}

// DeleteSession Controller
func (controller SessionController) DeleteSession(c *gin.Context) {

	_userID, exist := c.Get("UserID")
	userToken, _ := c.Get("UserToken")

	if !exist {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, _ := strconv.Atoi(_userID.(string))

	// Remove session
	session := Session{UserID: userID, Token: userToken.(string), Timestamp: time.Now().Unix()}
	sessionService.DeleteSession(session)

	c.AbortWithStatus(http.StatusNoContent)
}
