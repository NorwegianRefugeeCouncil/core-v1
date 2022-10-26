package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/nrc-no/notcore/pkg/meta"
	"github.com/nrc-no/notcore/pkg/validation"
	"github.com/stretchr/testify/assert"
)

func TestConstructors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "NewBadRequest",
			err:  NewBadRequest("bad request"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    400,
				Reason:  meta.StatusReasonBadRequest,
				Message: "bad request",
			}},
		}, {
			name: "NewUnauthorized",
			err:  NewUnauthorized("unauthorized"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    401,
				Reason:  meta.StatusReasonUnauthorized,
				Message: "unauthorized",
			}},
		}, {
			name: "NewForbidden",
			err:  NewForbidden("bar", "foo", errors.New("fuzz")),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    403,
				Reason:  meta.StatusReasonForbidden,
				Details: &meta.StatusDetails{ResourceType: "bar", Name: "foo"},
				Message: "bar foo is forbidden: fuzz",
			}},
		}, {
			name: "NewNotFound",
			err:  NewNotFound("bar", "foo"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    404,
				Reason:  meta.StatusReasonNotFound,
				Details: &meta.StatusDetails{ResourceType: "bar", Name: "foo"},
				Message: "bar foo not found",
			}},
		}, {
			name: "NewAlreadyExists",
			err:  NewAlreadyExists("bar", "foo"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    409,
				Reason:  meta.StatusReasonAlreadyExists,
				Details: &meta.StatusDetails{ResourceType: "bar", Name: "foo"},
				Message: "bar foo already exists",
			}},
		}, {
			name: "NewConflict",
			err:  NewConflict("bar", "foo", errors.New("fuzz")),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    409,
				Reason:  meta.StatusReasonConflict,
				Details: &meta.StatusDetails{ResourceType: "bar", Name: "foo"},
				Message: "operation cannot be fulfilled on bar foo: fuzz",
			}},
		}, {
			name: "NewInternalError",
			err:  NewInternalError(errors.New("fuzz")),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    500,
				Reason:  meta.StatusReasonInternalError,
				Message: "internal error occurred: fuzz",
				Details: &meta.StatusDetails{Causes: []meta.StatusCause{{Message: "fuzz"}}},
			}},
		}, {
			name: "NewGone",
			err:  NewGone("bar"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    410,
				Reason:  meta.StatusReasonGone,
				Message: "bar",
			}},
		}, {
			name: "NewInvalid",
			err: NewInvalid("bar", "foo", validation.ErrorList{
				validation.Required(validation.NewPath("snizz"), "fuzz"),
			}),
			want: &StatusError{ErrStatus: meta.Status{
				Status: meta.StatusFailure,
				Code:   422,
				Reason: meta.StatusReasonInvalid,
				Details: &meta.StatusDetails{
					ResourceType: "bar",
					Name:         "foo",
					Causes: []meta.StatusCause{
						{
							Field:   "snizz",
							Message: "Required: fuzz",
							Reason:  validation.ErrorTypeRequired,
						},
					},
				},
				Message: "bar \"foo\" is invalid: snizz: Required: fuzz",
			}},
		}, {
			name: "NewInvalid (nil errors)",
			err:  NewInvalid("bar", "foo", nil),
			want: &StatusError{ErrStatus: meta.Status{
				Status: meta.StatusFailure,
				Code:   422,
				Reason: meta.StatusReasonInvalid,
				Details: &meta.StatusDetails{
					ResourceType: "bar",
					Name:         "foo",
					Causes:       []meta.StatusCause{},
				},
				Message: "bar \"foo\" is invalid",
			}},
		}, {
			name: "NewServiceUnavailable",
			err:  NewServiceUnavailable("bar"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    503,
				Reason:  meta.StatusReasonServiceUnavailable,
				Message: "bar",
			}},
		}, {
			name: "NewMethodNotSupported",
			err:  NewMethodNotSupported("bar", "fuzz"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    405,
				Reason:  meta.StatusReasonMethodNotAllowed,
				Details: &meta.StatusDetails{ResourceType: "bar"},
				Message: "fuzz is not supported on resources of kind bar",
			}},
		}, {
			name: "NewServerTimeout",
			err:  NewServerTimeout("bar", "fuzz", 1),
			want: &StatusError{ErrStatus: meta.Status{
				Status: meta.StatusFailure,
				Code:   504,
				Reason: meta.StatusReasonServerTimeout,
				Details: &meta.StatusDetails{
					ResourceType:      "bar",
					RetryAfterSeconds: 1,
				},
				Message: "The fuzz operation against bar could not be completed at this time. Please retry again",
			}},
		}, {
			name: "NewTooManyRequests",
			err:  NewTooManyRequests("foo"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    429,
				Reason:  meta.StatusReasonTooManyRequests,
				Message: "foo",
			}},
		}, {
			name: "NewRequestEntityTooLarge",
			err:  NewRequestEntityTooLarge("foo"),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    413,
				Reason:  meta.StatusReasonRequestEntityTooLarge,
				Message: "request entity too large: foo",
			}},
		}, {
			name: "NewTimeoutError",
			err:  NewTimeoutError("foo", 1),
			want: &StatusError{ErrStatus: meta.Status{
				Status:  meta.StatusFailure,
				Code:    504,
				Reason:  meta.StatusReasonTimeout,
				Message: "foo",
				Details: &meta.StatusDetails{
					RetryAfterSeconds: 1,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.err)
		})
	}
}

