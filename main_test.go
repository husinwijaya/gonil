package main

import (
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	os.Args = []string{"", "model/sample.go"}
	main()
	defer os.Remove("model/sample_nillable.go")
	out1, _ := os.ReadFile("model/sample_nillable.go")
	out2, _ := os.ReadFile("model/sample_nillable_result.go")
	if string(out1) != string(out2) {
		t.Fatal()
	}
}
