package gordepends

import "github.com/CREDOProject/go-rdepends/providers"

// Retrieves the list of dependencies from a package.
func DependsOn(packagePath string, providersOptional ...providers.Provider) ([]providers.Dependency, error) {
	var configuredProviders []providers.Provider
	if len(providersOptional) > 0 {
		configuredProviders = providersOptional
	} else {
		configuredProviders = providers.DefaultProviders()
	}
	var dependencyList []providers.Dependency
	for _, provider := range configuredProviders {
		list, err := provider.Parse(packagePath)
		if err != nil {
			// TODO: Implement error logic.
		}
		dependencyList = append(dependencyList, list...)
	}
	return dependencyList, nil
}
