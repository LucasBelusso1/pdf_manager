package dto

type FileInputDTO struct {
	Name    string
	Content []byte
}

func NewFile(name string, content []byte) *FileInputDTO {
	return &FileInputDTO{Name: name, Content: content}
}
