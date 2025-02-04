package response

import "strconv"

type StartLine struct {
	Version       string
	StatusCode    int
	StatusMessage string
}

func (s StartLine) String() string {
	return string(s.Version) + " " + strconv.Itoa(s.StatusCode) + " " + s.StatusMessage
}
