package pdfill_test

import (
	"fmt"
	"testing"

	"github.com/pjsoftware/pdfill"
)

func TestErrorTypes(t *testing.T) {
	testErrorOutput(t, nil, pdfill.ENOERROR,
		"", "", "")

	err1 := fmt.Errorf("Simple error")
	testErrorOutput(t, err1, pdfill.EINTERNAL,
		"An internal error has occurred: Simple error", "NoContext", "Simple error")

	err2 := fmt.Errorf("Extended error: %v", err1)
	testErrorOutput(t, err2, pdfill.EINTERNAL,
		"An internal error has occurred: Extended error: Simple error", "NoContext", "Extended error: Simple error")

	gte1 := &pdfill.Error{
		Code:    pdfill.EEXENOTFOUND,
		Message: "Unable to locate PDFill.exe",
	}
	testErrorOutput(t, gte1, pdfill.EEXENOTFOUND,
		"Unable to locate PDFill.exe", "NoContext", "<E_EXE_NOT_FOUND> Unable to locate PDFill.exe")

	gte2 := &pdfill.Error{Op: "Testing", Err: gte1}
	testErrorOutput(t, gte2, pdfill.EEXENOTFOUND,
		"Unable to locate PDFill.exe", "NoContext", "Testing: <E_EXE_NOT_FOUND> Unable to locate PDFill.exe")

	gte3 := &pdfill.Error{Op: "StillTesting", Context: "SomeContext", Err: gte2}
	testErrorOutput(t, gte3, pdfill.EEXENOTFOUND,
		"Unable to locate PDFill.exe", "SomeContext", "StillTesting/SomeContext: Testing: <E_EXE_NOT_FOUND> Unable to locate PDFill.exe")
}

func testErrorOutput(t *testing.T, err error, expc, expm, expx, exps string) {
	if got := pdfill.ErrorCode(err); got != expc {
		t.Errorf("Expected Code '%s' but got '%s'", expc, got)
	}
	if got := pdfill.ErrorMessage(err); got != expm {
		t.Errorf("Expected Message '%s' but got '%s'", expm, got)
	}
	if got := pdfill.ErrorContext(err); got != expx {
		t.Errorf("Expected Context '%s' but got '%s'", expx, got)
	}
	if err != nil {
		if got := err.Error(); got != exps {
			t.Errorf("Expected String '%s' but got '%s'", exps, got)
		}
	}
}

func testErrorCode(t *testing.T, err error, want string) {
	if err == nil {
		t.Errorf("Expected error '%s' but none occurred", want)
	}
	got := pdfill.ErrorCode(err)
	if got != want {
		t.Errorf("Expected error code %s; got %s", want, err)
	}
}

func testErrorContext(t *testing.T, err error, want string) {
	got := pdfill.ErrorContext(err)
	if got != want {
		t.Errorf("Expected error context '%s'; got '%s': %s", want, got, err)
	}
}
