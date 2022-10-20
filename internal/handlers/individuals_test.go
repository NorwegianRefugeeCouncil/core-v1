package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/golang/mock/gomock"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/stretchr/testify/assert"
)

type mockRenderer struct{}

func (m mockRenderer) RenderView(w http.ResponseWriter, r *http.Request, templateName string, data viewParams) {
	w.WriteHeader(http.StatusOK)
}

var countries = []*api.Country{
	{
		ID:   "1",
		Name: "Country 1",
	}, {
		ID:   "2",
		Name: "Country 2",
	},
}

func Test_IndividualSelectedCountry(t *testing.T) {

	tests := []struct {
		name                  string
		selectedCountryID     string
		queryParamCountryID   string
		accessDenied          bool
		wantStatusCode        int
		wantHeaders           map[string]string
		additionalQueryParams map[string]string
		repoMock              func(t *testing.T, m *db.MockIndividualRepo)
		authMock              func(t *testing.T, m *auth.MockInterface)
	}{
		{
			// This test checks that a bad query parameter (such as invalid cast)
			// will result in a 400 Bad Request response.
			name:           "bad query param",
			wantStatusCode: http.StatusBadRequest,
			additionalQueryParams: map[string]string{
				"take": "abc",
			},
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Any()).Times(0)
			},
		},
		{
			// This test checks that a query parameter that is not a valid country ID
			// will result in a 404 Not Found response.
			name:                "query country id unknown",
			wantStatusCode:      http.StatusNotFound,
			queryParamCountryID: "abc",
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Any()).Times(0)
			},
		},
		{
			// This test checks that when neither the query parameter or session country ID
			// is set, the user is redirected to select a country
			name:           "no country selected. should redirect to select a country",
			wantStatusCode: http.StatusSeeOther,
			wantHeaders: map[string]string{
				"Location": "/countries/select",
			},
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Any()).Times(0)
			},
		}, {
			// This test checks that when only the session country ID is set, the user
			// will be redirected to the list within that country.
			name:              "selected country in session only. should redirect",
			selectedCountryID: "1",
			wantStatusCode:    http.StatusSeeOther,
			wantHeaders: map[string]string{
				"Location": "/individuals?country_id=1",
			},
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Any()).Times(0)
			},
		}, {
			// This test checks that when the query parameter country ID is set, the user
			// will access the list within that country.
			name:                "selected country in query param only. should use this country id",
			queryParamCountryID: "1",
			wantStatusCode:      http.StatusOK,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				expectOptions := api.GetAllOptions{CountryID: "1", Take: 20}
				m.EXPECT().GetAll(Any(), Eq(expectOptions)).Times(1)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("1")).Times(1).Return(true)
			},
		}, {
			// This test checks that when the query parameter country ID is not allowed for
			// the user, the user will receive a 403 Forbidden response.
			name:                "selected country in query param only, access denied. should return forbidden",
			queryParamCountryID: "1",
			wantStatusCode:      http.StatusForbidden,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("1")).Times(1).Return(false)
			},
		}, {
			// This test checks that when both the query parameter country ID and session country ID
			// are the same, the user will get the list for that country
			name:                "same selected country in session and query param",
			selectedCountryID:   "1",
			queryParamCountryID: "1",
			wantStatusCode:      http.StatusOK,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				expectOptions := api.GetAllOptions{CountryID: "1", Take: 20}
				m.EXPECT().GetAll(Any(), Eq(expectOptions)).Times(1)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("1")).Times(1).Return(true)
			},
		}, {
			// This test checks that when both the query parameter country ID and session country ID are the same
			// but the user does not have access to that country, the user will receive a 403 Forbidden response.
			name:                "same selected country in session and query param. access denied. should return forbidden",
			selectedCountryID:   "1",
			queryParamCountryID: "1",
			wantStatusCode:      http.StatusForbidden,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("1")).Times(1).Return(false)
			},
		}, {
			// This test checks that when the query parameter country ID and session country ID are different,
			// the user will get the list for the query parameter country ID
			name:                "different selected country in session and query param. Should use the country from query param",
			selectedCountryID:   "1",
			queryParamCountryID: "2",
			wantStatusCode:      http.StatusOK,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				expectOptions := api.GetAllOptions{CountryID: "2", Take: 20}
				m.EXPECT().GetAll(Any(), Eq(expectOptions)).Times(1)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("2")).Times(1).Return(true)
			},
		}, {
			// This test checks that when the query parameter country ID and session country ID are different
			// but the user doesn't have access to the query parameter country ID, the user will receive a 403 Forbidden response.
			name:                "different selected country in session and query param. access denied in both.",
			selectedCountryID:   "1",
			queryParamCountryID: "2",
			wantStatusCode:      http.StatusForbidden,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Any()).Times(1).Return(false)
			},
		}, {
			// This test checks that when the query parameter country ID and session country ID are different
			// and the user has access to the query parameter country ID, the user will get access to the
			// list of individuals within that country
			name:                "different selected country in session and query param. allowed only in query param",
			selectedCountryID:   "1",
			queryParamCountryID: "2",
			wantStatusCode:      http.StatusOK,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				expectOptions := api.GetAllOptions{CountryID: "2", Take: 20}
				m.EXPECT().GetAll(Any(), Eq(expectOptions)).Times(1)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("2")).Times(1).Return(true)
			},
		}, {
			// This test checks that when the query parameter country ID and session country ID are different
			// and the user is only allowed in the session country ID, the user will receive
			// a 403 Forbidden response.
			name:                "different selected country in session and query param. allowed only in session",
			selectedCountryID:   "1",
			queryParamCountryID: "2",
			wantStatusCode:      http.StatusForbidden,
			repoMock: func(t *testing.T, m *db.MockIndividualRepo) {
				m.EXPECT().GetAll(Any(), Any()).Times(0)
			},
			authMock: func(t *testing.T, m *auth.MockInterface) {
				m.EXPECT().CanReadWriteToCountryID(Eq("2")).Times(1).Return(false)
			},
		},
	}

	var setup = func(
		selectedCountryID string,
		queryParamCountryID string,
	) (*http.Request, *httptest.ResponseRecorder) {

		req := httptest.NewRequest(http.MethodGet, "/individuals", nil)
		rec := httptest.NewRecorder()

		ctx := req.Context()
		if len(selectedCountryID) != 0 {
			ctx = utils.WithSelectedCountryID(ctx, selectedCountryID)
		}
		if len(queryParamCountryID) != 0 {
			q := req.URL.Query()
			q.Add("country_id", queryParamCountryID)
			req.URL.RawQuery = q.Encode()
		}
		ctx = utils.WithCountries(ctx, countries)
		return req.WithContext(ctx), rec
	}

	r := &mockRenderer{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, rec := setup(tt.selectedCountryID, tt.queryParamCountryID)
			if tt.additionalQueryParams != nil {
				q := req.URL.Query()
				for k, v := range tt.additionalQueryParams {
					q.Add(k, v)
				}
				req.URL.RawQuery = q.Encode()
			}

			ctrl := NewController(t)
			defer ctrl.Finish()

			individualRepo := db.NewMockIndividualRepo(ctrl)

			authIntf := auth.NewMockInterface(ctrl)
			req = req.WithContext(utils.WithAuthContext(req.Context(), authIntf))

			if tt.authMock != nil {
				tt.authMock(t, authIntf)
			}
			if tt.repoMock != nil {
				tt.repoMock(t, individualRepo)
			}

			HandleIndividuals(r, individualRepo).ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatusCode, rec.Code)
			if tt.wantHeaders != nil {
				for k, v := range tt.wantHeaders {
					assert.Equal(t, v, rec.Header().Get(k))
				}
			}
		})
	}
}
