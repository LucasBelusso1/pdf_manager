package dto

type FileInput struct {
	Name    string
	Content []byte
}

type FileOutput struct {
	Name         string
	Pages        int
	Error        bool
	ErrorMessage string
}
