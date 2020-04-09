package pdfill_test

import (
	"testing"

	"github.com/pjsoftware/pdfill"
)

func TestFindExe(t *testing.T) {
	_, err := pdfill.FindPDFill()
	if err != nil {
		t.Fatalf("Error locating EXE: %s", err) // No exe; other tests WILL fail in odd ways
	}

	_, err = pdfill.FindPDFill("../NoSuchFile.exe")
	exp := pdfill.EEXENOTFOUND
	if got := pdfill.ErrorCode(err); got != exp {
		t.Errorf("Expected '%s' error code; got '%s'", exp, got)
	}

}
