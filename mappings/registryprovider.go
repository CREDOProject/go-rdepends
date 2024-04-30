package mappings

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

const REGISTRY_URL = "https://raw.githubusercontent.com/CREDOProject/package-suggestions/main/registry.json"

type Matcher struct {
	Matcher  string       `json:"matcher"`
	Packages []Dependency `json:"packages"`
}

// Describes a Dependency to be parsed in the client application.
type Dependency struct {
	Name           string `json:"name"`
	PackageManager string `json:"manager"`
}

// registryMappingProvider implements MappingProvider.
type registryMappingProvider struct {
	Mappings map[*regexp.Regexp][]Dependency
}

func NewRegistryMappingProvider() *registryMappingProvider {
	return &registryMappingProvider{}
}

func (r *registryMappingProvider) Get(name string) (deps []Dependency) {
	err := r.retrieveMappings()
	if err != nil {
		return nil
	}
	for regex, dependencies := range r.Mappings {
		if ok := regex.MatchString(name); ok {
			deps = append(deps, dependencies...)
		}
	}
	return
}

func (r *registryMappingProvider) retrieveMappings() (err error) {
	if len(r.Mappings) > 0 {
		return
	}
	r.Mappings = make(map[*regexp.Regexp][]Dependency)
	response, err := http.Get(REGISTRY_URL)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	proxyData := []Matcher{}
	err = json.Unmarshal(body, &proxyData)
	if err != nil {
		return
	}
	for _, p := range proxyData {
		regex, err := regexp.Compile(p.Matcher)
		if err != nil {
			return err
		}
		r.Mappings[regex] = append(r.Mappings[regex], p.Packages...)
	}
	return
}
