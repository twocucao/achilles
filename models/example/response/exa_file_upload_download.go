package response

import "achilles/models/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
