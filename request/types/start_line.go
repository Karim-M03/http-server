package types


type StartLine struct {
	Method Method
	Resource Resource
	Version string
}

func (r StartLine) String() string {
	return string(r.Method) + " " + r.Resource.String() + " " + r.Version
}
