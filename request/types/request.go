package types


type Request struct {
	StartLine StartLine
	Headers map[string]string
	Body *Body
}




func (r Request) String() (string){

	startLineString := r.StartLine.String(true)
	var headersString string
	for key, value := range r.Headers{
		headersString += key + ": " + value + "\n"
	}
	
	bodyString := string(r.Body.Data)

	return startLineString + "\n" + headersString + "\n" + bodyString
	
}