func TestErrorIs(t *testing.T) {

	errWithReason := func(reason meta.StatusReason) error {
		return &StatusError{ErrStatus: meta.Status{Reason: reason}}
	}
	errWithCode := func(code int32) error {
		return &StatusError{ErrStatus: meta.Status{Code: code}}
	}

	tests := []struct {
		name string
		err  error
		fn   func(err error) bool
		want bool
	}{
		{
			name: "IsAlreadyExists (reason)",
			err:  errWithReason(meta.StatusReasonAlreadyExists),
			fn:   IsAlreadyExists,
			want: true,
		}, {
			name: "IsAlreadyExists (wrong type)",
			err:  errWithReason(""),
			fn:   IsAlreadyExists,
			want: false,
		}, {
			name: "IsBadRequest (reason)",
			err:  errWithReason(meta.StatusReasonBadRequest),
			fn:   IsBadRequest,
			want: true,
		}, {
			name: "IsBadRequest (http)",
			err:  errWithCode(http.StatusBadRequest),
			fn:   IsBadRequest,
			want: true,
		}, {
			name: "IsBadRequest (wrong type)",
			err:  errWithCode(200),
			fn:   IsBadRequest,
			want: false,
		}, {
			name: "IsConflict (reason)",
			err:  errWithReason(meta.StatusReasonConflict),
			fn:   IsConflict,
			want: true,
		}, {
			name: "IsConflict (http)",
			err:  errWithCode(http.StatusConflict),
			fn:   IsConflict,
			want: true,
		}, {
			name: "IsConflict (wrong type)",
			err:  errWithCode(200),
			fn:   IsConflict,
			want: false,
		}, {
			name: "IsForbidden (reason)",
			err:  errWithReason(meta.StatusReasonForbidden),
			fn:   IsForbidden,
			want: true,
		}, {
			name: "IsForbidden (http)",
			err:  errWithCode(http.StatusForbidden),
			fn:   IsForbidden,
			want: true,
		}, {
			name: "IsForbidden (wrong type)",
			err:  errWithCode(200),
			fn:   IsForbidden,
			want: false,
		}, {
			name: "IsInvalid (reason)",
			err:  errWithReason(meta.StatusReasonInvalid),
			fn:   IsInvalid,
			want: true,
		}, {
			name: "IsInvalid (http)",
			err:  errWithCode(http.StatusUnprocessableEntity),
			fn:   IsInvalid,
			want: true,
		}, {
			name: "IsInvalid (wrong type)",
			err:  errWithCode(200),
			fn:   IsInvalid,
			want: false,
		}, {
			name: "IsGone (reason)",
			err:  errWithReason(meta.StatusReasonGone),
			fn:   IsGone,
			want: true,
		}, {
			name: "IsGone (http)",
			err:  errWithCode(http.StatusGone),
			fn:   IsGone,
			want: true,
		}, {
			name: "IsGone (wrong type)",
			err:  errWithCode(200),
			fn:   IsGone,
			want: false,
		}, {
			name: "IsRequestEntityTooLarge (reason)",
			err:  errWithReason(meta.StatusReasonRequestEntityTooLarge),
			fn:   IsRequestEntityTooLarge,
			want: true,
		}, {
			name: "IsRequestEntityTooLarge (code)",
			err:  errWithCode(http.StatusRequestEntityTooLarge),
			fn:   IsRequestEntityTooLarge,
			want: true,
		}, {
			name: "IsRequestEntityTooLarge (wrong type)",
			err:  errWithCode(200),
			fn:   IsRequestEntityTooLarge,
			want: false,
		}, {
			name: "IsUnexpectedServerError",
			err: &StatusError{
				ErrStatus: meta.Status{
					Details: &meta.StatusDetails{
						Causes: []meta.StatusCause{
							{
								Reason: validation.ErrorTypeUnexpectedServerResponse,
							},
						},
					},
				},
			},
			fn:   IsUnexpectedServerError,
			want: true,
		}, {
			name: "IsUnexpectedServerError (wrong type)",
			err:  errWithCode(200),
			fn:   IsUnexpectedServerError,
			want: false,
		}, {
			name: "IsMethodNotSupported (reason)",
			err:  errWithReason(meta.StatusReasonMethodNotAllowed),
			fn:   IsMethodNotSupported,
			want: true,
		}, {
			name: "IsMethodNotSupported (http)",
			err:  errWithCode(http.StatusMethodNotAllowed),
			fn:   IsMethodNotSupported,
			want: true,
		}, {
			name: "IsMethodNotSupported (wrong type)",
			err:  errWithCode(200),
			fn:   IsMethodNotSupported,
			want: false,
		}, {
			name: "IsUnauthorized (reason)",
			err:  errWithReason(meta.StatusReasonUnauthorized),
			fn:   IsUnauthorized,
			want: true,
		}, {
			name: "IsUnauthorized (http)",
			err:  errWithCode(http.StatusUnauthorized),
			fn:   IsUnauthorized,
			want: true,
		}, {
			name: "IsUnauthorized (wrong type)",
			err:  errWithCode(200),
			fn:   IsUnauthorized,
			want: false,
		}, {
			name: "IsNotAcceptable (reason)",
			err:  errWithReason(meta.StatusReasonNotAcceptable),
			fn:   IsNotAcceptable,
			want: true,
		}, {
			name: "IsNotAcceptable (http)",
			err:  errWithCode(http.StatusNotAcceptable),
			fn:   IsNotAcceptable,
			want: true,
		}, {
			name: "IsNotAcceptable (wrong type)",
			err:  errWithCode(200),
			fn:   IsNotAcceptable,
			want: false,
		}, {
			name: "IsUnsupportedMediaType (reason)",
			err:  errWithReason(meta.StatusReasonUnsupportedMediaType),
			fn:   IsUnsupportedMediaType,
			want: true,
		}, {
			name: "IsUnsupportedMediaType (http)",
			err:  errWithCode(http.StatusUnsupportedMediaType),
			fn:   IsUnsupportedMediaType,
			want: true,
		}, {
			name: "IsUnsupportedMediaType (wrong type)",
			err:  errWithCode(200),
			fn:   IsUnsupportedMediaType,
			want: false,
		}, {
			name: "IsServiceUnavailable (reason)",
			err:  errWithReason(meta.StatusReasonServiceUnavailable),
			fn:   IsServiceUnavailable,
			want: true,
		}, {
			name: "IsServiceUnavailable (http)",
			err:  errWithCode(http.StatusServiceUnavailable),
			fn:   IsServiceUnavailable,
			want: true,
		}, {
			name: "IsServiceUnavailable (wrong type)",
			err:  errWithCode(200),
			fn:   IsServiceUnavailable,
			want: false,
		}, {
			name: "IsTooManyRequests (reason)",
			err:  errWithReason(meta.StatusReasonTooManyRequests),
			fn:   IsTooManyRequests,
			want: true,
		}, {
			name: "IsTooManyRequests (http)",
			err:  errWithCode(http.StatusTooManyRequests),
			fn:   IsTooManyRequests,
			want: true,
		}, {
			name: "IsTooManyRequests (wrong type)",
			err:  errWithCode(200),
			fn:   IsTooManyRequests,
			want: false,
		}, {
			name: "IsNotFound (reason)",
			err:  errWithReason(meta.StatusReasonNotFound),
			fn:   IsNotFound,
			want: true,
		}, {
			name: "IsNotFound (http)",
			err:  errWithCode(http.StatusNotFound),
			fn:   IsNotFound,
			want: true,
		}, {
			name: "IsNotFound (wrong type)",
			err:  errWithCode(200),
			fn:   IsNotFound,
			want: false,
		}, {
			name: "IsTimeout (reason)",
			err:  errWithReason(meta.StatusReasonTimeout),
			fn:   IsTimeout,
			want: true,
		}, {
			name: "IsTimeout (http)",
			err:  errWithCode(http.StatusGatewayTimeout),
			fn:   IsTimeout,
			want: true,
		}, {
			name: "IsTimeout (wrong type)",
			err:  errWithCode(200),
			fn:   IsTimeout,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fn(tt.err))
		})
	}
}

