package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		endpoint       string
		body           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
		isJSON         bool
	}{
		{
			name:           "Valid POST /hello",
			method:         http.MethodPost,
			endpoint:       "/hello",
			body:           `{"name": "Somchai"}`,
			handler:        HelloHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Hello Somchai"}`,
			isJSON:         true,
		},
		{
			name:           "Invalid Method GET /hello",
			method:         http.MethodGet,
			endpoint:       "/hello",
			body:           "",
			handler:        HelloHandler,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method Not Allowed\n",
			isJSON:         false,
		},
		{
			name:           "Valid GET /terminate",
			method:         http.MethodGet,
			endpoint:       "/terminate",
			body:           "",
			handler:        TerminateHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is shutting down...\n",
			isJSON:         false,
		},
		{
			name:           "Invalid Method POST /terminate",
			method:         http.MethodPost,
			endpoint:       "/terminate",
			body:           "",
			handler:        TerminateHandler,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method Not Allowed\n",
			isJSON:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.endpoint, bytes.NewBufferString(tt.body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.handler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.isJSON && tt.expectedStatus == http.StatusOK {
				var actualResp Response
				json.Unmarshal(rr.Body.Bytes(), &actualResp)
				var expectedResp Response
				json.Unmarshal([]byte(tt.expectedBody), &expectedResp)
				assert.Equal(t, expectedResp, actualResp)
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}
