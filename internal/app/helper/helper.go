package helper

import (
	"io"
	"os"
	"path"
)

// CreateFakeFile ...
func CreateFakeFile(fileName string, fileType string) error {
	var fileExtension string
	var targetFileDirectory string
	switch fileType {
	case "pdf":
		fileExtension = "pdf"
		targetFileDirectory = "documents"
	case "image":
		fileExtension = "png"
		targetFileDirectory = "images"
	}
	curDirectory := path.Dir(".")
	sourceFile, err := os.Open(path.Join(curDirectory, "samples", "sample"+"."+fileType))
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	newFile, err := os.Create(path.Join(curDirectory, "files", targetFileDirectory, fileName+"."+fileExtension))
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
