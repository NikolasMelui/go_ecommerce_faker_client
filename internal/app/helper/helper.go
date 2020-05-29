package helper

import (
	"io"
	"os"
)

// CreateFakeFile ...
func CreateFakeFile(name string, size int64) error {

	sourceFile, err := os.Open("./data/sample/sample.pdf")
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	newFile, err := os.Create("./data/" + name)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
