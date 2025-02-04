package types


type StartLine struct {
	Method Method
	Resource Resource
	Version string
}

func (s StartLine) String(withVersion bool) string {
	if withVersion{
		return string(s.Method) + " " + s.Resource.String() + " " + s.Version
	}
	return string(s.Method) + " " + s.Resource.String()
	
}
