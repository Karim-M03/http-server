package types


type Request struct {
	StartLine StartLine
	Headers map[string]string
	Body *Body
}
