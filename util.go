package pdfill

import (
	"os"
)

// FileExists checks whether the named file exists
func FileExists(fn string) bool {
	info, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FolderExists checks whether the named file exists
func FolderExists(fn string) bool {
	info, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
