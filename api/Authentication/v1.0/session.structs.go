package authentication

// SessionController Struct
type SessionController struct {
}

// SessionService Struct
type SessionService struct {
}

// SessionRepository Struct
type SessionRepository struct {
}

// Session Struct
type Session struct {
	UserID    int
	Token     string
	Timestamp int64
}
