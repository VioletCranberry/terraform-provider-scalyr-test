package client

type PutFileRequest struct {
	Path            string `json:"path"`
	Content         string `json:"content"`
	PrettyPrint     bool   `json:"prettyprint"`
	ExpectedVersion int64  `json:"expectedVersion,omitempty"`
}

type PutFileResponse struct {
	ApiResponse
	Path string `json:"path"`
}

func (client *ApiClient) CreateConfigFile(file_path string, file_content string) (*PutFileResponse, error) {
	request_body := &PutFileRequest{
		Path:        file_path,
		Content:     file_content,
		PrettyPrint: true,
	}
	response, err := client.PostRequest("/api/putFile", request_body, &PutFileResponse{})
	return response.Result().(*PutFileResponse), err
}
