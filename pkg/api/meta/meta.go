package meta

import (
	"github.com/nrc-no/notcore/pkg/api/validation"
)

type Status struct {
	Status  string         `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
	Reason  StatusReason   `json:"reason,omitempty"`
	Details *StatusDetails `json:"details,omitempty"`
	Code    int32          `json:"code,omitempty"`
}

type StatusDetails struct {
	Name              string        `json:"name,omitempty"`
	Causes            []StatusCause `json:"causes,omitempty"`
	ResourceType      string        `json:"type,omitempty"`
	RetryAfterSeconds int64         `json:"retryAfterSeconds,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusFailure = "failure"
)

type StatusReason string

const (
	StatusReasonUnknown               StatusReason = "Unknown"
	StatusReasonUnauthorized          StatusReason = "Unauthorized"
	StatusReasonForbidden             StatusReason = "Forbidden"
	StatusReasonNotFound              StatusReason = "NotFound"
	StatusReasonAlreadyExists         StatusReason = "AlreadyExists"
	StatusReasonConflict              StatusReason = "Conflict"
	StatusReasonGone                  StatusReason = "Gone"
	StatusReasonInvalid               StatusReason = "Invalid"
	StatusReasonServerTimeout         StatusReason = "ServerTimeout"
	StatusReasonTimeout               StatusReason = "Timeout"
	StatusReasonTooManyRequests       StatusReason = "TooManyRequests"
	StatusReasonBadRequest            StatusReason = "BadRequest"
	StatusReasonMethodNotAllowed      StatusReason = "MethodNotAllowed"
	StatusReasonNotAcceptable         StatusReason = "NotAcceptable"
	StatusReasonRequestEntityTooLarge StatusReason = "RequestEntityTooLarge"
	StatusReasonUnsupportedMediaType  StatusReason = "UnsupportedMediaType"
	StatusReasonInternalError         StatusReason = "InternalError"
	StatusReasonExpired               StatusReason = "Expired"
	StatusReasonServiceUnavailable    StatusReason = "ServiceUnavailable"
)

type StatusCause struct {
	Reason  validation.ErrorType `json:"reason,omitempty"`
	Message string               `json:"message,omitempty"`
	Field   string               `json:"field,omitempty"`
}
