package domain_navigation_test

import (
	"bytes"
	"dns/internal/app"
	dns "dns/internal/domain_navigation"
	"dns/internal/models"
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestController_DataBankLocator(t *testing.T) {

	tests := []struct {
		Name          string
		c             *app.Config
		input         *dns.DataLocationRequest
		expected      *dns.Response
		expectedError *models.ErrorResponse
		expectedCode  int
	}{
		{
			Name: "Happy path with values in coordinate",
			input: &dns.DataLocationRequest{
				Vel: 20,
				X:   123.12,
				Y:   456.56,
				Z:   789.89,
			},
			expected: &dns.Response{
				Status:   models.OK,
				Location: 1389.57,
			},
			c: &app.Config{
				SectorId: 1,
				Name:     "dns",
				Version:  "v1",
				Port:     3002,
			},
			expectedCode: 200,
		},
		{
			Name: "Happy path with values in coordinate with Sector Id 0",
			input: &dns.DataLocationRequest{
				Vel: 20,
				X:   123.12,
				Y:   456.56,
				Z:   789.89,
			},
			expected: &dns.Response{
				Status:   models.OK,
				Location: 20,
			},
			c: &app.Config{
				SectorId: 0,
				Name:     "dns",
				Version:  "v1",
				Port:     3002,
			},
			expectedCode: 200,
		},
		{
			Name: "Happy path with 0 values in coordinate with Sector Id",
			input: &dns.DataLocationRequest{
				Vel: 20,
				X:   0,
				Y:   0,
				Z:   700.89,
			},
			expected: &dns.Response{
				Status:   models.OK,
				Location: 720.89,
			},
			c: &app.Config{
				SectorId: 1,
				Name:     "dns",
				Version:  "v1",
				Port:     3002,
			},
			expectedCode: 400,
		},
		{
			Name: "X coordinate missing the request",
			input: &dns.DataLocationRequest{
				Vel: 20,
				Y:   0,
				Z:   700.89,
			},
			expectedError: &models.ErrorResponse{
				Code:   400,
				Status: "NOK",
				Reason: "required field X is missing",
			},
			c: &app.Config{
				SectorId: 1,
				Name:     "dns",
				Version:  "v1",
				Port:     3002,
			},
			expectedCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			s := app.NewServer(test.c)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Errorf("error casued during the Marshalling: %s", err.Error())
			}

			r := httptest.NewRequest("POST", "/dns/v1/locate-data-bank", bytes.NewBuffer(body))
			s.InitRouter().ServeHTTP(w, r)

			if w.Code != test.expectedCode {
				t.Errorf("expected %d but got %d", test.expectedCode, w.Code)
			}

			res := &dns.Response{}
			er := &models.ErrorResponse{}
			if test.expectedCode == 200 {
				if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
					t.Errorf("error while unmarshalling the result %s", err.Error())
				}
				if !reflect.DeepEqual(res, test.expected) {
					t.Errorf("unexpected result %v", res)
				}

			} else {
				if err := json.Unmarshal(w.Body.Bytes(), er); err != nil {
					t.Errorf("error while unmarshalling the result %s", err.Error())
				}
				if test.expectedError != nil {
					if !reflect.DeepEqual(test.expectedError, er) {
						t.Errorf("unexpected result %v", res)
					}
				}
			}

		})

	}

}
