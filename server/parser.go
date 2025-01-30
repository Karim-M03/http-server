package server

import (
	"strings"
)

type StartLine struct{
	Method 		Method
	Resource	[]string
	Version		string
}



// parseInput analyze received data defining the action based on HTTP method.
func ParseInput(buffer []byte) string {
	request := strings.TrimSpace(string(buffer))
	if request == "" {
		return "Invalid request: empty input"
	}

	startLine, headers, body, err := DivideMessage(buffer)
	tokens := strings.Split(request, " ")
	if len(tokens) == 0 {
		return "Invalid request: no tokens found"
	}


	method := strings.ToUpper(tokens[0])
	if response, found := httpMethods[method]; found {
		return response
	}
	return "Invalid request: method not allowed"
}


