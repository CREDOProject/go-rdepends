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
func (a anticonf) Parse(extractpath string) ([]Dependency, error) {
	var dependencies []Dependency
	anticonfFiles, err := uf.ExecsInPath(extractpath, anticonfLooksLike)
	if err != nil {
		return nil, err
	}
	for _, anticonfEntry := range anticonfFiles {
		fileContent, err := os.ReadFile(anticonfEntry)
		if err != nil {
			return nil, err
		}
		for key, dep := range ap.Parse(string(fileContent)) {
			matchList := anticonfMgrRegex.FindStringSubmatch(key)
			if len(matchList) < 1 {
				continue
			}
			if pkgMgr, ok := pkgMgrMap[matchList[0]]; ok {
				dependencies = append(dependencies, Dependency{
					Name:           dep,
					PackageManager: pkgMgr,
					Suggestion:     false,
				})
			}
		}
	}
	return dependencies, nil
}

// Returns a new instance of anticonf Provider.
func NewAnticonf() Provider {
	return anticonf{}
}