func Test_reasonAndCodeForErrorUnknown(t *testing.T) {
	actualReason, actualCode := reasonAndCodeForError(errors.New("bla"))
	assert.Equal(t, meta.StatusReasonUnknown, actualReason)
	assert.Equal(t, int32(0), actualCode)
}

func Test_ReasonForError_Unknown(t *testing.T) {
	actualReason := ReasonForError(errors.New("bla"))
	assert.Equal(t, meta.StatusReasonUnknown, actualReason)
}

func TestStatusError_Error(t *testing.T) {
	err := &StatusError{
		ErrStatus: meta.Status{
			Reason:  meta.StatusReasonAlreadyExists,
			Message: "foo",
		},
	}
	assert.Equal(t, "foo", err.Error())
}

func TestHasStatusCause(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		cause    validation.ErrorType
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			cause:    validation.ErrorTypeTooShort,
			expected: false,
		}, {
			name:     "no status error",
			err:      errors.New("foo"),
			cause:    validation.ErrorTypeTooShort,
			expected: false,
		}, {
			name:     "no status cause",
			err:      &StatusError{ErrStatus: meta.Status{}},
			cause:    validation.ErrorTypeTooShort,
			expected: false,
		}, {
			name: "status cause",
			err: &StatusError{
				meta.Status{
					Details: &meta.StatusDetails{
						Causes: []meta.StatusCause{{Reason: validation.ErrorTypeTooShort}},
					},
				},
			},
			cause:    validation.ErrorTypeTooShort,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, HasStatusCause(tt.err, tt.cause))
		})
	}

}
