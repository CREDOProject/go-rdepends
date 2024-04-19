package gordepends

import (
	"os"
	"path"
	"testing"
)

func Test_Depends(t *testing.T) {
	pathToDependency := "./test_assets/openssl_2.1.1.tar.gz"
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	newPath := path.Join(dir, pathToDependency)
	dep, err := DependsOn(newPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(dep)
	if len(dep) < 1 {
		t.Error("Expected at least one dependency.")
	}
}
