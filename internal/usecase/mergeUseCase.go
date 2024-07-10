package usecase

import (
	"bytes"
	"io"

	"github.com/LucasBelusso1/pdf_manager/internal/usecase/dto"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type MergeUseCase struct {
	FileInputDtos []dto.FileInputDTO
	Writer        io.Writer
}

func NewMergeUseCase(fileInputDTOs []dto.FileInputDTO, writer io.Writer) *MergeUseCase {
	return &MergeUseCase{FileInputDtos: fileInputDTOs, Writer: writer}
}

func (muc *MergeUseCase) Merge() error {
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
