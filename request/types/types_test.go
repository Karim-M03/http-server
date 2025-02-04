package types_test

import (
	"karim/http_server/request/types"
	"reflect"
	"testing"
)

func TestRequest_String(t *testing.T) {
	startLine := types.StartLine{
		Method:  "GET",
		Resource: types.Resource{
			Path: []string{"", ""},
		},
		Version: "HTTP/1.1",
	}

	tests := []struct {
		name    string
		request types.Request
		want    string
	}{
		{
			name: "Test with valid data",
			request: types.Request{
				StartLine: startLine,
				Headers: map[string]string{
					"Content-Type": "application/json",
					"Connection":   "keep-alive",
				},
				Body: &types.Body{
					Data: []byte(`{"message": "success"}`),
				},
			},
			want: startLine.String(true) + "\nContent-Type: application/json\nConnection: keep-alive\n\n{\"message\": \"success\"}",
		},
		{
			name: "Test with empty headers and body",
			request: types.Request{
				StartLine: startLine,
				Headers:   map[string]string{},
				Body: &types.Body{
					Data: []byte(""),
				},
			},
			want: startLine.String(true) + "\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.request.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
