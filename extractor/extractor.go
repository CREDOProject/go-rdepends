package extractor

import (
	"path/filepath"
	"strings"
)

func Validate(path string) bool {
	ext := filepath.Ext(path)
	ext = strings.ToLower(ext)
	return ext == "tar.gz"
}
