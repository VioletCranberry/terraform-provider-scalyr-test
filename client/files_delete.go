package client

type DeleteFileRequest struct {
	Path            string `json:"path"`
	DeleteFile      bool   `json:"deleteFile"`
	ExpectedVersion int64  `json:"expectedVersion"`
}

type DeleteFileResponse struct {
	ApiResponse
}

func (client *ApiClient) DeleteConfigFile(file_path string) (*DeleteFileResponse, error) {
	request_body := &DeleteFileRequest{
		Path:       file_path,
		DeleteFile: true,
	}
	response, err := client.PostRequest("/api/putFile", request_body, &DeleteFileResponse{})
	return response.Result().(*DeleteFileResponse), err
}
