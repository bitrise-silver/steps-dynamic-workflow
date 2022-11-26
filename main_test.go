package main

import (
	_ "embed"
	"os"
	"testing"
)

//go:embed tests/expected.yml
var expected string

func TestMatchingOutput(t *testing.T) {
	output, err := Generate("tests/bitrise.tmpl", "tests/values.yml")
	if err != nil {
		t.Fatalf("Error while generating template: %s", err)
	}
	f, _ := os.Create("tests/output.yml")
	defer f.Close()
	f.Write(output.Bytes())
	if expected != output.String() {
		t.Fatalf("Unmatched output")
	}
}
