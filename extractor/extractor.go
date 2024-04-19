package extractor

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var (
	targzre = regexp.MustCompile(`^.*\.tar\.gz$`)
)

// Validate a path to see if it is a tar.gz.
func Validate(path string) bool {
	last := filepath.Base(path)
	return targzre.MatchString(last)
}

// Extract a .tar.gz and returns a new temporary directory where the file is
// extracted.
func Extract(filePath string) (*string, error) {
	tempDirectory, err := os.MkdirTemp(os.TempDir(), "")
	if err != nil {
		return nil, err
	}

	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	uncompressedStream, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	tarReader := tar.NewReader(uncompressedStream)
	for true {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		headerName := path.Join(tempDirectory, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(headerName, 0755); err != nil {
				return nil, err
			}
		case tar.TypeReg:
			outFile, err := os.Create(headerName)
			if err != nil {
				return nil, err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return nil, err
			}
			err = outFile.Close()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Unknown type: %b in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return &tempDirectory, nil
}
