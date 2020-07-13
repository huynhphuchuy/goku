package errors

// ErrorMsg string
type ErrorMsg string

const (
	TOKEN_REQUIRED ErrorMsg = "Token is required!"
	INVALID_TOKEN  ErrorMsg = "Token is invalid or expired!"
)
