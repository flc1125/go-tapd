package tapd

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_ErrorResponse(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		want            string
		unwrapErr       error
		isErrorResponse bool
	}{
		{"error", errors.New("error"), "error", nil, false},                                                                                                                                         //nolint:lll
		{"error response", &ErrorResponse{err: errors.New("error")}, "error", errors.New("error"), true},                                                                                            //nolint:lll
		{"error response with raw body", &ErrorResponse{rawBody: &RawBody{Status: 1, Info: "info"}}, "code: 1, info: info", nil, true},                                                              //nolint:lll
		{"error response with http response", &ErrorResponse{response: &http.Response{StatusCode: 404}, err: errors.New("error")}, "status code: 404, err: error", errors.New("error"), true},       //nolint:lll
		{"error response with http response and raw body", &ErrorResponse{response: &http.Response{StatusCode: 404}, rawBody: &RawBody{Status: 1, Info: "info"}}, "code: 1, info: info", nil, true}, //nolint:lll
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.err.Error())
			assert.Equal(t, tt.unwrapErr, errors.Unwrap(tt.err))
			assert.Equal(t, tt.isErrorResponse, IsErrorResponse(tt.err))
		})
	}
}
