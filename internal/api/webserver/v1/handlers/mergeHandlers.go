package handlers

import (
	"io"
	"net/http"

	"github.com/LucasBelusso1/pdf_manager/internal/usecase"
	"github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/gin-gonic/gin"
)

func Merge(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err,
			"code":    multipartErr,
		})
		return
	}

	var usecaseFileInputDTOs []dto.FileInputDTO
	for _, file := range form.File["files[]"] {
		if file.Header.Get("Content-Type") != "application/pdf" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   true,
				"message": "Incorrenct content type",
				"code":    contentTypeError,
			})
			return
		}

		multipartFile, err := file.Open()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   true,
				"message": err,
				"code":    openMultipartFileErr,
			})
			return
		}

		fileContent, err := io.ReadAll(multipartFile)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   true,
				"message": err,
				"code":    openMultipartFileErr,
			})
			return
		}

		defer multipartFile.Close()

		usecaseFileInputDTOs = append(usecaseFileInputDTOs, dto.FileInputDTO{
			Name:    file.Filename,
			Content: fileContent,
		})
	}

	mergeUseCase := usecase.NewMergeUseCase(usecaseFileInputDTOs, c.Writer)
	err = mergeUseCase.Merge()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err,
			"code":    openMultipartFileErr,
		})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Status(http.StatusOK)
}
