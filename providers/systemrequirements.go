package providers

import (
	"os"
	"regexp"

	"github.com/CREDOProject/go-rdepends/mappings"
	"github.com/CREDOProject/sharedutils/files"
	"pault.ag/go/debian/control"
)

var (
	descriptionRegex = regexp.MustCompile("(?i)^DESCRIPTION")
)

func descriptionLooksLike(name string) bool {
	return descriptionRegex.MatchString(name)
}

type systemrequirements struct {
	mappingProviders []mappings.MappingsProvider
}

type DescriptionFile struct {
	SystemRequirements []string `control:"SystemRequirements" delim:", " strip:"\n\r\t "`
}

// Parse implements Provider.
func (p systemrequirements) Parse(extractpath string) ([]Dependency, error) {
	dependencies := []Dependency{}
	descFilePaths, err := files.FilesInPath(extractpath, descriptionLooksLike)
	if err != nil {
		return nil, err
	}
	for _, descFilePath := range descFilePaths {
		var data DescriptionFile
		descFile, err := os.Open(descFilePath)
		if err != nil {
			return nil, err
		}
		defer descFile.Close()
		err = control.Unmarshal(&data, descFile)
		if err != nil {
			return nil, err
		}
	outer:
		for _, dep := range data.SystemRequirements {
			for _, provider := range p.mappingProviders {
				if data := provider.Get(dep); data != nil {
					for _, d := range data {
						dependencies = append(dependencies, Dependency{
							Name:           d.Name,
							PackageManager: d.PackageManager,
							Suggestion:     false,
						})
					}
					continue outer
				}
			}
			dependencies = append(dependencies, Dependency{
				Name:           dep,
				PackageManager: "",
				Suggestion:     true,
			})
		}
	}
	return dependencies, nil
}

// Returns a new instance of plain Provider.
func NewSystemRequirements(mappingProviders ...mappings.MappingsProvider) Provider {
	return systemrequirements{mappingProviders: mappingProviders}
}
