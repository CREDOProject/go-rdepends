package providers

import (
	"github.com/CREDOProject/go-rdepends/mappings"
)

type filename struct {
	mappingProviders []mappings.MappingsProvider
}

// Parse implements Provider.
func (f filename) Parse(path string) ([]Dependency, error) {
	dependencies := []Dependency{}
	for _, provider := range f.mappingProviders {
		for _, dependency := range provider.Get(path) {
			dependencies = append(dependencies, Dependency{
				Name:           dependency.Name,
				PackageManager: dependency.PackageManager,
				Suggestion:     false,
			})
		}
	}
	return dependencies, nil
}

func NewFilename(mappingProviders ...mappings.MappingsProvider) Provider {
	return filename{mappingProviders: mappingProviders}
}
