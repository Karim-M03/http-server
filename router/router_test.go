package router_test

import (
	"bytes"
	"fmt"
	"io"
	"karim/http_server/client"
	"karim/http_server/response"
	"karim/http_server/router"
	"net"
	"testing"
)

func TestRouter_HandleConnection(t *testing.T) {
	tests := []struct {
		name          string
		requestData   []byte
		expectedResp  string
		registerRoute bool
	}{
		{
			name:          "Valid request with registered route",
			requestData:   []byte("GET /test HTTP/1.1\r\nHost: localhost\r\n\r\n"),
			expectedResp:  "{\"status\": 200, \"message\": \"Route handled\", \"data\": null\"}",
			registerRoute: true,
		},
		{
			name:          "Valid request with unregistered route",
			requestData:   []byte("GET /unregistered HTTP/1.1\r\nHost: localhost\r\n\r\n"),
			expectedResp:  "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nRoute not found",
			registerRoute: false,
		},
		{
			name:          "Invalid request data",
			requestData:   []byte("INVALID REQUEST"),
			expectedResp:  "HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid request format",
			registerRoute: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			serverConn, clientConn := net.Pipe()
			defer clientConn.Close()
			defer serverConn.Close()



			// Create a new router
			r := router.NewRouter()
			if tt.registerRoute {
				err := r.RegisterEndpoint("GET /test", func(c *client.Client) {
					c.SendResponse(response.TmpResponse{
						Status:  200,
						Message: "Route handled",
						Data:    nil,
					})
				})

				if err != nil {
					t.Fatalf("Failed to register route: %v", err)
				}
			}

			// Run the server connection handler
			go r.HandleConnection(serverConn)

			// Send request
			_, err := clientConn.Write(tt.requestData)
			if err != nil {
				t.Fatalf("Failed to write request: %v", err)
			}

			// Read response
			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, clientConn)
			if err != nil && err != io.EOF {
				t.Fatalf("Failed to read response: %v", err)
			}
			
			t.Logf("Response: %s", buf.String())
		})
	}
}
