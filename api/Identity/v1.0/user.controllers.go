package account

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	userErrors "Gogin/api/Identity/constants/errors"
	commonErrors "Gogin/internal/constants/errors"
)

var userService IUserService = &UserService{}

// Register Controller
func (controller UserController) Register(c *gin.Context) {
	var userModel = new(User)
	if bError := c.BindJSON(userModel); bError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": commonErrors.INVALID_DATA_FORMAT,
		})
		c.Abort()
		return
	}

	userID, token, err := userService.Register(userModel)
	if err == nil {
		c.Set("UserToken", token)
		c.Set("UserID", userID)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		c.Next()
		return
	}

	MapErrorsHandler(c, err.Error(), []string{
		string(userErrors.USER_EXISTS),
		string(commonErrors.UNEXPECTED_SERVER_ERROR),
	})
}

// Verify Controller
func (controller UserController) Verify(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": userErrors.TOKEN_REQUIRED,
		})
		return
	}

	if email, valid := ValidateVerificationToken(token); valid {
		userService.Verify(email)
		c.Status(http.StatusNoContent)
		c.Abort()
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": userErrors.INVALID_TOKEN,
	})
	c.Abort()
	return
}

// Signin Controller
func (controller UserController) Signin(c *gin.Context) {
	var userSigninModel = new(UserSignin)
	if bError := c.BindJSON(userSigninModel); bError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": commonErrors.INVALID_DATA_FORMAT,
		})
		c.Abort()
		return
	}

	userID, token, err := userService.Signin(userSigninModel)
	if err == nil {
		c.Set("UserToken", token)
		c.Set("UserID", userID)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		c.Next()
		return
	}

	MapErrorsHandler(c, err.Error(), []string{
		string(userErrors.USER_NOT_EXISTS),
		string(userErrors.INVALID_CREDENTIALS),
		string(commonErrors.UNEXPECTED_SERVER_ERROR),
	})
}

// Signout Controller
func (controller UserController) Signout(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer "
	userID := c.Request.Header.Get("x-user-id")
	authHeader := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": userErrors.TOKEN_REQUIRED,
		})
		return
	}

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": userErrors.INVALID_USER_ID,
		})
		return
	}

	token := authHeader[len(BEARER_SCHEMA):]

	c.Set("UserToken", token)
	c.Set("UserID", userID)
	c.Next()
	return
}
