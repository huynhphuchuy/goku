package authentication

import "github.com/gin-gonic/gin"

// ISessionController Interface
type ISessionController interface {
	CreateSession(*gin.Context)
	DeleteSession(*gin.Context)
}

// ISessionService Interface
type ISessionService interface {
	CreateSession(Session) error
	DeleteSession(Session)
	IsSessionExisted(Session) bool
}

// ISessionRepository Interface
type ISessionRepository interface {
	CreateSession(Session) error
	DeleteSession(Session)
	IsSessionExisted(Session) bool
}
