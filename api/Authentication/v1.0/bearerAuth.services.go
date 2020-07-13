package authentication

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// IssueJwtToken Service
func (service BearerAuthService) IssueJwtToken(userID int, username string, email string, name string) string {
	duration, _ := strconv.Atoi(os.Getenv("AUTH_EXP"))
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"UserID":   userID,
		"Username": username,
		"Name":     name,
		"Email":    email,
		"exp":      time.Now().UTC().Add(time.Duration(duration) * time.Second).Unix(),
	})

	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("PRIVATE_TOKEN_SECRET")))

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(signKey)

	return tokenString
}

// VerifyToken Service
func (service BearerAuthService) VerifyToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("")
		}

		signKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("PUBLIC_TOKEN_SECRET")))
		return signKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["UserID"].(float64)), nil
	}

	return 0, errors.New("")
}
