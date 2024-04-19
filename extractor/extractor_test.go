package extractor

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func Test_Validate(t *testing.T) {
	type testutil struct {
		Path   string
		Expect bool
	}

	testExpect := []testutil{
		{Path: "/tmp/a.tar.gz", Expect: true},
		{Path: "/tmp/a.a.a.a.a.tar.gz", Expect: true},
		{Path: "/tmp/a.zip", Expect: false},
		{Path: "/tmp/a.tar", Expect: false},
		{Path: "/tmp/a.gz", Expect: false},
		{Path: "/tmp/foobar", Expect: false},
	}

	for _, v := range testExpect {
		validate := Validate(v.Path)
		if validate != v.Expect {
			t.Errorf("Expected %t for %s, got %t", v.Expect, v.Path, validate)
		}
	}
}

func Test_Extract(t *testing.T) {
	pathToExtract := "../test_assets/test.tar.gz"
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	path, err := Extract(path.Join(dir, pathToExtract))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", *path)
}

func Text_ExtractNoFile(t *testing.T) {
	pathToExtract := "../test_assets/notexists"
	dir, err := os.Getwd()
	if err == nil {
		t.Error(err)
	}
	path, err := Extract(path.Join(dir, pathToExtract))
	if err == nil {
		t.Error("Expected error")
	}
	fmt.Printf("%s", *path)
}
