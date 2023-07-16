package client

type GetFileRequest struct {
	Path            string `json:"path"`
	ExpectedVersion int64  `json:"expectedVersion,omitempty"`
}

type GetFileResponse struct {
	ApiResponse
	Path       string  `json:"path"`
	Version    int64   `json:"version"`
	CreateDate ApiTime `json:"createDate"`
	ModDate    ApiTime `json:"modDate"`
	Content    string  `json:"content"`
}

func (client *ApiClient) GetConfigFile(file_path string) (*GetFileResponse, error) {
	request_body := &GetFileRequest{
		Path: file_path,
	}
	response, err := client.PostRequest("/api/getFile", request_body, &GetFileResponse{})
	return response.Result().(*GetFileResponse), err
}
