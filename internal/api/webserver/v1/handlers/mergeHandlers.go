package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
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

	var filesForMerge []io.ReadSeeker
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

		reader := bytes.NewReader(fileContent)
		filesForMerge = append(filesForMerge, reader)
	}

	err = api.MergeRaw(filesForMerge, c.Writer, false, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err,
			"code":    mergeErr,
		})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Status(http.StatusOK)

}
