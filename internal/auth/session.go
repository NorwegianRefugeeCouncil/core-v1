package auth

import (
	"fmt"
	"time"
)

type Session interface {
	IsAuthenticated() bool
	GetUserID() string
	GetUserGroups() []string
	GetUserEmail() string
	GetIssuer() string
	GetSubject() string
	GetExpiration() time.Time
	GetIssuedAt() time.Time
	WillExpireIn(duration time.Duration) bool
}

type session struct {
	isAuthenticated bool
	userGroups      []string
	userEmail       string
	issuer          string
	subject         string
	expiration      time.Time
	issuedAt        time.Time
	nrcOrganisation string
}

func (s session) IsAuthenticated() bool {
	return s.isAuthenticated
}

func (s session) WillExpireIn(duration time.Duration) bool {
	return s.expiration.Sub(time.Now()) < duration
}

func (s session) GetUserID() string {
	return fmt.Sprintf("%s:%s", s.issuer, s.subject)
}

func (s session) GetUserGroups() []string {
	var ret = make([]string, len(s.userGroups))
	copy(ret, s.userGroups)
	return ret
}

func (s session) GetUserEmail() string {
	return s.userEmail
}

func (s session) GetIssuer() string {
	return s.issuer
}

func (s session) GetSubject() string {
	return s.subject
}

func (s session) GetExpiration() time.Time {
	return s.expiration
}

func (s session) GetIssuedAt() time.Time {
	return s.issuedAt
}

func NewAuthenticatedSession(
	userGroups []string,
	nrcOrganisation string,
	userEmail string,
	issuer string,
	subject string,
	expiration time.Time,
	issuedAt time.Time,
) Session {
	return &session{
		isAuthenticated: true,
		userGroups:      userGroups,
		userEmail:       userEmail,
		issuer:          issuer,
		subject:         subject,
		expiration:      expiration,
		issuedAt:        issuedAt,
		nrcOrganisation: nrcOrganisation,
	}
}
