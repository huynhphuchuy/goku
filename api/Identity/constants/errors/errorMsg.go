package errors

// ErrorMsg String
type ErrorMsg string

const (
	USER_EXISTS         ErrorMsg = "User exists!"
	TOKEN_REQUIRED      ErrorMsg = "Token is required!"
	INVALID_TOKEN       ErrorMsg = "Token is invalid or expired!"
	USER_NOT_EXISTS     ErrorMsg = "User not exist!"
	INVALID_CREDENTIALS ErrorMsg = "Credentials are invalid!"
	INVALID_USER_ID     ErrorMsg = "Invalid User ID!"
)
