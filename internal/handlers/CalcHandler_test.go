package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/niiilov/go-yandex-1/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           models.Request
		expectedStatus int
		expectedResult interface{}
	}{
		{
			name:           "Valid Expression", //200
			method:         http.MethodPost,
			body:           models.Request{Expression: "2+2"},
			expectedStatus: http.StatusOK,
			expectedResult: models.Response{Result: "4"},
		},
		{
			name:           "Method Not Allowed", //405
			method:         http.MethodGet,
			body:           models.Request{},
			expectedStatus: http.StatusMethodNotAllowed,
			expectedResult: models.Errors{Error: "Method Not Allowed"},
		},
		{
			name:           "Unprocessable Entity", //422
			method:         http.MethodPost,
			body:           models.Request{Expression: "2+"},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResult: models.Errors{Error: "Unprocessable Entity"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.body)
			req := httptest.NewRequest(test.method, "/api/v1/calculate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			http.HandlerFunc(CalcHandler).ServeHTTP(recorder, req)

			if recorder.Code != test.expectedStatus {
				t.Errorf("Expected Status %d, Got %d", test.expectedStatus, recorder.Code)
			}

			var result models.Response
			if recorder.Code != http.StatusOK && recorder.Code == test.expectedStatus {
				return
			}

			err := json.NewDecoder(recorder.Body).Decode(&result)
			if err != nil {
				t.Fatalf("Failed To Decode Response: %v", err)
			}

			if result != test.expectedResult {
				t.Errorf("Expected Result %+v, Got %+v", test.expectedResult, result)
			}
		})
	}
}
