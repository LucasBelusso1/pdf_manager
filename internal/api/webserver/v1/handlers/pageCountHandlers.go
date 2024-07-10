package handlers

import (
	"io"
	"net/http"

	handlersDTO "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/handlers/dto"
	"github.com/LucasBelusso1/pdf_manager/internal/usecase"
	usecaseDTO "github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/gin-gonic/gin"
)

func CountPage(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	var fileInputDTOs []usecaseDTO.FileInput
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

		fileInputDTOs = append(fileInputDTOs, usecaseDTO.FileInput{
			Name:    file.Filename,
			Content: fileContent,
		})
	}

	var pageCountOutputDTOs []handlersDTO.PageCountOutputDTO

	pageCountUseCase := usecase.NewPageCountUseCase(fileInputDTOs)
	filesWithPageCounter := pageCountUseCase.CountPages()

	for _, fileCounted := range filesWithPageCounter {
		pageCountOutputDTOs = append(pageCountOutputDTOs, handlersDTO.PageCountOutputDTO{
			Name:         fileCounted.Name,
			PagesQty:     fileCounted.Pages,
			Error:        fileCounted.Error,
			ErrorMessage: fileCounted.ErrorMessage,
		})
	}

	c.JSON(http.StatusOK, pageCountOutputDTOs)
}
