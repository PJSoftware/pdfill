package pdfill

// PDFill holds the link to the pdfill.exe
type PDFill struct {
	exe string
}

const exePath string = `C:\Program Files (x86)\PlotSoft\PDFill\pdfill.exe`

// FindPDFill locates the pdfill.exe or returns an error
func FindPDFill(fn ...string) (*PDFill, error) {
	ep := exePath
	if len(fn) > 0 {
		ep = fn[0]
	}

	if fileExists(ep) {
		exe := new(PDFill)
		exe.exe = ep
		return exe, nil
	}

	return nil, &Error{
		Code:    EEXENOTFOUND,
		Message: "Unable to locate PDFill.exe",
	}
}
