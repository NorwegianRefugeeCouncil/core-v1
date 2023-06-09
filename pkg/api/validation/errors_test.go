package validation

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyStringer interface {
	fmt.Stringer
}

type dummyStringer struct {
	s string
}

func (d dummyStringer) String() string {
	return d.s
}

func TestError_ErrorBody(t *testing.T) {
	tests := []struct {
		name string
		err  *Error
		want string
	}{
		{
			name: "nil",
			err:  nil,
			want: "<nil>",
		}, {
			name: "NotFound",
			err:  NotFound(NewPath("root"), 2),
			want: "root: Not Found: 2",
		}, {
			name: "Required",
			err:  Required(NewPath("root"), "detail"),
			want: "root: Required: detail",
		}, {
			name: "Duplicate",
			err:  Duplicate(NewPath("root"), 2),
			want: "root: Duplicate value: 2",
		}, {
			name: "Invalid",
			err:  Invalid(NewPath("root"), 2, "detail"),
			want: "root: Invalid value: 2: detail",
		}, {
			name: "NotSupported",
			err:  NotSupported(NewPath("root"), 3, []string{"1", "2"}),
			want: "root: Unsupported value: 3: supported values: \"1\", \"2\"",
		}, {
			name: "Forbidden",
			err:  Forbidden(NewPath("root"), "detail"),
			want: "root: Forbidden: detail",
		}, {
			name: "TooLong",
			err:  TooLong(NewPath("root"), 2, 1),
			want: "root: Too long: must have at most 1 bytes",
		}, {
			name: "TooLongMaxLength",
			err:  TooLongMaxLength(NewPath("root"), 2, 1),
			want: "root: Too long: value may not be longer than 1",
		}, {
			name: "TooLongMaxLength_minus",
			err:  TooLongMaxLength(NewPath("root"), 2, -1),
			want: "root: Too long: value is too long",
		}, {
			name: "TooShort",
			err:  TooShort(NewPath("root"), 2, 3),
			want: "root: Too short: must have at least 3 bytes",
		}, {
			name: "TooShortMinLength",
			err:  TooShortMinLength(NewPath("root"), 2, 3),
			want: "root: Too short: value may not be shorter than 3",
		}, {
			name: "TooShortMinLength_minus",
			err:  TooShortMinLength(NewPath("root"), 2, -1),
			want: "root: Too short: value is too short",
		}, {
			name: "InternalError",
			err:  InternalError(NewPath("root"), errors.New("detail")),
			want: "root: Internal error: detail",
		}, {
			name: "InternalError_nil",
			err:  InternalError(NewPath("root"), nil),
			want: "root: Internal error: <nil>",
		}, {
			name: "InternalError_empty",
			err:  InternalError(NewPath("root"), errors.New("")),
			want: "root: Internal error",
		}, {
			name: "TooMany",
			err:  TooMany(NewPath("root"), 10, 2),
			want: "root: Too many: 10: must have at most 2 items",
		}, {
			name: "TooMany_Singular",
			err:  TooMany(NewPath("root"), 10, 1),
			want: "root: Too many: 10: must have at most 1 item",
		}, {
			name: "TooMany_MaxQuantityMinus",
			err:  TooMany(NewPath("root"), 2, -1),
			want: "root: Too many: 2: too many items",
		}, {
			name: "TooMany_ActualQuantityMinus",
			err:  TooMany(NewPath("root"), -2, 2),
			want: "root: Too many: must have at most 2 items",
		}, {
			name: "TooFew",
			err:  TooFew(NewPath("root"), 2, 3),
			want: "root: Too few: 2: must have at least 3 items",
		}, {
			name: "TooFew_Singular",
			err:  TooFew(NewPath("root"), 0, 1),
			want: "root: Too few: 0: must have at least 1 item",
		}, {
			name: "TooFew_MinQuantityMinus",
			err:  TooFew(NewPath("root"), 2, -3),
			want: "root: Too few: 2: too few items",
		}, {
			name: "TooFew_ActualQuantityMinus",
			err:  TooFew(NewPath("root"), -2, 3),
			want: "root: Too few: must have at least 3 items",
		}, {
			name: "TypeInvalid",
			err:  TypeInvalid(NewPath("root"), 3, "detail"),
			want: "root: Invalid value type: 3: detail",
		}, {
			name: "IntValue",
			err:  TypeInvalid(NewPath("root"), 4, "detail"),
			want: "root: Invalid value type: 4: detail",
		}, {
			name: "Int32Value",
			err:  TypeInvalid(NewPath("root"), int32(5), "detail"),
			want: "root: Invalid value type: 5: detail",
		}, {
			name: "Int64Value",
			err:  TypeInvalid(NewPath("root"), int64(6), "detail"),
			want: "root: Invalid value type: 6: detail",
		}, {
			name: "Float32Value",
			err:  TypeInvalid(NewPath("root"), float32(7), "detail"),
			want: "root: Invalid value type: 7: detail",
		}, {
			name: "Float64Value",
			err:  TypeInvalid(NewPath("root"), float64(8), "detail"),
			want: "root: Invalid value type: 8: detail",
		}, {
			name: "StringValue",
			err:  TypeInvalid(NewPath("root"), "3", "detail"),
			want: "root: Invalid value type: \"3\": detail",
		}, {
			name: "BoolValue",
			err:  TypeInvalid(NewPath("root"), true, "detail"),
			want: "root: Invalid value type: true: detail",
		}, {
			name: "Stringer",
			err:  TypeInvalid(NewPath("root"), dummyStringer{s: "4"}, "detail"),
			want: "root: Invalid value type: \"4\": detail",
		}, {
			name: "StringerPointer",
			err:  TypeInvalid(NewPath("root"), &dummyStringer{s: "5"}, "detail"),
			want: "root: Invalid value type: \"5\": detail",
		}, {
			name: "StringerInterface",
			err:  TypeInvalid(NewPath("root"), DummyStringer(&dummyStringer{s: "6"}), "detail"),
			want: "root: Invalid value type: \"6\": detail",
		}, {
			name: "NilInterface",
			err:  TypeInvalid(NewPath("root"), DummyStringer(nil), "detail"),
			want: "root: Invalid value type: null: detail",
		}, {
			name: "NilStruct",
			err:  TypeInvalid(NewPath("root"), (*dummyStringer)(nil), "detail"),
			want: "root: Invalid value type: null: detail",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAggregate(t *testing.T) {
	tests := []struct {
		name string
		errs Aggregate
		want string
	}{
		{
			name: "single",
			errs: NewAggregate([]error{errors.New("single")}),
			want: "single",
		}, {
			name: "multiple",
			errs: NewAggregate([]error{errors.New("first"), errors.New("second")}),
			want: "[first, second]",
		}, {
			name: "duplicate",
			errs: NewAggregate([]error{errors.New("first"), errors.New("first")}),
			want: "first",
		}, {
			name: "nested",
			errs: NewAggregate([]error{
				errors.New("first"),
				NewAggregate([]error{errors.New("second")}),
			}),
			want: "[first, second]",
		}, {
			name: "nested duplicate",
			errs: NewAggregate([]error{
				errors.New("first"),
				NewAggregate([]error{errors.New("first")}),
			}),
			want: "first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.errs.Error(); got != tt.want {
				t.Errorf("aggregate.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

type DummyErr struct {
	msg string
}

func (e *DummyErr) Error() string {
	return e.msg
}

func (e *DummyErr) Is(target error) bool {
	if t, ok := target.(*DummyErr); ok {
		return e.msg == t.msg
	}
	return false
}

func TestAggregate_Is(t *testing.T) {

	err1 := func() error {
		return &DummyErr{msg: "first"}
	}
	err2 := func() error {
		return &DummyErr{msg: "second"}
	}
	err3 := func() error {
		return &DummyErr{msg: "third"}
	}

	tests := []struct {
		name string
		errs Aggregate
		err  error
		want bool
	}{
		{
			name: "single",
			errs: NewAggregate([]error{err1()}),
			err:  err1(),
			want: true,
		}, {
			name: "multiple",
			errs: NewAggregate([]error{err1(), err2()}),
			err:  err2(),
			want: true,
		}, {
			name: "duplicate",
			errs: NewAggregate([]error{err1(), err1()}),
			err:  err1(),
			want: true,
		}, {
			name: "nested",
			errs: NewAggregate([]error{
				err1(),
				NewAggregate([]error{err2()}),
			}),
			err:  err2(),
			want: true,
		}, {
			name: "no match",
			errs: NewAggregate([]error{err1(), err2()}),
			err:  err3(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.errs.Is(tt.err); got != tt.want {
				t.Errorf("aggregate.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisit(t *testing.T) {
	tests := []struct {
		name string
		errs []error
		want []error
	}{
		{
			name: "single",
			errs: []error{errors.New("single")},
			want: []error{errors.New("single")},
		}, {
			name: "multiple",
			errs: []error{errors.New("first"), errors.New("second")},
			want: []error{errors.New("first"), errors.New("second")},
		}, {
			name: "nested",
			errs: []error{
				errors.New("first"),
				NewAggregate([]error{
					errors.New("second"),
					errors.New("third"),
				}),
			},
			want: []error{
				errors.New("first"),
				errors.New("second"),
				errors.New("third"),
			},
		}, {
			name: "nested aggregate",
			errs: []error{
				errors.New("first"),
				NewAggregate([]error{
					errors.New("second"),
					NewAggregate([]error{
						errors.New("third"),
					}),
				}),
			},
			want: []error{
				errors.New("first"),
				errors.New("second"),
				errors.New("third"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []error
			agg := aggregate(tt.errs)
			agg.visit(func(err error) bool {
				got = append(got, err)
				return false
			})
			assert.Equal(t, tt.want, got)
		})
	}

}
