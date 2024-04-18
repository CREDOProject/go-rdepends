package providers

import (
	"os"
	"regexp"

	ap "github.com/CREDOProject/go-anticonf-parser"
	uf "github.com/CREDOProject/sharedutils/files"
)

var (
	anticonfRegex    = regexp.MustCompile("(?i)^configure.*")
	anticonfMgrRegex = regexp.MustCompile("^PKG_(.*)_NAME")
)

var pkgMgrMap = map[string]string{
	"DEB": "apt",
}

func anticonfLooksLike(name string) bool {
	return anticonfRegex.MatchString(name)
}

// anticonf implements logic to parse anticonf files.
type anticonf struct{}

// Parse implements Provider.
func (a anticonf) Parse(path string) ([]Dependency, error) {
	var dependencies []Dependency
	anticonfFiles, err := uf.ExecsInPath(path, anticonfLooksLike)
	if err != nil {
		return nil, err
	}
	for _, anticonfEntry := range anticonfFiles {
		fileContent, err := os.ReadFile(anticonfEntry)
		if err != nil {
			return nil, err
		}
		for key, dep := range ap.Parse(string(fileContent)) {
		}
	}
	return nil, nil
}

// Returns a new instance of anticonf Provider.
func NewAnticonf() Provider {
	return anticonf{}
}
