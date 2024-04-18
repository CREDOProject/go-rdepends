package extractor

import "testing"

func Validate_Test(t *testing.T) {
	type testutil struct {
		Path   string
		Expect bool
	}

	testExpect := []testutil{
		{Path: "/tmp/a.tar.gz", Expect: true},
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
