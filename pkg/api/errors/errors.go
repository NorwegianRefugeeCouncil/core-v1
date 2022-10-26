package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nrc-no/notcore/pkg/api/meta"
	"github.com/nrc-no/notcore/pkg/validation"
)

type StatusError struct {
	ErrStatus meta.Status
}

type APIStatus interface {
	Status() meta.Status
}

func (s *StatusError) Error() string {
	return s.ErrStatus.Message
}

func (s *StatusError) Status() meta.Status {
	return s.ErrStatus
}

func (s *StatusError) DebugError() (string, []interface{}) {
	if out, err := json.MarshalIndent(s.ErrStatus, "", "  "); err == nil {
		return "server response object: %s", []interface{}{string(out)}
	}
	return "server response object: %#v", []interface{}{s.ErrStatus}
}

func ErrorFrom(err error) *StatusError {
	switch t := err.(type) {
	case *StatusError:
		return t
	case APIStatus:
		return &StatusError{t.Status()}
	}
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusInternalServerError,
		Reason:  meta.StatusReasonInternalError,
		Message: err.Error(),
	}}
}

func HasStatusCause(err error, causeType validation.ErrorType) bool {
	_, ok := AsStatusCause(err, causeType)
	return ok
}

func AsStatusCause(err error, name validation.ErrorType) (meta.StatusCause, bool) {
	status, ok := err.(APIStatus)
	if (ok || errors.As(err, &status)) && status.Status().Details != nil {
		for _, cause := range status.Status().Details.Causes {
			if cause.Reason == name {
				return cause, true
			}
		}
	}
	return meta.StatusCause{}, false
}

func NewNotFound(resourceType, name string) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusNotFound,
		Reason: meta.StatusReasonNotFound,
		Details: &meta.StatusDetails{
			ResourceType: resourceType,
			Name:         name,
		},
		Message: fmt.Sprintf("%s %s not found", resourceType, name),
	}}
}

func NewAlreadyExists(resourceType, name string) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusConflict,
		Reason: meta.StatusReasonAlreadyExists,
		Details: &meta.StatusDetails{
			ResourceType: resourceType,
			Name:         name,
		},
		Message: fmt.Sprintf("%s %s already exists", resourceType, name),
	}}
}

func NewUnauthorized(reason string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusUnauthorized,
		Reason:  meta.StatusReasonUnauthorized,
		Message: reason,
	}}
}

func NewForbidden(resourceType, name string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusForbidden,
		Reason: meta.StatusReasonForbidden,
		Details: &meta.StatusDetails{
			ResourceType: resourceType,
			Name:         name,
		},
		Message: fmt.Sprintf("%s %s is forbidden: %v", resourceType, name, err),
	}}
}

func NewConflict(resourceType, name string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusConflict,
		Reason: meta.StatusReasonConflict,
		Details: &meta.StatusDetails{
			ResourceType: resourceType,
			Name:         name,
		},
		Message: fmt.Sprintf("operation cannot be fulfilled on %s %s: %v", resourceType, name, err),
	}}
}

func NewGone(message string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusGone,
		Reason:  meta.StatusReasonGone,
		Message: message,
	}}
}

func NewInvalid(resourceType, name string, fieldErrors validation.ErrorList) *StatusError {
	causes := make([]meta.StatusCause, 0, len(fieldErrors))
	for i := range fieldErrors {
		err := fieldErrors[i]
		causes = append(causes, meta.StatusCause{
			Reason:  err.Type,
			Message: err.ErrorBody(),
			Field:   err.Field,
		})
	}
	err := &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusUnprocessableEntity,
		Reason: meta.StatusReasonInvalid,
		Details: &meta.StatusDetails{
			Name:         name,
			ResourceType: resourceType,
			Causes:       causes,
		},
	}}
	aggregatedErrs := fieldErrors.ToAggregate()
	var identifier string
	if len(name) != 0 {
		identifier = fmt.Sprintf("%s %q", resourceType, name)
	} else {
		identifier = resourceType
	}
	if aggregatedErrs == nil {
		err.ErrStatus.Message = fmt.Sprintf("%s is invalid", identifier)
	} else {
		err.ErrStatus.Message = fmt.Sprintf("%s is invalid: %v", identifier, aggregatedErrs)
	}
	return err
}

func NewBadRequest(reason string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusBadRequest,
		Reason:  meta.StatusReasonBadRequest,
		Message: reason,
	}}
}

func NewServiceUnavailable(reason string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusServiceUnavailable,
		Reason:  meta.StatusReasonServiceUnavailable,
		Message: reason,
	}}
}

func NewMethodNotSupported(resourceType, action string) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusMethodNotAllowed,
		Reason: meta.StatusReasonMethodNotAllowed,
		Details: &meta.StatusDetails{
			ResourceType: resourceType,
		},
		Message: fmt.Sprintf("%s is not supported on resources of kind %s", action, resourceType),
	}}
}

