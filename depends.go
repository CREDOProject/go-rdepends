package gordepends

import (
	"errors"
	"os"
	"path"

	"github.com/CREDOProject/go-rdepends/extractor"
	"github.com/CREDOProject/go-rdepends/providers"
)

var (
	ErrNotTar = errors.New("File provided is not a tar.gz file.")
)

// Default configuredProviders.
var configuredProviders = providers.DefaultProviders()

// Retrieves the list of dependencies from a package.
func DependsOn(packagePath string,
	providersOptional ...providers.Provider) ([]providers.Dependency, error) {
	if !extractor.Validate(packagePath) {
		return nil, ErrNotTar
	}
	extractPath, err := extractor.Extract(packagePath)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(*extractPath)
	if len(providersOptional) > 0 {
		configuredProviders = providersOptional
	}
	var dependencyList []providers.Dependency
	// Reads all the subnodes of the extractPath.
	dirs, err := os.ReadDir(*extractPath)
	if err != nil {
		return nil, err
	}
	for _, provider := range configuredProviders {
		list, err := provider.Parse(*extractPath)
		if err != nil {
			// TODO: Implement error logic.
		}
		dependencyList = append(dependencyList, list...)
		for _, d := range dirs {
			if d.IsDir() { // Scan if it's a directory.
				list, err := provider.Parse(path.Join(*extractPath, d.Name()))
				if err != nil {
					// TODO: Implement error logic.
				}
				dependencyList = append(dependencyList, list...)
			}
		}
	}
	return dependencyList, nil
}
