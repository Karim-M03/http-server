package types

import (
	"errors"
	"strings"
)

type Resource struct {
	Protocol *string   
	Path     []string 
}

func (r Resource) String() string {
	if r.Protocol != nil{
		return *r.Protocol + "://" + joinWithSlash(r.Path)
	}
	return joinWithSlash(r.Path)
}

func joinWithSlash(parts []string) string {
	return strings.Join(parts, "/")
}

func CreateResource(resourceString string) (*Resource, error) {
	parts := strings.SplitN(resourceString, "://", 2)

	var protocol *string
	var pathString string

	if len(parts) == 1 {
		// no protocol specifies
		pathString = parts[0]
	} else {
		protocol = &parts[0]
		pathString = parts[1]

		if *protocol != "http" && *protocol != "https" {
			return nil, errors.New("invalid protocol")
		}
	}

	resource := &Resource{
		Protocol: protocol,
		Path:     strings.Split(pathString, "/"),
	}

	return resource, nil
}