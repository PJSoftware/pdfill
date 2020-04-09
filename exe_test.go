package pdfill_test

import (
	"testing"

	"github.com/pjsoftware/pdfill"
)

func TestFindExe(t *testing.T) {
	_, err := pdfill.FindPDFill()
	if err != nil {
		t.Errorf("Error locating EXE: %s", err)
	}

	_, err = pdfill.FindPDFill("../NoSuchFile.exe")
	exp := pdfill.EEXENOTFOUND
	if got := pdfill.ErrorCode(err); got != exp {
		t.Errorf("Expected '%s' error code; got '%s'", exp, got)
	}

}
