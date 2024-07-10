package handlers

import (
	"net/http"

	handlersDTO "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/handlers/dto"
	"github.com/LucasBelusso1/pdf_manager/internal/usecase"
	"github.com/gin-gonic/gin"
)

func CountPage(c *gin.Context) {
	files := GetFilesFromMultipart(c)

	if len(files) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Error when trying to get pdf",
			"code":    openMultipartFileErr,
		})
	}

	var pageCountOutputDTOs []handlersDTO.PageCountOutputDTO

	pageCountUseCase := usecase.NewPageCountUseCase(files)
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
