package account

import (
	t "Gogin/internal/helpers/emailTemplate"
	"Gogin/internal/platform/mailer"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MapErrorsHandler Util
func MapErrorsHandler(c *gin.Context, cError string, errors []string) {
	for _, err := range errors {
		if err == cError {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			c.Abort()
			return
		}
	}
}

// GenerateVerificationToken Util
func GenerateVerificationToken(email string, exp int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": email,
		"exp":   time.Now().UTC().Add(time.Duration(exp) * time.Second).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	return base64.StdEncoding.EncodeToString([]byte(tokenString))
}

// ValidateVerificationToken Util
func ValidateVerificationToken(tokenString string) (string, bool) {
	decodedToken, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return "", false
	}
	token, err := jwt.Parse(string(decodedToken), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims["Email"].(string), true
		}
	}
	return "", false
}

// HashPassword Util
func HashPassword(password string) string {
	sha := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(sha[:])
}

// SendVerificationEmail Util
func (u User) SendVerificationEmail() {
	go mailer.SendEmail(u.Email, "Verification Email!", t.Verification{
		Name: u.Name,
		Intros: []string{
			"Welcome to Golang! We're very excited to have you on board.",
		},
		Instructions: "To get started with Golang, please click here:",
		Button: t.Button{
			Color: "#22BC66",
			Text:  "Verify your account",
			Link:  "http://" + os.Getenv("HOST_NAME") + "/auth/bearer/v1.0/verify/" + GenerateVerificationToken(u.Email, 3600),
		},
		Outros: []string{
			"Need help, or have questions? Just reply to this email, we'd love to help.",
		},
	}.Init())
}
