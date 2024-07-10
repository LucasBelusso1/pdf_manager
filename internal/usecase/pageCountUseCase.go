package usecase

import (
	"bytes"

	"github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type PageCount struct {
	Files []dto.FileInput
}

func NewPageCountUseCase(files []dto.FileInput) *PageCount {
	return &PageCount{Files: files}
}

func (pc *PageCount) CountPages() []dto.FileOutput {
	var outputFiles []dto.FileOutput
	for _, file := range pc.Files {
		var outputFile dto.FileOutput
		outputFile.Name = file.Name

		fileReader := bytes.NewReader(file.Content)
		pagesQty, err := api.PageCount(fileReader, nil)

		if err != nil {
			outputFile.Error = true
			outputFile.ErrorMessage = err.Error()
			outputFiles = append(outputFiles, outputFile)
			continue
		}

		outputFile.Pages = pagesQty
		outputFiles = append(outputFiles, outputFile)
	}

	return outputFiles
}
