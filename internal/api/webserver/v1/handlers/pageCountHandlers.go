package handlers

import (
	"net/http"

	handlersDTO "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/handlers/dto"
	resterror "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/restError"
	"github.com/LucasBelusso1/pdf_manager/internal/usecase"
	"github.com/gin-gonic/gin"
)

func CountPage(c *gin.Context) {
	files, err := GetFilesFromMultipart(c, false)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if len(files) == 0 {
		err := resterror.NewBadRequestError("Not enought files")
		c.JSON(err.Code, err)
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
