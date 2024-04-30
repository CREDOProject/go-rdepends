package mappings

type MappingsProvider interface {
	// Retrieves a dependency from a MappingsProvider.
	Get(string) []Dependency
}
