package mappings

import (
	"regexp"
)

// Describes a Dependency to be parsed in the client application.
type Dependency struct {
	Name           string
	PackageManager string
}

// registryMappingProvider implements MappingProvider.
type registryMappingProvider struct {
	mappings map[*regexp.Regexp][]Dependency
}

func NewRegistryMappingProvider() registryMappingProvider {
	return registryMappingProvider{}
}

func (r registryMappingProvider) Get(name string) *Dependency {
	return nil
}
