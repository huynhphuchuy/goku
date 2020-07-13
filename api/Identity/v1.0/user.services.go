package account

import (
	"errors"

	userErrors "Gogin/api/Identity/constants/errors"

	authenticationV1_0 "Gogin/api/Authentication/v1.0"
)

var userRepository IUserRepository = UserRepository{}
var sessionService authenticationV1_0.ISessionService = authenticationV1_0.SessionService{}
var bearerAuth authenticationV1_0.IBearerAuthService = authenticationV1_0.BearerAuthService{}

// Register Service
func (service UserService) Register(u *User) (int, string, error) {
	accountUser := User(*u)
	user := GetUserByIdentity(accountUser.Username, accountUser.Email)

	// User not exist
	if user != nil {
		return 0, "", errors.New(string(userErrors.USER_EXISTS))
	}

	// Issue Jwt
	tokenString := bearerAuth.IssueJwtToken(accountUser.ID, accountUser.Username, accountUser.Email, accountUser.Name)

	// Hash password
	u.Password = HashPassword(u.Password)

	// Register user
	userID, err := userRepository.Register(*u)

	// Send verification email
	u.SendVerificationEmail()

	return userID, tokenString, err
}

// Verify Service
func (service UserService) Verify(email string) error {
	userRepository.Verify(email)
	return nil
}

// Signin Service
func (service UserService) Signin(userSignin *UserSignin) (int, string, error) {
	user := GetUserByIdentity(userSignin.Username, "")

	// User not exist
	if user == nil {
		return 0, "", errors.New(string(userErrors.USER_NOT_EXISTS))
	}

	// Issue Jwt
	tokenString := bearerAuth.IssueJwtToken(user.ID, userSignin.Username, user.Email, user.Name)

	// Hash password
	userSignin.Password = HashPassword(userSignin.Password)

	// Validate credentials
	if userRepository.ValidateCredentials(*userSignin) == false {
		return 0, "", errors.New(string((userErrors.INVALID_CREDENTIALS)))
	}

	return user.ID, tokenString, nil
}

// GetUserByIdentity Service
func GetUserByIdentity(username string, email string) *User {
	user := userRepository.GetUserByIdentity(username, email)
	return user
}
