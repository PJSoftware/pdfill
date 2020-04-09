package pdfill

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Split takes an input PDF and splits it per page to specified folder
func (p *PDFill) Split(pdf, out, prefix string) (int, error) {
	const op string = "pdfill.Split"

	if !FileExists(pdf) {
		return 0, &Error{
			Code:    EFILENOTFOUND,
			Message: fmt.Sprintf("Split input file '%s' not found", pdf),
		}
	}

	if err := ValidatePDF(pdf); err != nil {
		return 0, &Error{Op: op, Context: "Validating", Err: err}
	}

	if !FolderExists(out) {
		return 0, &Error{
			Code:    EFOLDERNOTFOUND,
			Message: fmt.Sprintf("Split output folder '%s' not found", pdf),
		}
	}

	tempdir, err := ioutil.TempDir("", "PDFSplit-"+prefix)
	if err != nil {
		return 0, &Error{Op: op, Context: prefix, Err: err}
	}

	cmd := exec.Command(p.exe, "extract", "9999", pdf, tempdir+"/SPLIT.pdf")
	err = cmd.Run()
	return 0, err
}
