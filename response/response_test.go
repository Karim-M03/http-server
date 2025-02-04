package response_test

import (
	"karim/http_server/response"
	"reflect"
	"testing"
)

func TestResponseType_String(t *testing.T) {
	startLine := response.StartLine{
		StatusCode: 200,
		StatusMessage: "OK",
		Version: "HTTP/1.1",
	}

	tests := []struct {
		name     string
		response response.ResponseType
		want     string
		wantErr  bool
	}{
		{
			name: "Test with valid data",
			response: response.ResponseType{
				StartLine: startLine,
				Headers: map[string]string{
					"Content-Type": "application/json",
					"Connection":   "keep-alive",
				},
				ResponseBody: []byte(`{"message": "success"}`),
			},
			want:    startLine.String() + "\nContent-Type: application/json\nConnection: keep-alive\n\n{\"message\": \"success\"}",
			wantErr: false,
		},
		{
			name: "Test with empty headers and body",
			response: response.ResponseType{
				StartLine:    startLine,
				Headers:      map[string]string{},
				ResponseBody: []byte(""),
			},
			want:    startLine.String() + "\n\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.response.String()
			if !reflect.DeepEqual((err != nil), tt.wantErr) {
				t.Errorf("ResponseType.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}