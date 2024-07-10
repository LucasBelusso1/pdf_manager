package usecase

import (
	"bytes"
	"io"

	"github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type Merge struct {
	FileInputDtos []dto.FileInputDTO
	Writer        io.Writer
}

func NewMergeUseCase(fileInputDTOs []dto.FileInputDTO, writer io.Writer) *Merge {
	return &Merge{FileInputDtos: fileInputDTOs, Writer: writer}
}

func (muc *Merge) Merge() error {
	var readersForMerge []io.ReadSeeker
	for _, file := range muc.FileInputDtos {
		reader := bytes.NewReader(file.Content)

		readersForMerge = append(readersForMerge, reader)
	}

	err := api.MergeRaw(readersForMerge, muc.Writer, false, nil)

	if err != nil {
		return err
	}

	return nil
}
