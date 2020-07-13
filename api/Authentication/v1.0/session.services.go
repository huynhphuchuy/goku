package authentication

var sessionRepository ISessionRepository = SessionRepository{}

// CreateSession Service
func (service SessionService) CreateSession(session Session) error {
	return sessionRepository.CreateSession(session)
}

// DeleteSession Service
func (service SessionService) DeleteSession(session Session) {
	sessionRepository.DeleteSession(session)
}

// IsSessionExisted Service
func (service SessionService) IsSessionExisted(session Session) bool {
	return sessionRepository.IsSessionExisted(session)
}
