package gordepends

import (
	"errors"
	"os"

	"github.com/CREDOProject/go-rdepends/extractor"
	"github.com/CREDOProject/go-rdepends/providers"
)

var (
	ErrNotTar = errors.New("File provided is not a tar.gz file.")
)

// Retrieves the list of dependencies from a package.
func DependsOn(packagePath string,
	providersOptional ...providers.Provider) ([]providers.Dependency, error) {
	if !extractor.Validate(packagePath) {
		return nil, ErrNotTar
	}
	extractPath, err := extractor.Extract(packagePath)
	defer os.RemoveAll(*extractPath)
	if err != nil {
		return nil, err
	}
	var configuredProviders []providers.Provider
	if len(providersOptional) > 0 {
		configuredProviders = providersOptional
	} else {
		configuredProviders = providers.DefaultProviders()
	}
	var dependencyList []providers.Dependency
	for _, provider := range configuredProviders {
		list, err := provider.Parse(*extractPath)
		if err != nil {
			// TODO: Implement error logic.
		}
		dependencyList = append(dependencyList, list...)
	}
	return dependencyList, nil
}
