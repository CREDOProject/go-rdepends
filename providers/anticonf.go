package providers

// anticonf implements logic to parse anticonf files.
type anticonf struct{}

// Parse implements Provider.
func (a anticonf) Parse(path string) ([]Dependency, error) {
	return nil, nil
}

// Returns a new instance of anticonf Provider.
func NewAnticonf() Provider {
	return anticonf{}
}
