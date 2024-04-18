package providers

type plain struct{}

// Parse implements Provider.
func (p plain) Parse(path string) ([]Dependency, error) {
	return nil, nil
}

// Returns a new instance of plain Provider.
func NewPlain() Provider {
	return plain{}
}
