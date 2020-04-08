package pdfill

// EXE holds the link to the pdfill.exe
type EXE struct {
	exe string
}

const exePath string = `C:\Program Files (x86)\PlotSoft\PDFill\pdfill.exe`

var testPath string = ""

// FindEXE locates the pdfill.exe or returns an error
func FindEXE() (*EXE, error) {
	fn := exePath
	if testPath != "" {
		fn = testPath
	}

	if fileExists(fn) {
		exe := new(EXE)
		exe.exe = fn
		return exe, nil
	}

	return nil, &Error{
		Code:    EEXENOTFOUND,
		Message: "Unable to locate PDFill.exe",
	}
}

// TestMissingExe is provided for testing purposes only
func TestMissingExe() (*EXE, error) {
	testPath = "./NoSuchFile.XXX"
	x, err := FindEXE()
	testPath = ""
	return x, err
}
