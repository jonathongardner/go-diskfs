package squashfs

import (
	"slices"
	"testing"
)

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
