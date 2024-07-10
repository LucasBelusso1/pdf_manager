package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/handlers/dto"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func CountPage(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	var pageCountOutputDTOs []dto.PageCountOutputDTO
	for _, file := range form.File["files[]"] {
		multipartFile, err := file.Open()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		defer multipartFile.Close()

		fileContent, err := io.ReadAll(multipartFile)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		reader := bytes.NewReader(fileContent)

		pagesQty, err := api.PageCount(reader, nil)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		pageCountOutputDTOs = append(pageCountOutputDTOs, dto.PageCountOutputDTO{
			Name:     file.Filename,
			PagesQty: pagesQty,
		})
	}

	c.JSON(http.StatusOK, pageCountOutputDTOs)
}
