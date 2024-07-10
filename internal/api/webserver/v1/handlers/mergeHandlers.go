package handlers

import (
	"net/http"

	"github.com/LucasBelusso1/pdf_manager/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Merge(c *gin.Context) {
	files := GetFilesFromMultipart(c)

	if len(files) < 2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Error when trying to get pdf",
			"code":    openMultipartFileErr,
		})
	}

	mergeUseCase := usecase.NewMergeUseCase(files, c.Writer)
	err := mergeUseCase.Merge()

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
