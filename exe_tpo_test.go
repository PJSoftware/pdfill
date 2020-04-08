package pdfill

// TestMissingExe is provided for testing purposes only
func MissingExeTest() (*EXE, error) {
	testPath = "./NoSuchFile.XXX"
	x, err := FindEXE()
	testPath = ""
	return x, err
}
