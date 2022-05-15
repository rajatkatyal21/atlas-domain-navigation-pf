package health_check_test

import (
	"dns/internal/app"
	"dns/internal/health_check"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_status(t *testing.T) {

	c := &app.Config{
		Name:    "dns",
		Version: "v1",
		Port:    3000,
	}

	s := app.NewServer(c)

	testCases := []struct {
		Name   string
		Result *health_check.Response
	}{
		{
			Name: "HealthCheck API OK",
			Result: &health_check.Response{
				Status: "OK",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			r, err := http.NewRequest("GET", "/dns/v1/status", nil)
			if err != nil {
				t.Fatalf("Failed to create the request %s", err.Error())
			}

			w := httptest.NewRecorder()

			s.InitRouter().ServeHTTP(w, r)

			if w.Code != 200 {
				t.Errorf("Test Failed. As expected 200 but got %d", w.Code)
			}

			res := &health_check.Response{}
			if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
				t.Errorf("error while unmarshalling the result %s", err.Error())
			}

			if !reflect.DeepEqual(res, tc.Result) {
				t.Errorf("unexpected result %v", res)
			}

		})

	}
}
