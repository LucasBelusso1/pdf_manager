package handlers

import (
	"io"

	"github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/gin-gonic/gin"
)

func GetFilesFromMultipart(c *gin.Context) []dto.FileInput {
	var files []dto.FileInput
	form, err := c.MultipartForm()

	if err != nil {
		return files
	}

	for _, fileHeader := range form.File["files[]"] {
		if fileHeader.Header.Get("Content-Type") != "application/pdf" {
			continue
		}

		multipartFile, err := fileHeader.Open()

		if err != nil {
			continue
		}

		defer multipartFile.Close()

		fileContent, err := io.ReadAll(multipartFile)

		if err != nil {
			continue
		}

		files = append(files, dto.FileInput{
			Name:    fileHeader.Filename,
			Content: fileContent,
		})
	}

	return files
}
