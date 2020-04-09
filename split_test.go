package pdfill_test

import (
	"testing"

	"github.com/pjsoftware/pdfill"
)

func TestSplit(t *testing.T) {
	const tf string = "testfiles/"

	testCases := []struct {
		name   string
		pdf    string
		folder string
		prefix string
		expC   int
		expEC  string
	}{
		{"Invalid File", "notValid.pdf", "testout", "notValid", 0, pdfill.EINVALIDPDF},
		{"Missing File", "noSuch.pdf", "testout", "notValid", 0, pdfill.EFILENOTFOUND},
		{"Missing Folder", "SinglePage.pdf", "noSuchFolder", "notValid", 0, pdfill.EFOLDERNOTFOUND},
		{"Single Page", "singlePage.pdf", "testout", "SP", 1, pdfill.ENOERROR},
		{"Multi-Page", "multiPage.pdf", "testout", "MP", 3, pdfill.ENOERROR},
	}

	pdf, _ := pdfill.FindPDFill()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			pdfin := tf + test.pdf
			folder := tf + test.folder
			gotC, gotE := pdf.Split(pdfin, folder, test.prefix)
			if gotC != test.expC {
				t.Errorf("Expected result count %d, got %d -- Err: %v", test.expC, gotC, gotE)
			}
			testErrorCode(t, gotE, test.expEC)
		})
	}
}
