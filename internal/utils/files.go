package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

func UploadFile(nameForm string, id int, c echo.Context) (string, error) {
	file, err := c.FormFile(nameForm)
	if err != nil {
		return "", err
	}

	location := fmt.Sprintf("static/files/images/%d/avatar", id)
	fileUbication := fmt.Sprintf("static/images/%d/avatar/%s", id, file.Filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(location, os.ModePerm); err != nil {
		return "", err
	}

	dst, err := os.Create(fileUbication)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	host := c.Request().Host
	schema := ""
	if c.Request().TLS != nil {
		schema = "https"
	} else {
		schema = "http"
	}

	return fmt.Sprintf("%s://%s/%s", schema, host, fileUbication), nil
}
