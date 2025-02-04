package response


type ResponseType struct{
	StartLine 		StartLine
	Headers			map[string]string
	ResponseBody 	[]byte
}


func (r ResponseType) String() (string, error){
	startLineString := r.StartLine.String()

	var headersString string
	for key, value := range r.Headers{
		headersString += key + ": " + value + "\n"
 	}

	responseBodyString := string(r.ResponseBody)


	return startLineString + "\n" + headersString + "\n" + responseBodyString, nil
}






