package providers

import (
	"os"
	"regexp"

	"github.com/CREDOProject/sharedutils/files"
	"pault.ag/go/debian/control"
)

var (
	descriptionRegex = regexp.MustCompile("(?i)^DESCRIPTION")
)

func descriptionLooksLike(name string) bool {
	return descriptionRegex.MatchString(name)
}

type systemrequirements struct{}

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
			dependencies = append(dependencies, Dependency{
				Name:           dep,
				PackageManager: "",
				Suggestion:     true,
			})
			for matcher, data := range mappings {
				if ok := matcher.MatchString(dep); ok {
					dependencies = append(dependencies, data...)
					continue outer
				}
			}
		}
	}
	return dependencies, nil
}

// Returns a new instance of plain Provider.
func NewSystemRequirements() Provider {
	return systemrequirements{}
}

// Fuzzy package mappings.
var mappings = map[*regexp.Regexp][]Dependency{
	// Matches GNU make, make, ...
	regexp.MustCompile("(?i)^(?:.*)make$"): {
		{
			Name:           "make",
			PackageManager: APT_NAME,
			Suggestion:     false,
		},
	},
	// Matches libcurl, ...
	regexp.MustCompile("(?i)^(?:.*)libcurl$"): {
		{
			Name:           "libcurl-dev",
			PackageManager: APT_NAME,
			Suggestion:     false,
		},
	},
}
