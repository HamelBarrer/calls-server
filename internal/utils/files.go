package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(location, fileName string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(location, os.ModePerm); err != nil {
		return err
	}

	dst, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
