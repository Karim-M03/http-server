package request

import (
	"errors"
	"karim/http_server/constants"
	"karim/http_server/request/types"
	"strings"
)

// 1 -> if Start Line, Headers and Body are present
// 2 -> if only Start Line
// -1 invalid request
func DivideMessage(message []byte) (*types.Request, int, error) {
	raws := strings.Split(strings.ReplaceAll(string(message), "\r", ""), "\n")

	if len(raws) == 0 {
		return nil, -1, errors.New("Invalid request format")
	}

	var request types.Request
	request.Body = &types.Body{Data: []byte{}}

	startLine, err := StartLineValidation(raws[0])
	if err != nil {
		return nil, -1, err
	}

	request.StartLine = *startLine
	request.Headers = make(map[string]string)

	var i int
	for i = 1; i < len(raws); i++ {
		line := strings.TrimSpace(raws[i])
		if line == "" { // Empty line indicates end of headers
			i++
			break
		}
		arr := strings.SplitN(raws[i], ":", 2)
		if len(arr) != 2 {
			return nil, -1, errors.New("Invalid header format")
		}
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])

		request.Headers[key] = value
	}

	// join remaining lines as body
	if i < len(raws) {
		bodyStr := strings.Join(raws[i:], "\n")
		if bodyStr != "" {
			bodyData := types.Body{Data: []byte(bodyStr)}
			request.Body = &bodyData
		}
	}

	// determine return code
	if len(request.Headers) > 0 || string(request.Body.Data) != "" {
		return &request, 1, nil // start Line, Headers, and Body present
	}
	return &request, 2, nil // only Start Line present
}



func StartLineValidation(startLineString string)(*types.StartLine, error){
	startLineTokens := strings.Fields(strings.TrimSpace(startLineString))
	if len(startLineTokens) != 3{
		return nil, errors.New("Invalid StartLine Format")
	}

	methodString := startLineTokens[0]
	resourceString := startLineTokens[1]
	versionString := startLineTokens[2]

	if !types.IsAllowedMethod(methodString){
		return nil, errors.New("Invalid Method")
	}

	method := types.Method(methodString)

	if versionString != constants.PROTOCOL_VERSION{
		return nil, errors.New("Invalid Version. Version allowed HTTP/1.1")
	}

	resource, err := types.CreateResource(resourceString)
	if err != nil{
		return nil, err
	}

	startLine := &types.StartLine{
		Method: method,
		Resource: *resource,
		Version: versionString,
	}

	return startLine, nil
	
}