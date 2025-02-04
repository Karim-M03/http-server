package request_test

import (
	"bytes"
	"karim/http_server/request"
	"karim/http_server/request/types"
	"reflect"
	"testing"
)

func TestDivideMessage(t *testing.T) {
	var tests = []struct {
		name   string
		input  []byte
		output *types.Request
		errMsg string
	}{
		{
			name: "Correct Request",
			input: []byte("GET / HTTP/1.1\nHost: example.com\n\nBody Content"),
			output: &types.Request{
				StartLine: types.StartLine{
					Method: types.Method("GET"),
					Resource: types.Resource{
						Path:     []string{"", ""},
					},
					Version: "HTTP/1.1",
				},
				Headers: map[string]string{
					"Host": "example.com",
				},
				Body: &types.Body{Data: []byte("Body Content")},
			},
			errMsg: "",
		},
		{
			name:  "Wrong Request",
			input: []byte("GET / HTTP/1.2\nHost: example.com\n\nBody Content"),
			output: nil,
			errMsg: "Invalid Version. Version allowed HTTP/1.1",
		},
		{
			name:  "Wrong Method",
			input: []byte("CIAO / HTTP/1.1\nHost: example.com\n\nBody Content"),
			output: nil,
			errMsg: "Invalid Method",
		},
		{
			name:  "Wrong Header Format",
			input: []byte("GET / HTTP/1.1\nHost: example.com\nContent-Type application/json\n\nBody Content"),
			output: nil,
			errMsg: "Invalid header format",
		},
		{
			name:  "Multiple Headers",
			input: []byte("GET / HTTP/1.1\nHost: example.com\nContent-Type:application/json\n\nBody Content"),
			output: &types.Request{
				StartLine: types.StartLine{
					Method: types.Method("GET"),
					Resource: types.Resource{
						Path:     []string{"", ""},
					},
					Version: "HTTP/1.1",
				},
				Headers: map[string]string{
					"Host": "example.com",
					"Content-Type": "application/json",
				},
				Body: &types.Body{Data: []byte("Body Content")},
			},
			errMsg: "",
		},
		{
			name:  "Different Body",
			input: []byte("GET / HTTP/1.1\nHost: example.com\nContent-Type:application/json\n\n{\n\t'test':'hello'\n}"),
			output: &types.Request{
				StartLine: types.StartLine{
					Method: types.Method("GET"),
					Resource: types.Resource{
						Path:     []string{"", ""},
					},
					Version: "HTTP/1.1",
				},
				Headers: map[string]string{
					"Host": "example.com",
					"Content-Type": "application/json",
				},
				Body: &types.Body{Data: []byte("{\n\t'test':'hello'\n}")},
			},
			errMsg: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, _, err := request.DivideMessage(test.input)

			if test.output == nil {
				if request != nil {
					t.Errorf("Expected request to be nil, but got: %+v", request)
				}
				if err == nil || err.Error() != test.errMsg {
					t.Errorf("Expected error %q, but got: %v", test.errMsg, err)
				}
				return
			}

			if request.Body == nil {
				request.Body = &types.Body{Data: []byte{}}
			}
			if test.output.Body == nil {
				test.output.Body = &types.Body{Data: []byte{}}
			}

	

			t.Logf("Got: %#v", request)
			t.Logf("Want: %#v", test.output)
			t.Logf("Got Body: %q", request.Body.Data)
			t.Logf("Want Body: %q", test.output.Body.Data)

			if !reflect.DeepEqual(request.StartLine, test.output.StartLine) {
				t.Errorf("StartLine mismatch: got %+v, want %+v", request.StartLine, test.output.StartLine)
			}
			if !reflect.DeepEqual(request.Headers, test.output.Headers) {
				t.Errorf("Headers mismatch: got %+v, want %+v", request.Headers, test.output.Headers)
			}
			if !bytes.Equal(request.Body.Data, test.output.Body.Data) {
				t.Errorf("Body mismatch: got %q, want %q", request.Body.Data, test.output.Body.Data)
			}
		})
	}
}
