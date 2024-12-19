package file

// FileUploadSuccessResponse represents the response structure for a successful file upload operation.
type FileUploadSuccessResponse struct {
	Filename string `json:"filename"`
	FileType string `json:"fileType"`
	Url      string `json:"url"`
}

type FileLengthResponse struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
	Length   int    `json:"length"`
}
