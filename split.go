package pdfill

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
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

	tempdir, err := ioutil.TempDir("", "PDFSplit-"+prefix+"-")
	if err != nil {
		return 0, &Error{Op: op, Context: "TempDir", Err: err}
	}

	cmd := exec.Command(p.exe, "extract", "9999", pdf, tempdir+"/SPLIT.pdf")
	err = cmd.Run()
	if err != nil {
		sc := cleanupTempDir(tempdir)
		return sc, &Error{Op: op, Context: "Run", Err: err}
	}

	pc, err := moveFromTemp(tempdir, out, prefix)
	cleanupTempDir(tempdir)
	return pc, err
}

func cleanupTempDir(tf string) int {
	fc := 0
	folder, _ := os.Open(tf)
	files, _ := folder.Readdir(0)

	for _, file := range files {
		name := file.Name()
		path := tf + "/" + name
		os.Remove(path)
		fc++
		fmt.Println("Removed file:", path)
	}
	os.Remove(tf)
	return fc
}

func moveFromTemp(from string, to string, prefix string) (int, error) {
	op := "pdfill.moveFromTemp"
	fc := 0
	folder, _ := os.Open(from)
	files, _ := folder.Readdir(0)

	var err error
	re := regexp.MustCompile(`_(\d+)[.]pdf`)
	for _, file := range files {
		name := file.Name()
		fromPath := from + "/" + name

		toFile := prefix + "-Page-"
		match := re.FindAllStringSubmatch(name, -1)
		pn := `0000` + match[0][1]
		toFile += pn[len(pn)-4:] + ".pdf"
		toPath := to + "/" + toFile

		_, err = copyFile(fromPath, toPath)
		if err != nil {
			return 0, &Error{Op: op, Err: err}
		}
		fc++
	}
	return fc, err
}

func copyFile(src, dst string) (int64, error) {
	const op string = "pdfill.copyFile"
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, &Error{Op: op, Context: "SourceStat:" + src, Err: err}
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, &Error{Op: op, Context: "SourceStat", Err: fmt.Errorf("%s is not a regular file", src)}
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, &Error{Op: op, Context: "SourceOpen:" + src, Err: err}
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, &Error{Op: op, Context: "DestCreate:" + dst, Err: err}
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return 0, &Error{Op: op, Context: "Copy", Err: err}
	}

	return nBytes, nil
}
