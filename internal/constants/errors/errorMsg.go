package commonErrors

type ErrorMsg string

const (
	UNEXPECTED_SERVER_ERROR ErrorMsg = "Unexpected server error!"
	INVALID_DATA_FORMAT     ErrorMsg = "Invalid data format!"
)
