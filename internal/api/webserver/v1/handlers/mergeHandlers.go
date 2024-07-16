package handlers

import (
	"net/http"

	resterror "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/restError"
	"github.com/LucasBelusso1/pdf_manager/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Merge(c *gin.Context) {
	files, err := GetFilesFromMultipart(c, true)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if len(files) < 2 {
		err := resterror.NewBadRequestError("Not enought files")
		c.JSON(err.Code, err)
	}

	mergeUseCase := usecase.NewMergeUseCase(files, c.Writer)
	mergeError := mergeUseCase.Merge()

	if mergeError != nil {
		err := resterror.NewInternalServerError(mergeError.Error())
		c.JSON(err.Code, err)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Status(http.StatusOK)
}
