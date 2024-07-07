package usecase

import "github.com/LucasBelusso1/pdf_manager/internal/entities"

type MergeUseCase struct {
}

func NewMergeUseCase() *MergeUseCase {
	return &MergeUseCase{}
}

func (muc *MergeUseCase) Merge(fileEntities []entities.FileEntity) {

}
