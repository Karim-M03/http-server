package server

// Method represents the allowed HTTP methods
type Method string

// Allowed methods
const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	OPTION Method = "OPTION"
	DELETE Method = "DELETE"
)
