package squashfs

import (
	"reflect"
	"slices"
	"testing"
)

func TestSplitPathPreservesBackslashInName(t *testing.T) {
	p := `foo/bar\abc`
	expected := []string{"foo", `bar\abc`}

	actual := splitPath(p)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("splitPath(%q) = %#v, expected %#v", p, actual, expected)
	}
}

func TestSplitPathPreservesLiteralBackslash(t *testing.T) {
	got := splitPath(`/bar/baz\baz`)
	want := []string{
		"bar",
		`baz\baz`,
	}
	if !slices.Equal(got, want) {
		t.Fatalf("splitPath() = %#v, want %#v", got, want)
	}
}
