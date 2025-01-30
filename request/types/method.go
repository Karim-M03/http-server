package types

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

// IsAllowedMethod checks if a given method is one of the allowed HTTP methods
func IsAllowedMethod(method string) bool {
	switch Method(method) {
	case GET, POST, PUT, OPTION, DELETE:
		return true
	default:
		return false
	}
}
