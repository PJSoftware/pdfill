package pdfill

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// ValidatePDF does simple validation on the specified file
func ValidatePDF(fn string) error {
	const op string = "pdfill.ValidatePDF"

	if fn[len(fn)-4:] != ".pdf" {
		return &Error{
			Code:    EBADFILENAME,
			Message: fmt.Sprintf("Expected 'pdf' extension: '%s'", fn),
		}
	}

	pdf, err := os.Open(fn)
	if err != nil {
		return &Error{Op: op, Context: "Opening", Err: err}
	}
	defer pdf.Close()

	scanner := bufio.NewScanner(pdf)
	lc := 0
	for scanner.Scan() && lc == 0 {
		txt := scanner.Text()
		lc++
		if ok, _ := regexp.MatchString(`^%PDF-`, txt); ok {
			return nil
		}
	}

	return &Error{
		Code:    EINVALIDPDF,
		Message: fmt.Sprintf("'%s' does not appear to be a valid PDF", fn),
	}
}
