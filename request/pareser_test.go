package request_test

import (
	"fmt"
	"karim/http_server/request"
	"testing"
)



func TestDivideMessage(t *testing.T) {
	result, code, err := request.DivideMessage([]byte("GET / HTTP/1.1\nHost: example.com\n\nBody Content"))

	fmt.Printf("Request:\nStart Line: %s %s %s\nHeaders:\n",
		result.StartLine.Method, result.StartLine.Resource.String(), result.StartLine.Version)

	
	if code != 1{
		for key, value := range result.Headers {
			fmt.Printf("  %s: %s\n", key, value)
		}
	
		if result.Body != nil {
			fmt.Printf("Body:\n%s | Len %d", string(result.Body.Data), len(string(result.Body.Data)))
		} else {
			fmt.Println("Body: <nil>")
		}
	}

	fmt.Printf("Code: %d\n", code)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Error: <nil>")
	}
}

