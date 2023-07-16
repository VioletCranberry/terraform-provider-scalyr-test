package client

type ListFilesResponse struct {
	ApiResponse
	Paths string `json:"paths"`
}

func (client *ApiClient) ListConfigFiles() (*ListFilesResponse, error) {
	response, err := client.PostRequest("/api/listFiles", `{}`, &ListFilesResponse{})
	return response.Result().(*ListFilesResponse), err
}
