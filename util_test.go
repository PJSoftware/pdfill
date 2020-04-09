package pdfill_test

import (
	"testing"

	"github.com/pjsoftware/pdfill"
)

func TestFileExists(t *testing.T) {
	testCases := []struct {
		name string
		file string
		exp  bool
	}{
		{"Unpathed", "util_test.go", true},
		{"Pathed", "testfiles/notValid.pdf", true},
		{"Not Existing", "no_such.file", false},
		{"Folder", "testfiles", false},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if got := pdfill.FileExists(test.file); test.exp != got {
				t.Errorf("Expected %v, got %v", test.exp, got)
			}
		})
	}
}

func TestFolderExists(t *testing.T) {
	testCases := []struct {
		name string
		file string
		exp  bool
	}{
		{"Unpathed", "testfiles", true},
		{"Pathed", "testfiles/testout", true},
		{"Not Existing", "no_such_folder", false},
		{"File", "testfiles/notValid.pdf", false},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if got := pdfill.FolderExists(test.file); test.exp != got {
				t.Errorf("Expected %v, got %v", test.exp, got)
			}
		})
	}
}
