package handlers

import (
	"fmt"
	"io"

	resterror "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/restError"
	"github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/gin-gonic/gin"
)

func GetFilesFromMultipart(c *gin.Context, skipAnyError bool) ([]dto.FileInput, *resterror.RestError) {
	form, err := c.MultipartForm()

	if err != nil {
		return nil, resterror.NewInternalServerError("Error trying to get multipart form")
	}

	var files []dto.FileInput
	for _, fileHeader := range form.File["files[]"] {
		contentType := fileHeader.Header.Get("Content-Type")
		if contentType != "application/pdf" {
			if skipAnyError {
				errorMessage := fmt.Sprintf("Wrong content type of file %s: %s", fileHeader.Filename, contentType)
				return nil, resterror.NewBadRequestError(errorMessage)
			}

			continue
		}

		multipartFile, err := fileHeader.Open()

		if err != nil {
			if skipAnyError {
				return nil, resterror.NewInternalServerError(err.Error())
			}

			continue
		}

		defer multipartFile.Close()

		fileContent, err := io.ReadAll(multipartFile)

		if err != nil {
			if skipAnyError {
				return nil, resterror.NewInternalServerError(err.Error())
			}

			continue
		}

		files = append(files, dto.FileInput{
			Name:    fileHeader.Filename,
			Content: fileContent,
		})
	}

	return files, nil
}
