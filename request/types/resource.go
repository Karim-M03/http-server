package types

import (
	"errors"
	"strings"
)

type Resource struct {
	Path     []string 
}

func (r Resource) String() string {
	return joinWithSlash(r.Path)
}

func joinWithSlash(parts []string) string {
	return strings.Join(parts, "/")
}

func CreateResource(resourceString string) (*Resource, error) {
	parts := strings.SplitN(resourceString, "://", 2)

	var pathString string

	if len(parts) == 1 {
		pathString = parts[0]
	} else {
		if parts[0] != "http" && parts[0] !="https"{
			return nil, errors.New("Invalid Protocol")
		}
		pathString = parts[1]
	}
	resource := &Resource{
		Path:     strings.Split(pathString, "/"),
	}

	return resource, nil
}