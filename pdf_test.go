package pdfill_test

import (
	"testing"

	"github.com/pjsoftware/pdfill"
)

func TestValidatePDF(t *testing.T) {
	testCases := []struct {
		name  string
		file  string
		expEC string
	}{
		{"Single Page", "testfiles/singlePage.pdf", pdfill.ENOERROR},
		{"Multi-Page", "testfiles/multiPage.pdf", pdfill.ENOERROR},
		{"Bad Name", "testfiles/singlePage", pdfill.EBADFILENAME},
		{"Invalid", "testfiles/notValid.pdf", pdfill.EINVALIDPDF},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := pdfill.ValidatePDF(test.file)
			testErrorCode(t, got, test.expEC)
		})
	}
}