func NewServerTimeout(resourceType string, operation string, seconds int64) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusGatewayTimeout,
		Reason: meta.StatusReasonServerTimeout,
		Details: &meta.StatusDetails{
			ResourceType:      resourceType,
			RetryAfterSeconds: seconds,
		},
		Message: fmt.Sprintf("The %s operation against %s could not be completed at this time. Please retry again", operation, resourceType),
	}}
}

func NewTimeoutError(message string, seconds int64) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusGatewayTimeout,
		Reason: meta.StatusReasonTimeout,
		Details: &meta.StatusDetails{
			RetryAfterSeconds: seconds,
		},
		Message: message,
	}}
}

func NewTooManyRequests(message string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusTooManyRequests,
		Reason:  meta.StatusReasonTooManyRequests,
		Message: message,
	}}
}

func NewRequestEntityTooLarge(message string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Code:    http.StatusRequestEntityTooLarge,
		Reason:  meta.StatusReasonRequestEntityTooLarge,
		Message: fmt.Sprintf("request entity too large: %s", message),
	}}
}

func NewInternalError(err error) *StatusError {
	return &StatusError{meta.Status{
		Status: meta.StatusFailure,
		Code:   http.StatusInternalServerError,
		Reason: meta.StatusReasonInternalError,
		Details: &meta.StatusDetails{
			Causes: []meta.StatusCause{{Message: err.Error()}},
		},
		Message: fmt.Sprintf("internal error occurred: %v", err),
	}}
}

func IsNotFound(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonNotFound, http.StatusNotFound)
}

func IsAlreadyExists(err error) bool {
	return ReasonForError(err) == meta.StatusReasonAlreadyExists
}

func IsConflict(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonConflict, http.StatusConflict)
}

func IsInvalid(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonInvalid, http.StatusUnprocessableEntity)
}

func IsGone(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonGone, http.StatusGone)
}

func IsNotAcceptable(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonNotAcceptable, http.StatusNotAcceptable)
}

func IsUnsupportedMediaType(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonUnsupportedMediaType, http.StatusUnsupportedMediaType)
}

func IsMethodNotSupported(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonMethodNotAllowed, http.StatusMethodNotAllowed)
}

func IsServiceUnavailable(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonServiceUnavailable, http.StatusServiceUnavailable)
}

func IsBadRequest(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonBadRequest, http.StatusBadRequest)
}

func IsUnauthorized(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonUnauthorized, http.StatusUnauthorized)
}

func IsForbidden(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonForbidden, http.StatusForbidden)
}

func IsTimeout(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonTimeout, http.StatusGatewayTimeout)
}

func IsTooManyRequests(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonTooManyRequests, http.StatusTooManyRequests)
}

func IsRequestEntityTooLarge(err error) bool {
	return isErrorOrCode(err, meta.StatusReasonRequestEntityTooLarge, http.StatusRequestEntityTooLarge)
}

func IsUnexpectedServerError(err error) bool {
	status, ok := err.(APIStatus)
	if (ok || errors.As(err, &status)) && status.Status().Details != nil {
		for _, cause := range status.Status().Details.Causes {
			if cause.Reason == validation.ErrorTypeUnexpectedServerResponse {
				return true
			}
		}
	}
	return false
}

func isErrorOrCode(err error, reason meta.StatusReason, httpStatus int32) bool {
	errReason, code := reasonAndCodeForError(err)
	if errReason == reason {
		return true
	}
	if _, ok := knownReasons[errReason]; !ok && code == httpStatus {
		return true
	}
	return false
}

func ReasonForError(err error) meta.StatusReason {
	if status, ok := err.(APIStatus); ok || errors.As(err, &status) {
		return status.Status().Reason
	}
	return meta.StatusReasonUnknown
}

func reasonAndCodeForError(err error) (meta.StatusReason, int32) {
	if status, ok := err.(APIStatus); ok || errors.As(err, &status) {
		return status.Status().Reason, status.Status().Code
	}
	return meta.StatusReasonUnknown, 0
}

var knownReasons = map[meta.StatusReason]struct{}{
	// StatusReasonUnknown:                 {},
	meta.StatusReasonUnauthorized:          {},
	meta.StatusReasonForbidden:             {},
	meta.StatusReasonNotFound:              {},
	meta.StatusReasonAlreadyExists:         {},
	meta.StatusReasonConflict:              {},
	meta.StatusReasonGone:                  {},
	meta.StatusReasonInvalid:               {},
	meta.StatusReasonServerTimeout:         {},
	meta.StatusReasonTimeout:               {},
	meta.StatusReasonTooManyRequests:       {},
	meta.StatusReasonBadRequest:            {},
	meta.StatusReasonMethodNotAllowed:      {},
	meta.StatusReasonNotAcceptable:         {},
	meta.StatusReasonRequestEntityTooLarge: {},
	meta.StatusReasonUnsupportedMediaType:  {},
	meta.StatusReasonInternalError:         {},
	meta.StatusReasonExpired:               {},
	meta.StatusReasonServiceUnavailable:    {},
}
