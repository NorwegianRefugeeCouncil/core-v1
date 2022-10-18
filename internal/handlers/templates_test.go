package handlers

import (
	"testing"

	"github.com/nrc-no/notcore/internal/validation"
	"github.com/stretchr/testify/assert"
)

func TestViewDataGetErrors(t *testing.T) {

	testCases := []struct {
		name     string
		viewData ViewData
		expect   validation.ValidationErrors
	}{
		{
			name:     "nil view data",
			viewData: nil,
			expect:   validation.ValidationErrors{},
		}, {
			name:     "empty view data",
			viewData: ViewData{},
			expect:   validation.ValidationErrors{},
		},
		{
			name: "view data without errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{},
			},
			expect: validation.ValidationErrors{},
		},
		{
			name: "view data with errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{
					"field1": "error1",
				},
			},
			expect: validation.ValidationErrors{
				"field1": "error1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, tc.viewData.GetErrors())
		})
	}

}

func TestViewDataError(t *testing.T) {
	testCases := []struct {
		name      string
		viewData  ViewData
		fieldName string
		expect    string
	}{
		{
			name:      "nil view data",
			viewData:  nil,
			fieldName: "field",
			expect:    "",
		}, {
			name:      "empty view data",
			viewData:  ViewData{},
			fieldName: "field",
			expect:    "",
		},
		{
			name: "view data without errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{},
			},
			fieldName: "field",
			expect:    "",
		},
		{
			name: "view data with errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{
					"field": "error1",
				},
			},
			fieldName: "field",
			expect:    "error1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, tc.viewData.Error(tc.fieldName))
		})
	}
}

func TestViewDataHasError(t *testing.T) {
	testCases := []struct {
		name      string
		viewData  ViewData
		fieldName string
		expect    bool
	}{
		{
			name:      "nil view data",
			viewData:  nil,
			fieldName: "field",
			expect:    false,
		}, {
			name:      "empty view data",
			viewData:  ViewData{},
			fieldName: "field",
			expect:    false,
		},
		{
			name: "view data without errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{},
			},
			fieldName: "field",
			expect:    false,
		}, {
			name: "different field",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{
					"bla": "error1",
				},
			},
			fieldName: "foo",
			expect:    false,
		}, {
			name: "view data with errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{
					"field": "error1",
				},
			},
			fieldName: "field",
			expect:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, tc.viewData.HasError(tc.fieldName))
		})
	}
}

func TestViewDataHasErrors(t *testing.T) {
	testCases := []struct {
		name     string
		viewData ViewData
		expect   bool
	}{
		{
			name:     "nil view data",
			viewData: nil,
			expect:   false,
		}, {
			name:     "empty view data",
			viewData: ViewData{},
			expect:   false,
		},
		{
			name: "view data without errors",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{},
			},
			expect: false,
		}, {
			name: "different field",
			viewData: ViewData{
				"Errors": validation.ValidationErrors{
					"bla": "error1",
				},
			},
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, tc.viewData.HasErrors())
		})
	}
}
