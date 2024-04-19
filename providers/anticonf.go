package providers

import (
	"os"
	"path"
	"regexp"

	ap "github.com/CREDOProject/go-anticonf-parser"
)

var (
	anticonfRegex    = regexp.MustCompile("(?i)^configure.*")
	anticonfMgrRegex = regexp.MustCompile("^PKG_(.*)_NAME")
)

var pkgMgrMap = map[string]string{
	"DEB":  "apt",
	"RPM":  "rpm",
	"BREW": "brew",
}

func anticonfLooksLike(name string) bool {
	return anticonfRegex.MatchString(name)
}

// anticonf implements logic to parse anticonf files.
type anticonf struct{}

// Parse implements Provider.
func (a anticonf) Parse(extractpath string) ([]Dependency, error) {
	var dependencies []Dependency
	files, err := os.ReadDir(extractpath)
	if err != nil {
		return nil, err
	}
	var anticonfFiles []string
	for _, f := range files {
		if anticonfLooksLike(f.Name()) {
			anticonfFiles = append(anticonfFiles, f.Name())
		}
	}
	for _, anticonfEntry := range anticonfFiles {
		fileContent, err := os.ReadFile(path.Join(extractpath, anticonfEntry))
		if err != nil {
			return nil, err
		}
		for key, dep := range ap.Parse(string(fileContent)) {
			matchList := anticonfMgrRegex.FindStringSubmatch(key)
			if matchList == nil || len(matchList) < 1 {
				continue
			}
			if pkgMgr, ok := pkgMgrMap[matchList[1]]; ok {
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
