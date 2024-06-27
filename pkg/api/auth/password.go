package auth

import (
	"net/url"
)

type SessionAuthorizer struct {
	sessionID string
}

func NewSessionAuthorizer(sessionID string) *SessionAuthorizer {
	return &SessionAuthorizer{sessionID: sessionID}
}

func (a *SessionAuthorizer) Apply(values url.Values) {
	values.Set("_sid", a.sessionID)
}
