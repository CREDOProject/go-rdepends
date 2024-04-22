package providers

const (
	APT_NAME  = "apt"
	RPM_NAME  = "rpm"
	BREW_NAME = "brew"
)

// Describes what is a Provider.
type Provider interface {
	// Parse the dependencies
	Parse(path string) ([]Dependency, error)
}

// Returns a slice of DefaultProviders.
func DefaultProviders() []Provider {
	return []Provider{
		NewAnticonf(),
		NewSystemRequirements(),
	}
}

// Describes a Dependency to be parsed in the client application.
type Dependency struct {
	Name           string
	PackageManager string
	Suggestion     bool
}
